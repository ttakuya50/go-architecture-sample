.PHONY: run
run:
	go run api/*.go

.PHONY: build
build:
	go build -v api/main.go

.PHONY: test
test:
	go test -shuffle=on -race ./...

.PHONY: docker_up
docker_up:
	docker-compose -f docker-compose.yml up --build -d

.PHONY: sqlboiler
sqlboiler:
	sqlboiler mysql -o api/infra/mysql -p mysql -d -c sqlboiler.toml
	goimports -local github.com/ttakuya50/go-architecture-sample/api -w ./

.PHONY: mock
mock:
	go generate -x -run="mockgen.*" ./...

.PHONY: fmt
fmt:
	goimports -local github.com/ttakuya50/go-architecture-sample/api -w ./

.PHONY: install
install:
	go install github.com/golang/mock/mockgen@v1.6.0
	brew install openapi-generator
	docker pull swaggerapi/swagger-editor

.PHONY: swagger-editor
swagger-editor:
	docker run -d -p 80:8090 swaggerapi/swagger-editor
	# http://localhost にアクセス

.PHONY:openapi-generator
openapi-generator:
	openapi-generator generate -i swagger/openapi.yml -g go-server --package-name handler
	mv ./go/*.go ./api/handler/
	rm -r ./go && rm -r .openapi-generator

