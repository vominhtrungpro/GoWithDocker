# ==============================================================================
# Main
run: 
	go run main.go

docker-up:
	@docker compose up -d

create-topic:
	@docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic quickstart

write-message:
	@docker exec --interactive --tty broker kafka-console-producer --bootstrap-server broker:9092 --topic quickstart

read-message:
	@docker exec --interactive --tty broker kafka-console-consumer --bootstrap-server broker:9092 --topic quickstart --from-beginning