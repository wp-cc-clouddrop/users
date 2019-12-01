package adapter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	. "users/internal/types"
)

type UserDBCloud interface {
	Connect() error
	Disconnect() error

	Insert(collection string, obj UserDataI) error
	Update(collection string, id string, obj interface{}) error
	Get(collection string, id string) (map[string]interface{}, error)
	Find(collection string, fieldname string, value string) (map[string]interface{}, error)
	Delete(collection string, id string) error
}

func ReadJSONConfig(filePath string) DBAccess {
	link, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dat, readErr := ioutil.ReadAll(link)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var access DBAccess
	parseErr := json.Unmarshal(dat, &access)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	return access
}
