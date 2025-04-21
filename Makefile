run:
	docker compose up --build
proto-user:
	protoc --go_out=proto/gen --go-grpc_out=proto/gen proto/user/user_service.proto