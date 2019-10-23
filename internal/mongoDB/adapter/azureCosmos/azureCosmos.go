package azureCosmos

import (
	"encoding/json"
	. "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"github.com/globalsign/mgo"
	"log"
	"users/internal/mongoDB/adapter"
)

const filePath string = "azureConfig.json"

type AzureCosmos struct {
	session  *mgo.Session
	database *mgo.Database
}

func (azure *AzureCosmos) Connect() error {
	access := adapter.ReadJSONConfig(filePath)

	s, err := NewMongoDBClientWithCredentials(
		access.User,
		access.Password,
		access.Host,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = s.Ping()

	if err != nil {
		log.Fatal(err)
	}
	azure.session = s
	azure.database = s.DB(access.Database)
	return nil
}

func (azure *AzureCosmos) Disconnect() error {
	azure.database.Session.Close()
	return nil
}

func (azure *AzureCosmos) Insert(collection string, obj interface{}) error {
	insertErr := azure.database.C(collection).Insert(obj)
	return insertErr
}

func (azure *AzureCosmos) Update(collection string, id string, obj interface{}) error {
	updErr := azure.database.C(collection).UpdateId(id, obj)
	return updErr
}

func (azure *AzureCosmos) Get(collection string, id string) ([]byte, error) {
	var data *interface{}
	getErr := azure.database.C(collection).FindId(id).One(&data)
	if getErr != nil { // no entrie found
		return nil, getErr
	}
	bytes, parseErr := json.Marshal(data)
	return bytes, parseErr
}

func (azure *AzureCosmos) Delete(collection string, id string) error {
	panic("implement me")
}
