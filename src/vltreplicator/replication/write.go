package replication


import (
)

func (d *ReplicaConfig) writeKeys ( key string, value string ) ( error) {

	path := d.WritePath

	writeKeyName := path + "/" + key
	m := make(map[string]interface{})
	m["value"] = value
	writeData := m
	_, err := d.client.Logical().Write(writeKeyName, writeData)
	if err != nil {
		return err
	}

	return nil
}
