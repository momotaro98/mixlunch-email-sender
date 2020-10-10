NAME=mixlunch-email-notification
TAG=0.1
TARGET_DATE=2018-10-31T20:01:23Z # Variable

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_DIR=build
BINARY_NAME=mixlunch-email-notification

.PHONY: all
all: build test

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: run
run:
	mkdir -p ./$(BINARY_DIR)/templates
	cp ./templates/body.html.tpl ./$(BINARY_DIR)/templates/body.html.tpl
	make build
	cd $(BINARY_DIR) && ./$(BINARY_NAME) -target-date ${TARGET_DATE}
	make clean

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

.PHONY: docker-run
docker-run:
	docker build -t mixlunch-email-notification -t ${NAME}:${TAG} . \
		&& docker run --rm --net=mixlunch-service-api_grpc-net \
		-e MAIL_DOMAIN_NAME=${MAIL_DOMAIN_NAME} \
		-e MAILGUN_PRIVATE_API_KEY=${MAILGUN_PRIVATE_API_KEY} \
		-e PARTYSERVICE_API_ADDRESS=mixlunch-service-api_grpc-server \
		-e PARTYSERVICE_API_PORT=${PARTYSERVICE_API_PORT} \
		-e TARGET_DATE=${TARGET_DATE} \
		${NAME}:${TAG}
