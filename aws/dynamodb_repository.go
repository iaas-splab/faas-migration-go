package aws

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/iaas-splab/faas-migration-go/core"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"os"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

const TableNameEnvName = "DYNAMODB_TABLE"

var TableName = os.Getenv(TableNameEnvName)

type DynamoRepo struct {
	TableName string

	session *session.Session
	db      *dynamo.DB
	table   *dynamo.Table
}

func (r *DynamoRepo) Put(i *core.ToDoItem) (*core.ToDoItem, error) {
	if len(i.ID) == 0 {
		i.ID = uuid.New().String()
	}
	err := r.table.Put(*i).Run()
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (r *DynamoRepo) Get(k string) (*core.ToDoItem, error) {
	var item core.ToDoItem
	err := r.table.Get("ID", k).One(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *DynamoRepo) List() ([]core.ToDoItem, error) {
	items := make([]core.ToDoItem, 0)

	err := r.table.Scan().All(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *DynamoRepo) Delete(k string) error {
	return r.table.Delete("ID", k).Run()
}

func NewDynamoRepo() core.Repository {
	return &DynamoRepo{
		TableName: TableName,
	}
}

func (r *DynamoRepo) Init() error {
	r.session = session.Must(session.NewSession())

	r.db = dynamo.New(r.session)

	table := r.db.Table(r.TableName)
	r.table = &table

	return nil
}
