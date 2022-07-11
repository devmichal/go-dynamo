package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
	"os"
)

func ConnectDynamo() (db *dynamodb.DynamoDB) {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file ")
	}

	region := os.Getenv("DB_REGION")
	host := os.Getenv("DB_HOST")
	keyId := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	token := os.Getenv("AWS_SESSION_TOKEN")

	os.Setenv("AWS_ACCESS_KEY_ID", keyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", accessKey)
	os.Setenv("AWS_SESSION_TOKEN", token)

	creds := credentials.NewEnvCredentials()
	creds.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
		Endpoint:    aws.String(host)})

	if err != nil {
		panic(err)
	}
	return dynamodb.New(sess)
}
