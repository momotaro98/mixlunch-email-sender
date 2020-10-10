# mixlunch-email-notification

## Run

Run [mixlunch-service-api](https://github.com/momotaro98/mixlunch-service-api)

```
$ cd mixlunch-service-api
$ make docker-run
```

Run email batch

```
$ cd mixlunch-email-notification
# Default DATE=2018-10-31
$ make run
# Or specify date
$ make run DATE=2019-05-31
```

### (Option) Run with Docker

```
make docker-run MAILGUN_PRIVATE_API_KEY=$MAILGUN_PRIVATE_API_KEY
```

## Development

Let's check `Makefile`

### How to update gRPC/protocol buffer

```
$ cd mixlunch-email-notification
$ protoc -I pb/ pb/mixlunch.proto --go_out=plugins=grpc:pb
```

