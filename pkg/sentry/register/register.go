package register

import (
	"context"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/paralus/paralus/pkg/sentry/cryptoutil"
	sentryrpc "github.com/paralus/paralus/proto/rpc/sentry"
	commonv3 "github.com/paralus/paralus/proto/types/commonpb/v3"
	"github.com/paralus/paralus/proto/types/sentry"
	"github.com/pkg/errors"
	"github.com/rs/xid"

	bootstrapclientv2 "github.com/paralus/paralus/api/def/clients/sentry/client"
	bootstrapapiv2 "github.com/paralus/paralus/api/def/clients/sentry/client/bootstrap"

	"github.com/paralus/paralus/pkg/log"
	"github.com/paralus/paralus/pkg/sentry/util"
)

var (
	_log = log.GetLogger()
)

const (
	schemeHTTP  = "http"
	schemeHTTPS = "https"
	schemeGRPC  = "grpc"

	modeServer = "server"
	modeClient = "client"
)

// Config is the config used for client registration
type Config struct {
	// TemplateToken is token of the template the client is registering to
	TemplateToken string
	// Scheme is the scheme for talking to the registration endpoint
	// accepted schemes http/https/grpc

	// TemplateName is the name of the boostrap template
	// it can be passed instead of template token to
	// auto register servers
	TemplateName string

	Scheme string

	Mode string

	// Addr is the address of the registation endpoint
	Addr string

	// ClientID is the unique client identifier
	// it can be supplied or generated by the client
	// to perform auto registration
	ClientID string

	// ClientIP is the ip at which client is reachable
	ClientIP string

	// Name of the client
	// can be part of the supplied client config or
	// generated by the client to perform auto registration
	Name string

	// PEM encoded private key used to generate
	// CSR for client registration
	// If not set private key is generated before registation
	// which can be accessed from config
	PrivateKey []byte

	// CSR is the certificate signing request
	// this is auto generated based on the PrivateKey and Token
	CSR []byte

	// Certificate is the certificate signed certificate
	// returned by bootstrap service after successful registration
	Certificate []byte

	// CACertificate is the CA Certificate of the Issuer returned by
	// the bootstrap service after successful registration
	CACertificate []byte

	// ServerHost is the host the server
	ServerHost string

	// ServerPort is port the registered server should listen on
	// it is returned after registration
	ServerPort int
}

