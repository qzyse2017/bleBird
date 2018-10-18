package config


type ServerConfig struct {
	Port string `yaml:"port"`
	dbName string `yaml:"dbname"`
	dbConn string `yaml:"dbConnn"`
}

func ConfOrDefault(conf *ServerConfig) {
	if conf.Port = "" {
		conf.Port = 
	}
}

func LoadFromYaml(data) {

}