package adapter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	. "users/internal/types"
)

type MongoDBCloud interface {
	Connect() error
	Disconnect() error

	Insert(collection string, obj interface{}) error
	Update(collection string, id string, obj interface{}) error
	Get(collection string, id string) (interface{}, error)
	Delete(collection string, id string) error
}

func ReadJSONConfig(filePath string) MongoAccess {
	link, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dat, readErr := ioutil.ReadAll(link)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var access MongoAccess
	parseErr := json.Unmarshal(dat, &access)

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	return access
}
