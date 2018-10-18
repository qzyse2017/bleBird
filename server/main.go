package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/qzyse2017/bleBird/server/config"
)

func main() {
	configFileName := flag.String("conf", "./server_conf.yaml", "the configuration file path" )
	conf := loadConfig(*configFileName)

}

func loadConfig(confFilePath string) *config.serverConfig {
	var conf *config.ServerConfig

	if _, err := os.Stat(confFilePath); os.isNotExist(err) {
		log.Fatal(err)
		//TODO--let the logger be pluggable
	} else {
		//file exist, load confifuration from file
		data, err := ioutil.ReadFile(confFilePath)
		if err != nil {
			log.Fatal(err)
		}

		conf, err = config.LoadFromYaml(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	conf = config.ConfOrdefault(conf)
	return conf
	//TODO -- use factory method and bridge pattern to make data related operation pluggable.
}