func registerHTTP(ctx context.Context, config *Config) error {

	tc := &bootstrapclientv2.TransportConfig{
		Host:     config.Addr,
		BasePath: "/",
		Schemes:  []string{config.Scheme},
	}

	bc := bootstrapclientv2.NewHTTPClientWithConfig(nil, tc)

	transport := http.DefaultTransport.(*http.Transport).Clone()

	if bootstrapCA := os.Getenv("BOOTSTRAP_CA_CERT"); strings.TrimSpace(bootstrapCA) != "" {
		_log.Infow("using rootCA for bootstrap")
		rootCAs := x509.NewCertPool()
		decodedCA, err := base64.StdEncoding.DecodeString(bootstrapCA)
		if err != nil {
			err = errors.Wrap(err, "unable to decode bootstrap CA")
			return err
		}
		if ok := rootCAs.AppendCertsFromPEM(decodedCA); !ok {
			err = errors.New("unable to append bootstrap CA to cert pool")
			return err
		}
		transport.TLSClientConfig.RootCAs = rootCAs
	} else if insecure := os.Getenv("ALLOW_INSECURE_BOOTSTRAP"); strings.TrimSpace(insecure) == "true" {
		_log.Infow("using InsecureSkipVerify for bootstrap")
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	resp, err := bc.Bootstrap.BootstrapRegisterBootstrapAgent(&bootstrapapiv2.BootstrapRegisterBootstrapAgentParams{
		Context:       ctx,
		TemplateToken: fmt.Sprintf("template/%s", config.TemplateToken),
		Body: bootstrapapiv2.BootstrapRegisterBootstrapAgentBody{
			Token:     config.ClientID,
			Csr:       config.CSR,
			IPAddress: config.ClientIP,
			Name:      config.Name,
		},
		HTTPClient: httpClient,
	}, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	config.Certificate = resp.Payload.Certificate
	config.CACertificate = resp.Payload.CaCertificate

	return nil
}

func registerGRPC(ctx context.Context, config *Config) error {
	sp := sentryrpc.NewSentryPool(config.Addr, 1)

	spc, err := sp.NewClient(ctx)
	if err != nil {
		return err
	}

	// if mode is server
	// fetch host from template
	if config.Mode == modeServer {
		template, err := spc.GetBootstrapAgentTemplate(ctx, &sentry.BootstrapAgentTemplate{
			Metadata: &commonv3.Metadata{
				Name: config.TemplateName,
			},
		})
		if err != nil {
			return err
		}

		// Only for external services
		for _, host := range template.Spec.Hosts {
			if host.Type == sentry.BootstrapTemplateHostType_HostTypeExternal {
				config.ServerHost, config.ServerPort = util.ParseAddr(host.Host)
			}
		}

		privKey, err := cryptoutil.DecodePrivateKey(config.PrivateKey, cryptoutil.NoPassword)
		if err != nil {
			return err
		}

		csr, err := cryptoutil.CreateCSR(pkix.Name{
			CommonName:         template.Metadata.Description,
			Country:            []string{"USA"},
			Organization:       []string{"Rafay Systems Inc"},
			OrganizationalUnit: []string{template.Metadata.Description},
			Province:           []string{"California"},
			Locality:           []string{"Sunnyvale"},
		}, privKey)
		if err != nil {
			return err
		}

		config.CSR = csr
	}

	if config.TemplateToken == "" {
		config.TemplateToken = "template/-"
	}

	resp, err := spc.RegisterBootstrapAgent(ctx, &sentryrpc.RegisterAgentRequest{
		TemplateToken: config.TemplateToken,
		TemplateName:  config.TemplateName,
		Csr:           config.CSR,
		Name:          config.Name,
		IpAddress:     config.ClientIP,
		Token:         config.ClientID,
	})
	if err != nil {
		return err
	}

	config.CACertificate = resp.CaCertificate
	config.Certificate = resp.Certificate

	return nil
}

func prepareConfig(config *Config) error {
	if config.Addr == "" {
		return fmt.Errorf("empty registration Addr not allowed")
	}

	switch config.Scheme {
	case schemeHTTP, schemeHTTPS, schemeGRPC:
	case "":
	default:
		return fmt.Errorf("invalid scheme %s", config.Scheme)
	}

	if config.Scheme == "" {
		_log.Debug("scheme not defined, defaulting to http")
		config.Scheme = schemeHTTP
	}

	switch config.Mode {
	case modeServer, modeClient:
	case "":
	default:
		return fmt.Errorf("invalid mode %s", config.Mode)
	}

	if config.Mode == "" {
		config.Mode = modeClient
	}

	if config.TemplateToken == "" && config.Mode == modeClient {
		return fmt.Errorf("template token cannot be empty")
	}

	if config.PrivateKey == nil {
		privKey, err := cryptoutil.GenerateECDSAPrivateKey()
		if err != nil {
			return err
		}

		key, err := cryptoutil.EncodePrivateKey(privKey, cryptoutil.NoPassword)
		if err != nil {
			return err
		}

		config.PrivateKey = key

	}

	if config.ClientID == "" {
		config.ClientID = xid.New().String()
	}

	// for server csr is generated from host of the template
	if config.CSR == nil && config.Mode == modeClient {
		privKey, err := cryptoutil.DecodePrivateKey(config.PrivateKey, cryptoutil.NoPassword)
		if err != nil {
			return err
		}

		csr, err := cryptoutil.CreateCSR(pkix.Name{
			CommonName:         config.ClientID,
			OrganizationalUnit: []string{config.TemplateToken},
		}, privKey)
		if err != nil {
			return err
		}

		config.CSR = csr
	}

	return nil
}

// Register registers in bootstrap service
func Register(ctx context.Context, config *Config) error {
	err := prepareConfig(config)
	if err != nil {
		return err
	}

	switch config.Scheme {
	case schemeHTTP, schemeHTTPS:
		err = registerHTTP(ctx, config)
	case schemeGRPC:
		err = registerGRPC(ctx, config)
	}
	return err
}
