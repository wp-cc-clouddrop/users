package gcpFirestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"users/internal/types"
)

type GCPFirestore struct {
	client *firestore.Client
	ctx    context.Context
}

func (fs *GCPFirestore) Connect() error {
	// Sets your Google Cloud Platform project ID.
	projectID := "clouddrop-gcp"

	// Get a Firestore client.
	fs.ctx = context.Background()
	//only for local testing
	//jsonPath := "/home/dennis/Downloads/clouddrop-gcp-59a95369a26c.json"
	//client, err := firestore.NewClient(fs.ctx, projectID, option.WithCredentialsFile(jsonPath))
	client, err := firestore.NewClient(fs.ctx, projectID)
	fs.client = client
	if err != nil {
		return errors.New("connection to gcp firestore failed")
	}
	return nil
}

func (fs *GCPFirestore) Disconnect() error {
	fs.client.Close()
	return nil
}

func (fs *GCPFirestore) Insert(collection string, obj types.UserDataI) error {
	_, err := fs.client.Collection(collection).Doc(obj.Id()).Set(fs.ctx, obj)
	return err
}

func (fs *GCPFirestore) Update(collection string, id string, obj interface{}) error {

	return nil
}

func (fs *GCPFirestore) Get(collection string, id string) (map[string]interface{}, error) {
	doc, err := fs.client.Collection(collection).Doc(id).Get(fs.ctx)
	if err != nil {
		return nil, err
	} else {
		return doc.Data(), nil
	}
}

func (fs *GCPFirestore) Find(collection string, fieldname string, value string) (map[string]interface{}, error) {
	doc, err := fs.client.Collection(collection).Where(fieldname, "==", value).Documents(fs.ctx).Next()
	if err != nil {
		return nil, err
	} else {
		return doc.Data(), nil
	}
}

func (fs *GCPFirestore) Delete(collection string, id string) error {
	fs.client.Collection(collection).Doc(id).Delete(fs.ctx)
	return nil
}
