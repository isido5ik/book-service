markdown
Copy code
# Book Service

Привет! Это репозиторий для проекта "Book Service". Этот проект представляет собой простую веб-службу для управления книгами.

## В данном проекте использованы:

- **Rest Api**: Для построения веб-службы, которая обеспечивает доступ к данным о книгах через HTTP протокол.
- **PostgreSQL**: В качестве реляционной базы данных для хранения информации о книгах, авторах и других сущностях.
- **Sqlx**: Библиотека для работы с базой данных PostgreSQL в Go, обеспечивающая безопасные SQL-запросы и маппинг результатов запросов на структуры.
- **Redis**: Кэш-хранилище для улучшения производительности при доступе к данным, таким как информация о книгах.
- **Gin**: Фреймворк для создания веб-приложений на языке Go, облегчающий обработку HTTP запросов, роутинг и многое другое.

## Установка

1. Клонируйте репозиторий на ваш компьютер:

```bash
git clone https://github.com/isido5ik/book-service.git
```
2. Перейдите в каталог проекта:
```bash
cd book-service
```
3. Установите зависимости Go:
```bash
go mod tidy
```

4. Запустите приложение:
```bash
go run main.go
```

Перейдите по адресу http://localhost:8080 в вашем браузере для доступа к приложению.







