package configs

type AppConfig struct {
	LyricFetchLimits uint
	ApiServerPort    string
}

func LoadAppConfig() *AppConfig {
	return &AppConfig{
		LyricFetchLimits: 250,
		ApiServerPort:    "443",
	}
}
