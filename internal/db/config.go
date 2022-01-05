package db

type Config struct {
	Name string `koanf:"name"`
	URL  string `koanf:"url"`
}
