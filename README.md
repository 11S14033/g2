# G1

## Description
This is a fun project to imporove my knowledge about using kafka with go. This project maintain anouncment processing. This project will produce and consume anouncement model from/to kafka:

Here component in this project:
* Golang 
* gRPC (using libary evans for running using grpc server)
* MongoDB
* Docker
## Prerequisites
See Makefile for:

* Generate file pb
* Create topic
* Up and stop Kafka

## Running App
### Run Anouncement service
After the requirements are met:
```
start_anouncement:
	go run services/anouncement/*.go
```
To start Anouncement service:
```
make start_anouncement
```


## Service
Anouncement Services (rpc delivery)
```
PublishAnouncement
GetAnouncements
ConsumeAndSave
```

You can use reflection like evans to run service at gRPC server. See more detail about evans reflection at https://github.com/ktr0731/evans.
```
evans -p 50051 -r

call PublishAnouncement
call GetAnouncements
call ConsumeAndSave
```


## Tools Used
* All library can be seen in file go.mod

## Next Step
This project is still not finished and still doing:
* Unit testing
* Solve some isue
* Improve kafka (instance,partition, group, etc)
* Auto sacling using kubernates

