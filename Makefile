SRC_PATH:= ${PWD}

## Developing jobs
prepare:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/golang/mock/mockgen@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install github.com/daixiang0/gci@latest
	@go install github.com/rubenv/sql-migrate/...@latest
	@curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${shell go env GOPATH}/bin v1.50.1


migrate-create:
	@$(eval NAME := $(shell read -p "Enter new file name: " v && echo $$v))
	$(eval CMD:= $*)
	cd db;\
	sql-migrate new ${NAME}

# [up,down]
migrate-%:
	$(eval CMD:= $*)
	cd db;\
	sql-migrate $(CMD) -config=dbconfig.yml;


gen:
	## Go generate
	go generate ./...
	## Swagger generate
	@swag init -g app/delivery/http/routes/routes.go -o app/delivery/http/docs --exclude pkg,db,deployment,scripts,vendor
	@./scripts/gci.sh


mod:
	@go mod tidy
	@go mod vendor

gci:
	scripts/gci.sh

up:
	@go run ${SRC_PATH}/cmd/server/...	

test:
	@go test -cover ./...
