package calculate

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"musement.com/commission-service/internal/repository"
	"musement.com/commission-service/internal/validation"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {

	tables := struct {
		name      string
		inputJSON func() string
	}{
		inputJSON: func() string {
			return ReadJSONFile("../../fixture/payload/postcommission/post_commission.json")
		},
	}

	e := echo.New()
	e.Validator = validation.NewValidator()

	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/calculate",
		strings.NewReader(tables.inputJSON()),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	region := os.Getenv("us-east-1")
	host := os.Getenv("commission")
	keyId := os.Getenv("dummy1")
	accessKey := os.Getenv("dummy2")
	token := os.Getenv("dummy3")

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

	commissionRepository := repository.NewRuleRepository(dynamodb.New(sess))

	postCalculator := New(*commissionRepository)
	postCalculator.Handler(c)

	// assert
	assert.Equal(t, http.StatusOK, rec.Code)
}

func ReadJSONFile(fileName string) string {
	var data []byte
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err.Error())
	}

	buff := new(bytes.Buffer)

	if err := json.Compact(buff, data); err != nil {
		panic(err)
	}

	return buff.String()
}
