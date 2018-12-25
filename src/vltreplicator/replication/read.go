package replication

import (

	"log"
	"strings"
)

func (d *ReplicaConfig) ReadKeys() ([]interface{}, error ){
	path := d.ReaderPath
	if d.ReaderVersion == 2 {

		slice := strings.Split(path, "/")
		slice[0] += "/metadata"
		path = strings.Join(slice, "/")

	}
	listPath := path
	secrets, err := d.client.Logical().List(listPath)
	if err != nil {
		return nil, err
	}

	data := secrets.Data
	keys := data["keys"].([]interface{})

	return keys, nil


}

func (d *ReplicaConfig) ReadValue(key interface{}) (interface{}, error ) {
	path := d.ReaderPath
	if d.ReaderVersion == 2 {

		slice := strings.Split(path, "/")
		slice[0] += "/data"
		path = strings.Join(slice, "/")

	}
	keyName := path + "/" + key.(string)
	secret, err := d.client.Logical().Read(keyName)
	if err != nil {
		log.Print(err)
		return "", err
	}
	data := secret.Data["data"].(map[string]interface{})
	return data["value"], nil

}
