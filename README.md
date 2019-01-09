# go-container-demo

A small project that I am using to learn the [Go](https://github.com/golang/go) programming language. I intend to build a container that consumes a message from a queue, then writes a resulting file to AWS S3.

The aim will be to listen for a transaction message on a queue, and write a receipt file to S3 with details from the transaction message.

The integration test starts a [LocalStack](https://github.com/localstack/localstack) container. Please ensure that LocalStack is not running before starting the test.

### Prerequesites

- [Go 1.11](https://golang.org/doc/go1.11) for module support
- [Docker](https://www.docker.com) installed and daemon running

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
