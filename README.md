# Task tracker API

Небольшое API для работы с таск трекером
В написании использовались Fiber, PostgreSQL и Swaggo
Swagger-документация доступна по маршруту "http://localhost:3000/swagger/index.html"


### Для корректной работы необходимо создать в корне проекта файл .env со следующими переменными окружения:

```
DB_NAME=your_database_name        # Название базы данных
DB_USER=your_username             # Имя пользователя PostgreSQL
DB_PASSWORD=your_password         # Пароль пользователя PostgreSQL
DB_HOST=localhost                 # Адрес сервера базы данных
DB_PORT=5432                      # Порт подключения к PostgreSQL
```