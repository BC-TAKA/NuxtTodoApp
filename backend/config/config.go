package config

type DBConfig struct {
	ID       string
	Password string
	Host     string
	Port     int
	DB       string
}

type Config struct {
	DB DBConfig
}

func Readconfig() (Config, error) {
	conf := Config{}

	conf.DB = DBConfig{
		ID:       "sample_id",
		Password: "sample_pass",
		Host:     "sample_host",
		Port:     3306,
		DB:       "todolist",
	}

	return conf, nil
}
