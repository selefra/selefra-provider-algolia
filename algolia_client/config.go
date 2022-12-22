package algolia_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	AppID  string `yaml:"app_id,omitempty" mapstructure:"app_id"`
	APIKey string `yaml:"api_key,omitempty" mapstructure:"api_key"`
}
