# Описание

CRUD-приложение для управления списком задач (Task Manager), написанное на Go.

![image](https://user-images.githubusercontent.com/102806309/227769756-891816d7-de20-4bd5-bec4-188ff0c088c3.png)
![image](https://user-images.githubusercontent.com/102806309/227769771-a8974ac1-3283-4a55-b247-15009cfc6c18.png)


## Установка

1. Установить используя команду

```bash
git clone https://github.com/exdk/go-task-manager
```
2. Переименовать файл example.env в .env и прописать в нем настройки подключения к базе данных.  
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
