SWAGGER_OUT_DIR ?= api/swagger


hot:
	nodemon --exec go run -mod=vendor cmd/main.go server $* --signal SIGTERM

local-setup:
	cat local/*.env > .env && cp docker-compose.override.yml.dist docker-compose.override.yml

vendor:
	GOPRIVATE="gitlab.com/commission" go mod vendor

lint-go:
	@golangci-lint run ./...

lint-yaml:
	@yamllint .


lint: lint-go lint-yaml

clean:
	rm -rf bin/

run-%:
	go run -mod=vendor cmd/main.go $*

# swagger
swagger:
	swag init \
		--parseDependency \
		--output api/swagger \
		--dir cmd/server/ \
	&& rm api/swagger/swagger.yaml

test:
	grc go test -v -failfast -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
