package configs

var Conf configurations

type configurations struct {
	Server serverConfigurations
}

type serverConfigurations struct {
	Port int
}
