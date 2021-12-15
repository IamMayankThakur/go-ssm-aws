package ssm

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
)

var (
	ErrInvalidConfig       = errors.New("config cannot be nil")
	ErrConfigInvalidRegion = errors.New("region not found: ssm region is a required config")
	ErrSSMNotEnabled       = errors.New("ssm is disabled")
)

type Config struct {
	Enabled     bool
	SecretsPath string
	Region      string
}

type Client struct {
	client *ssm.SSM
	config *Config
}

func New(cfg *Config) (*Client, error) {
	if cfg == nil {
		return nil, ErrInvalidConfig
	}
	if !cfg.Enabled {
		return nil, ErrSSMNotEnabled
	}
	if cfg.Region == "" {
		return nil, ErrConfigInvalidRegion
	}
	opts := session.Options{
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Region:                        aws.String(cfg.Region),
		},
	}
	sess := session.Must(session.NewSessionWithOptions(opts))

	clt := Client{client: ssm.New(sess), config: cfg}
	return &clt, nil
}

func (c *Client) GetValueByName(name string, decrypt bool) (string, error) {
	path := fmt.Sprintf("%s/%s", c.config.SecretsPath, name)
	input := ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(decrypt),
	}

	out, err := c.client.GetParameter(&input)
	if err != nil {
		log.Fatalf("unable to get parameter from SSM at path: %s", path)
		return "", err
	}

	return *out.Parameter.Value, nil
}
