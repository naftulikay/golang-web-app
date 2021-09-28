package cmdConstants

const (
	DefaultEnvironment = "dev"
	DefaultMySQLHost   = "localhost"
	DefaultMySQLPort   = uint16(3306)
	DefaultListenHost  = "127.0.0.1"
	DefaultListenPort  = uint16(8080)
)

func DefaultValues() []interface{} {
	return []interface{}{
		DefaultEnvironment, DefaultMySQLHost, DefaultMySQLPort, DefaultListenHost, DefaultListenPort,
	}
}
