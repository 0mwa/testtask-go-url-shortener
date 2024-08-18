<div align="center">
    <b>Go URL Shortener Service</b>
</div>


## ТЗ

<details close>
Требуется создать микросервис сокращения url. Длина сокращенного URL-адреса должна быть как 
можно короче. Сокращенный URL может содержать цифры (0-9) и буквы (a-z, A-Z) 

Эндпоинты: 

POST http://localhost:8080/ 

Request: (body): http://cjdr17afeihmk.biz/123/kdni9/z9d112423421 

Response: http://localhost:8080/qtj5opu 

GET 

Request (url query): http://localhost:8080/qtj5opu 

Response (body): http://cjdr17afeihmk.biz/123/kdni9/z9d112423421 

Микросервис должен уметь хранить информацию в памяти и в postgres в зависимости от флага 
запуска -d 
</details>

## Запуск

Клонировать репозиторий
```shell
git clone https://github.com/0mwa/testtask-go-url-shortener.git
cd testtask-go-url-shortener
```

### Вариант с использованием базы данных
```shell
docker-compose up -d 
```

При первом использовании запустить скрипт, который создаст таблицу в Posrgres'e 
```shell
sh initDataBase.sh
```

Запустить приложение
```shell
go run cmd/main.go -d
```

### Вариант с использованием оперативной памяти
```shell
go run cmd/main.go
```

## Взаимодействие

### Создание короткой ссылки
Запрос:
```shell
curl -X POST http://localhost:8080 -d '{"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}'
```
Ответ:
```JSON
{
  "url":"http://localhost:8080/85cb0c"
}
```

### Получение оригинальной ссылки 
Запрос:
```shell
curl -X GET http://localhost:8080/85cb0c
```
Ответ:
```JSON
{
  "url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"
}
```

> [!NOTE]
> 1) Так совпало, что я уже делал похожий пет-проект с использованием Redis для практики, поэтому за основу взял его. 
> 
> 2) Для общения сервера решил использовать полноценный JSON.


