# mongo restapi server
web server for store events in mongo-db

### build
выполнить make в директории проекта - соберется сервис и поднимется контейнер с mongo-db на порту 2717
(требуется установленный docker-compose)

### run
./server - запустит сервис на порту 8080

### usage
* **POST localhost:8080/v1/start {"type":"e0"}** - создание события e0
* **POST localhost:8080/v1/finish {"type":"e0"}** - завершение события e0

### math
в директории /cmd/math выполнить go run main.go, принцип решения задачи:
- встреча возможна если расстояние с течением времени сокращается
- если изначально находились в одной точке, то встреча уже произошла
