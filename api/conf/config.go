package conf

type Config struct {
	App   *app   `toml:"app" json:"app"`
	MySQL *mysql `toml:"mysql" json:"mysql"`
	Auth  *auth  `toml:"auth" json:"auth"`
}
