# go-container-demo

A small project that I am using to learn the [Go](https://github.com/golang/go) programming language. I intend to build a container that consumes a message from a queue, then writes a resulting file to AWS S3.

The aim will be to listen for a transaction message on a queue, and write a receipt file to S3 with details from the transaction message.

Currently the integration test requires [LocalStack](https://github.com/localstack/localstack) to be running.
