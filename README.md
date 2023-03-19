# Описание

CRUD-приложение для управления списком задач (Task Manager), написанное на Go.

## Установка

1. Установить используя команду

```bash
git clone https://github.com/exdk/go-task-manager
```
2. Переименовать файл .env_example в .env и прописать в нем настройки подключения к базе данных.  
3. Выполнить миграции:
```bash
migrate -path db/migrations -database "mysql://DBUSER:DBPASS@tcp(DBHOST:DBPORT)/DBNAME" -verbose up
```  
Например:
```bash
migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/go-tm" -verbose up
```  

## Зависимости

Проект написан с использованием таких пакетов как:  
[golang-migrate/migrate](https://github.com/golang-migrate/migrate)  
[go-playground/validator](https://github.com/go-playground/validator/)  
[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)  
[joho/godotenv](https://github.com/joho/godotenv)  
[gorilla/mux](https://github.com/gorilla/mux)

## License

[MIT](https://choosealicense.com/licenses/mit/)
