package config

// Database struct for db connection
type Database struct {
	Host, Username, Password, Name string
	Port                           int
}
