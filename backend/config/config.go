package config

type DB struct {
	ID       string
	Password string
	Host     string
	Port     string
	DB       string
}

func Readconfig() DB {
	var d DB
	d.ID = "a1"
	d.Password = "a2"
	d.Host = "a3"
	d.Port = "a4"
	d.DB = "a5"
	return d
}
