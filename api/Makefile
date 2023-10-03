startDb:
	cd ./config/docker/ && \
	docker compose up -d

stopDb:
	cd ./config/docker/ && \
	docker compose down

migrateUp:
	migrate -database "mysql://root:dev@tcp(127.0.0.1)/todo_app?multiStatements=true" -source "file://data/migration" up

migrateDown:
	migrate -database "mysql://root:dev@tcp(127.0.0.1)/todo_app?multiStatements=true" -source "file://data/migration" down

migrateFix:
	migrate -database "mysql://root:dev@tcp(127.0.0.1)/todo_app?multiStatements=true" -source "file://data/migration" force $(VERSION)

genTodoMocks:
	mockgen -source=service/todo_service.go -destination=./service/mocks/todo_service_mocks.go;
	mockgen -source=repository/todo_repository.go -destination=./repository/mocks/todo_repository_mocks.go;

.PHONY: startDb stopDb migrateUp migrateDown migrateFix genTodoMocks
