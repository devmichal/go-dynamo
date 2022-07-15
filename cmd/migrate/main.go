package migrate

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/urfave/cli/v2"
	"musement.com/commission-service/internal/database"
	"musement.com/commission-service/internal/model"
)

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

	return c.Err()
}
