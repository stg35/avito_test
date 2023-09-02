# avito_test
## Запуск
`docker compose up -d`
## Swagger
Swagger находиться на localhost:8000/swagger/index.html
## Подробнее
Весь минимум сделал, а также реализовал добавление сегментов пользователю с TTL с помощью системы очередей asynq. Тесты, к сожалению, не успел написать.
## Описание роутов
1. Метод создания сегмента. Принимает название сегмента. - http://localhost:8000/api/segment [post]
Пример запроса:
`{
    "name": "AVITO_VOICE_MESSAGES"
 }`
2. Метод удаления сегмента. Принимает id сегмента. - http://localhost:8000/api/segment/{id} [delete]
3. Метод добавления пользователя в сегмент. Принимает список названий сегментов, которые нужно добавить пользователю, id пользователя и TTL. - http://localhost:8000/api/user/addSegments [patch]
Пример запроса:
`{
    "id": 2,
	"segments": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"],
	"ttl": 15
}`
Чтобы добавить пользователя в сегмент без ttl, нужно передать ttl равным 0.
4. Метод удаления пользователя из сегмента. Принимает список названий сегментов, которые нужно удалить у пользователя и id пользователя. - http://localhost:8000/api/user/deleteSegments [patch]
Пример запроса:
`{
    "id": 2,
	"segments": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"]
}`
5. Метод получения активных сегментов пользователя. Принимает на вход id пользователя. - http://localhost:8000/api/user/showSegments/{id} [get]
6. Метод создания пользователя. Принимает имя пользователя. - http://localhost:8000/api/user [post]
Пример запроса:
`{
    "username": "Alex"
 }`