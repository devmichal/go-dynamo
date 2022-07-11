package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"musement.com/commission-service/internal/model"
	"strconv"
)

type CommissionRepository struct {
	db *dynamodb.DynamoDB
}

func NewRuleRepository(db *dynamodb.DynamoDB) *CommissionRepository {
	return &CommissionRepository{
		db: db,
	}
}

func (h *CommissionRepository) CreateCommission(commission *model.Commission) error {
	_, err := h.db.PutItem(&dynamodb.PutItemInput{
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
				S: aws.String(commission.CreatedAt),
			},
		},
		TableName: aws.String(model.TableName),
	})

	return err
}

func (h *CommissionRepository) GetById(id string) *model.Commission {
	result, err := h.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(model.TableName),
	})

	if err != nil {
		panic(err)
	}

	commission := &model.Commission{}
	_ = dynamodbattribute.UnmarshalMap(result.Item, commission)

	return commission
}
