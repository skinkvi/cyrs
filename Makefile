# запуск программы
run:
	docker-compose up --build

# удалить все контейры 
down:
	docker-compose down --rmi all --volumes --remove-orphans

# сгенерировать swagger документациб
swag:
	swag init --generalInfo cmd/cyrs/main.go --output ./docs;

# удалить папку со swagger
deleteswag:
	rm -rf docs

