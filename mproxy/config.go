package mproxy

// Config application wide config object
type Config struct {
	Profile string
	Auth    auth   `toml:"auth"`
	Remote  remote `toml:"remote"`
}

type auth struct {
	Username string
	Password string
	Port     int
}

type remote struct {
	Username string
	Password string
	Host     string
	Port     int
}
