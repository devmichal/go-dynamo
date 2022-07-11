package migrate

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/urfave/cli/v2"
	"log"
	"musement.com/commission-service/internal/database"
	"musement.com/commission-service/internal/model"
	"os"
	"strconv"
	"time"
)

var RegionName = "us-east-1"
var TableName = "commission_super_pro"

/*func init() {
	connectDynamo()
}*/

// connectDynamo returns a dynamoDB client
func connectDynamos() (db *dynamodb.DynamoDB) {

	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")

	creds := credentials.NewEnvCredentials()
	creds.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      &RegionName,
		Credentials: creds,
		Endpoint:    aws.String("http://localhost:8000")})

	if err != nil {
		panic(err)
	}
	return dynamodb.New(sess)
}

// Create a new table
func createTable() {
	dbSvc := connectDynamos()
	_, err := dbSvc.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   &TableName,
	})

	if err != nil {
		panic(err)
	}
}

func GetItem(id string) (commission model.Commission, err error) {
	dbSvc := connectDynamos()
	result, err := dbSvc.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		TableName: &TableName,
	})

	if err != nil {
		return commission, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &commission)

	return commission, err

}

func PutItem(commission model.Commission) error {
	dbSvc := connectDynamos()
	_, err := dbSvc.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(commission.Id),
			},
			"Name": {
				S: aws.String(commission.Name),
			},
			"Status": {
				N: aws.String(strconv.Itoa(commission.Status)),
			},
			"Price": {
				S: aws.String(strconv.FormatFloat(commission.Price, 'E', 2, 32)),
			},
			"CreatedAt": {
				S: aws.String(time.Now().Format("2006:04:02 15:04")),
			},
		},
		TableName: &TableName,
	})

	return err
}

func PrintAndGet(id string) {
	p, err := GetItem(id)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v", p)
}

func Command(c *cli.Context) error {

	connect := database.ConnectDynamo()

	_, err := connect.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(model.TableName),
	})

	if err != nil {
		panic(err)
	}

	/*	fmt.Println("_______________start_________")

		id := uuid.New()
		fmt.Println(id.String())
		//createTable()
		commission := commission.Commission{
			Id:     id.String(),
			Name:   "Hobbit un expected journey",
			Status: 3,
			Price:  5.11,
		}
		err := PutItem(commission)

		if err != nil {
			panic(err)
		}

		PrintAndGet("44779514-6337-44bf-be91-56c83825198a")

		fmt.Println("_______________finish no error_________")*/

	return c.Err()
}
