package replication

import (
	"github.com/hashicorp/vault/api"

	"log"
	"os"
)

type ReplicaConfig struct {
	ReaderAddress string
	ReaderPath string
	ReaderVersion int
	WriteAddress string
	WritePath    string
	WriteVersion int
	client       *api.Client
}

func Replicate( c *ReplicaConfig ) {

        readerToken := os.Getenv("READER_TOKEN")
        writeToken := os.Getenv("WRITE_TOKEN")

	err := c.CreateConnection("reader", readerToken)

	if err != nil {
		log.Fatal(err)
	}
	keys, err := c.ReadKeys()
	if err != nil {
		log.Fatal(err)
	}

	storage := make(map[string]string)

	for _, key := range keys {

		value, _ := c.ReadValue(key)

		valueStr := value.(string)
		keyStr := key.(string)

		storage[keyStr] = valueStr

	}
	if err := c.CreateConnection("write", writeToken); err != nil {
		log.Fatal(err)
	}

	for key, value := range storage {
		log.Printf("Making replication for key `%s`", key)
		err := c.writeKeys(key, value)

		if err != nil {
			log.Fatal(err)
		}


	}
	log.Printf("Replication done!")

}

func (d *ReplicaConfig ) CreateConnection( role string, token string) ( error ) {
	var address string

	if role == "reader" {
		address = d.ReaderAddress
	} else {
		address = d.WriteAddress
	}

	client, err := api.NewClient(&api.Config{Address: address})
	if err != nil {
		return err
	}

	client.SetToken(token)
	d.client = client

	log.Printf("Created connection for `%s`", role)
	return nil
}
