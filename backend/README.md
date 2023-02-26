1. Запустить контейнеры postgres и redis: перейти в папку backend и ввести make rundb
2. Перейти в папку migrations: cd /migrations и выполнить команду go run migrate.go
3. Запустить сервер api: cd /cmd и выполнить команду go run main.go
4. Чтобы снести контейнеры с postgres и redis: перейти в папку backend и выполнить make down