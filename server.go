package remotectrl

// ServerConfig holds information for the remotectrl server
type ServerConfig struct {
	// Port is the port which will be used fot the server to listen
	Port string
	// User for the registration on the remotectrl server
	User string
	// Pass is the password phrase for the user to login on the server
	Pass string
}
