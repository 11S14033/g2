gen_pb:
	protoc commons/protocs/Numbers.proto --go_out=plugins=grpc:. 

start_anouncement:
	go run services/anouncement/*.go

up_kafka:
	docker-compose -f kafka/docker-compose.yml up -d 

stop_kafka:
	docker-compose -f kafka/docker-compose.yml stop

down_kafka:
	docker-compose -f kafka/docker-compose.yml down
	#docker-compose down --volumes 
	# Down and remove volumes
	#docker-compose down --volumes 

	# Down and remove images
	#docker-compose down --rmi <all|local> 

#create Topic:
create_topic:
	docker exec -it  kafka kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic diantest


#Consumer
consume:
	docker exec -it  kafka kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic diantest --from-beginning

#Producers
produce:
	docker exec -it kafka kafka-console-producer.sh --broker-list localhost:9092 --topic diantest