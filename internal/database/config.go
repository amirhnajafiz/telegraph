package database

// Config type for database.
// ConnectionTimeout: for connect context.
// Database: mongodb database name.
// MongoURL: URL for mongo cluster.
type Config struct {
	ConnectionTimeout int    `koanf:"connection_timeout"`
	Database          string `koanf:"database"`
	MongoURL          string `koanf:"mongo_url"`
}
