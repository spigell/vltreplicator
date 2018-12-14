package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"io/ioutil"

        "gopkg.in/yaml.v2"
        "time"


        "vltreplicator/replication"
)

var (
        config       = flag.String("config", "/etc/vault-replicator.yml", "path to config file")
        version      = flag.Bool("version", false, "show version")
        BuildVersion = "None"
)

type Config struct {
        Reader        
        Writer       
        Reload_period int
}

type Reader struct {
	Address  string
	Path     string
	Version  int
}

type Writer struct {
	Address  string
	Path     string
	Version  int
}



func main() {

        flag.Parse()
        if *version {
                fmt.Printf("%s\n", BuildVersion)
                os.Exit(0)
        }

        file, _ := os.Open(*config)
        configuration := Config{}
        target, _ := ioutil.ReadAll(file)

        err := yaml.Unmarshal(target, &configuration)
        if err != nil {
                log.Fatal("[ERROR] Error while parsing configuration: ", err)
        }

	config := &replication.ReplicaConfig{ReaderAddress: configuration.Reader.Address, ReaderPath: configuration.Reader.Path, ReaderVersion: configuration.Reader.Version,
		WriteAddress: configuration.Writer.Address, WritePath: configuration.Writer.Path }
        replication.Replicate(config)

        ticker := time.NewTicker(time.Duration(configuration.Reload_period) * time.Second)
        go func() {
        	for t := range ticker.C {
			log.Print("Start replication at ", t)
		        replication.Replicate(config)
        	}
        }()

        select{ }
}