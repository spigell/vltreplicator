package replication


import (
	"fmt"

)

func (d *ReplicaConfig) writeKeys ( key string, value string ) (e error) {

	path := d.WritePath

	writeKeyName := path + "/" + key
	m := make(map[string]interface{})
	m["key"] = value
	writeData := m
	_, err := d.client.Logical().Write(writeKeyName, writeData)
	if err != nil {
		fmt.Println(err)
		return
	}

	return nil
}