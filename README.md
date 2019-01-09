# go-container-demo

A small project that I am using to learn the [Go](https://github.com/golang/go) programming language. 

The aim will be to listen for a transaction message on a queue, and upload a receipt file to S3 with details from the transaction message.

### Prerequesites

- [Go 1.11](https://golang.org/doc/go1.11) for module support
- [Docker](https://www.docker.com) installed and daemon running
- [LocalStack](https://github.com/localstack/localstack) NOT started. LocalStack is started as a docker container with default ports.

### Current state

- Starts a LocalStack container
- Writing file to S3
- Running integration test to test S3 uploader using LocalStack
- Tested on macOS 10.13

### To Do

- Customer and transaction objects
- Listening on queue (SQS, probably) for a transaction message
- Joining message listener to S3 uploader
- Containerise
- Integration test everything with LocalStack
