buildapi:
	docker-compose build main-app

rundb:
	docker-compose up -d postgres redis

runapi:
	docker-compose up main-app

down:
	docker-compose down && docker rmi todos-api-image:1.0.0

restart:
	docker-compose down && docker rmi todos-api-image:1.0.0 && docker-compose up main-app

protoapigtwauth:
	protoc services/api-gateway/grpc-svc-clients/auth/pb/auth.proto --go_out=plugins=grpc:.

protoauth:
	protoc services/auth-svc/internal/pb/auth.proto --go_out=plugins=grpc:.