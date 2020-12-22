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
		ID:       "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "todo",
	}

	return conf, nil
}
