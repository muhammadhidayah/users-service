build:
	protoc -I. --go_out=. --micro_out=. \
		proto/users/users.proto	
	docker build -t user-service .
run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		user-service
runlocal:
	DB_HOST=localhost DB_USER=postgres DB_PASSWORD=docker DB_NAME=inact_mini \
	go run *.go
