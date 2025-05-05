run:
	docker compose up --build
proto-user:
	protoc --go_out=proto/gen --go-grpc_out=proto/gen proto/user/user_service.proto
proto-recipe:
	protoc --go_out=proto/gen --go-grpc_out=proto/gen proto/recipe/recipe_service.proto
proto-review:
	protoc --go_out=proto/gen --go-grpc_out=proto/gen proto/review/review_service.proto
proto: proto-user proto-recipe proto-review