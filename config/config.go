package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title      string      `yaml:"title"`
	Endpoints  []*Endpoint `yaml:"endpoints"`
	WebhookURL string      `yaml:"slack_webhook_url"`
}

type Endpoint struct {
	// Enable defines whether to enable the monitoring of the endpoint
	Enable bool `yaml:"enable" json:"enable"`
	// Name of the endpoint. Can be anything.
	Name string `yaml:"name" json:"name"`
	// URL to send the request to
	URL string `yaml:"url" json:"url"`
	// Method of the request made to the url of the endpoint
	Method string `yaml:"method,omitempty" json:"method,omitempty"`
	// Interval is the duration to wait between every status check
	Interval time.Duration `yaml:"interval,omitempty" json:"interval,omitempty"`
	// Condition used to determine the health of the endpoint
	Condition *Condition `yaml:"condition,omitempty" json:"condition,omitempty"`
}

type Condition struct {
	StatusCode int               `yaml:"status,omitempty"`
	Headers    map[string]string `yaml:"headers,omitempty"`
	Body       string            `yaml:"body,omitempty"`
}

// Load configuration file.
func Load(configFile string) (*Config, error) {
	cfg, err := readConfigurationFile(configFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func readConfigurationFile(fileName string) (config *Config, err error) {
	var bytes []byte
	if bytes, err = ioutil.ReadFile(fileName); err == nil {
		return parseAndValidateConfigBytes(bytes)
	}
	return
}

// parseAndValidateConfigBytes parses configuration file into a Config struct and validates.
func parseAndValidateConfigBytes(yamlBytes []byte) (config *Config, err error) {
	yamlBytes = []byte(os.ExpandEnv(string(yamlBytes)))
	if err = yaml.Unmarshal(yamlBytes, &config); err != nil {
		return
	}

	return
}
