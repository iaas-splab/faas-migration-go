package ibm

import (
	"fmt"
	"github.com/c-mueller/faas-migration-go/core"
	"github.com/cloudant-labs/go-cloudant"
	"github.com/google/uuid"
)

type Obejct map[string]interface{}

const CloudantEndpoint = "https://44f2aef2-a020-443d-b86e-378fe46e5ccd-bluemix.cloudantnosqldb.appdomain.cloud"
const CloudantUserName = "theressidespeneirdistome"
const CloudantPassword = "29a42a1434b189f4586ac238714dd3064d91ca17"

const DatabaseName = "items"

type cloudantItem struct {
	Id      string        `json:"_id"`
	Rev     string        `json:"_rev"`
	Element core.ToDoItem `json:"element"`
}

func NewCloudantRepository() core.Repository {
	repo := &cloudantRepository{}

	return core.Repository(repo)
}

type cloudantRepository struct {
	Client *cloudant.CouchClient
}

func (c *cloudantRepository) Init() error {
	client, err := cloudant.CreateClient(CloudantUserName, CloudantPassword, CloudantEndpoint, 5)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = client.Get(DatabaseName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.Client = client
	return nil
}

func (c *cloudantRepository) Put(item *core.ToDoItem) (*core.ToDoItem, error) {
	item.ID = uuid.New().String()

	e := cloudantItem{
		Id:      item.ID,
		Rev:     "1",
		Element: *item,
	}

	db, err := c.Client.GetOrCreate(DatabaseName)
	if err != nil {
		return nil, err
	}

	_, err = db.Set(&e)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *cloudantRepository) Get(id string) (*core.ToDoItem, error) {
	db, err := c.Client.GetOrCreate(DatabaseName)
	if err != nil {
		return nil, err
	}

	var e cloudantItem
	err = db.Get(id, nil, &e)
	if err != nil {
		return nil, err
	}

	return &e.Element, nil
}

func (c *cloudantRepository) List() ([]core.ToDoItem, error) {
	db, err := c.Client.GetOrCreate(DatabaseName)
	if err != nil {
		return nil, err
	}

	rows, err := db.All(cloudant.NewAllDocsQuery().Build())

	for {
		row, more := <-rows
		if more {
			fmt.Println(row.ID, row.Value.Rev) // prints document 'id' and 'rev'
		} else {
			break
		}
	}

	return make([]core.ToDoItem, 0), nil
}

func (c *cloudantRepository) Delete(id string) error {
	db, err := c.Client.GetOrCreate(DatabaseName)
	if err != nil {
		return err
	}
	return db.Delete(id, "1")
}
