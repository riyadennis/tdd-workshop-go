package main
type DBConnection struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}
type APIConfig struct {
	Host string
	Port string
}
type Config struct {
	DB DBConnection
	API APIConfig
}

type ConfigOption interface {
	ApplyConfig(*Config)
}
func NewConfig(){

}
