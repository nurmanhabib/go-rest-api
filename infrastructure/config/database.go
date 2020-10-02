package config

// DBConfig represent db config keys.
type DBConfig struct {
    DBDriver   string
    DBHost     string
    DBPort     string
    DBUser     string
    DBName     string
    DBPassword string
    DBLog      bool
}
