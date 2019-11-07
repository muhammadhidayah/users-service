build:
	protoc -I. --go_out=. --micro_out=. \
		proto/users/users.proto