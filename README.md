# Wildberries-L0

Сервис принимает заказы от брокера сообщений NATS-Streaming в формате JSON, 
полученные данные пишет в БД (PostgreSQL) и в кэш (in memory).
HTTP-сервер выдает данные по id из кэша.
