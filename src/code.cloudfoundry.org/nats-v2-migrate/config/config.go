package config

import (
	"encoding/json"
	"io/ioutil"

	"code.cloudfoundry.org/lager/lagerflags"
)

type Config struct {
	Address                         string   `json:"address"`
	Bootstrap                       bool     `json:"bootstrap"`
	NATSInstances                   []string `json:"nats_instances"`
	NATSPort                        int      `json:"nats_port"`
	NATSMigratePort                 int      `json:"nats_migrate_port"`
	NATSMigrateServers              []string `json:"nats_migrate_servers"`
	NATSMigrateServerCAFile         string   `json:"nats_migrate_server_ca_file"`
	NATSMigrateServerClientCertFile string   `json:"nats_migrate_server_client_cert_file"`
	NATSMigrateServerClientKeyFile  string   `json:"nats_migrate_server_client_key_file"`
	NATSV1BinPath                   string   `json:"nats_v1_bin_path"`
	NATSV2BinPath                   string   `json:"nats_v2_bin_path"`
	NATSConfigPath                  string   `json:"nats_config_path"`
	lagerflags.LagerConfig
}

func NewConfig(configPath string) (Config, error) {
	var cfg Config
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
