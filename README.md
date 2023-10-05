# Тестовое задание Medods "Сервис аутентификации"
[![forthebadge](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNDYuMTQwNjM0NTM2NzQzMTYiIGhlaWdodD0iMzUiIHZpZXdCb3g9IjAgMCAxNDYuMTQwNjM0NTM2NzQzMTYgMzUiPjxyZWN0IHdpZHRoPSIxMDMuNDg0MzgyNjI5Mzk0NTMiIGhlaWdodD0iMzUiIGZpbGw9IiMxZDdlYTAiIGRhdGEtZGFya3JlYWRlci1pbmxpbmUtZmlsbD0iIiBzdHlsZT0iLS1kYXJrcmVhZGVyLWlubGluZS1maWxsOiAjMTc2NTgwOyIvPjxyZWN0IHg9IjEwMy40ODQzODI2MjkzOTQ1MyIgd2lkdGg9IjQyLjY1NjI1MTkwNzM0ODYzIiBoZWlnaHQ9IjM1IiBmaWxsPSIjMzg5QUQ1IiBkYXRhLWRhcmtyZWFkZXItaW5saW5lLWZpbGw9IiIgc3R5bGU9Ii0tZGFya3JlYWRlci1pbmxpbmUtZmlsbDogIzIyNzJhMjsiLz48dGV4dCB4PSI1MS43NDIxOTEzMTQ2OTcyNjYiIHk9IjE3LjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInUm9ib3RvJywgc2Fucy1zZXJpZiIgZmlsbD0iI0ZGRkZGRiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgYWxpZ25tZW50LWJhc2VsaW5lPSJtaWRkbGUiIGxldHRlci1zcGFjaW5nPSIyIiBkYXRhLWRhcmtyZWFkZXItaW5saW5lLWZpbGw9IiIgc3R5bGU9Ii0tZGFya3JlYWRlci1pbmxpbmUtZmlsbDogI2U4ZTZlMzsiPk1BREUgV0lUSDwvdGV4dD48dGV4dCB4PSIxMjQuODEyNTA4NTgzMDY4ODUiIHk9IjE3LjUiIGZvbnQtc2l6ZT0iMTIiIGZvbnQtZmFtaWx5PSInTW9udHNlcnJhdCcsIHNhbnMtc2VyaWYiIGZpbGw9IiNGRkZGRkYiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtd2VpZ2h0PSI1MDAiIGFsaWdubWVudC1iYXNlbGluZT0ibWlkZGxlIiBsZXR0ZXItc3BhY2luZz0iMiIgZGF0YS1kYXJrcmVhZGVyLWlubGluZS1maWxsPSIiIHN0eWxlPSItLWRhcmtyZWFkZXItaW5saW5lLWZpbGw6ICNlOGU2ZTM7Ij5HTzwvdGV4dD48L3N2Zz4=)](https://forthebadge.com)
[![forthebadge](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMzQuMzQzNzU4NTgzMDY4ODUiIGhlaWdodD0iMzUiIHZpZXdCb3g9IjAgMCAxMzQuMzQzNzU4NTgzMDY4ODUgMzUiPjxyZWN0IHdpZHRoPSIxMDUuMjE4NzU3NjI5Mzk0NTMiIGhlaWdodD0iMzUiIGZpbGw9IiNmNTliMzIiIGRhdGEtZGFya3JlYWRlci1pbmxpbmUtZmlsbD0iIiBzdHlsZT0iLS1kYXJrcmVhZGVyLWlubGluZS1maWxsOiAjYWM2MDA4OyIvPjxyZWN0IHg9IjEwNS4yMTg3NTc2MjkzOTQ1MyIgd2lkdGg9IjI5LjEyNTAwMDk1MzY3NDMxNiIgaGVpZ2h0PSIzNSIgZmlsbD0iI2ZjYTIwMyIgZGF0YS1kYXJrcmVhZGVyLWlubGluZS1maWxsPSIiIHN0eWxlPSItLWRhcmtyZWFkZXItaW5saW5lLWZpbGw6ICNjYTgyMDI7Ii8+PHRleHQgeD0iNTIuNjA5Mzc4ODE0Njk3MjY2IiB5PSIxNy41IiBmb250LXNpemU9IjEyIiBmb250LWZhbWlseT0iJ1JvYm90bycsIHNhbnMtc2VyaWYiIGZpbGw9IiNGRkZGRkYiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGFsaWdubWVudC1iYXNlbGluZT0ibWlkZGxlIiBsZXR0ZXItc3BhY2luZz0iMiIgZGF0YS1kYXJrcmVhZGVyLWlubGluZS1maWxsPSIiIHN0eWxlPSItLWRhcmtyZWFkZXItaW5saW5lLWZpbGw6ICNlOGU2ZTM7Ij5CVUlMRCBXSVRIPC90ZXh0Pjx0ZXh0IHg9IjExOS43ODEyNTgxMDYyMzE2OSIgeT0iMTcuNSIgZm9udC1zaXplPSIxMiIgZm9udC1mYW1pbHk9IidNb250c2VycmF0Jywgc2Fucy1zZXJpZiIgZmlsbD0iI0ZGRkZGRiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZm9udC13ZWlnaHQ9IjUwMCIgYWxpZ25tZW50LWJhc2VsaW5lPSJtaWRkbGUiIGxldHRlci1zcGFjaW5nPSIyIiBkYXRhLWRhcmtyZWFkZXItaW5saW5lLWZpbGw9IiIgc3R5bGU9Ii0tZGFya3JlYWRlci1pbmxpbmUtZmlsbDogI2U4ZTZlMzsiPuKZpTwvdGV4dD48L3N2Zz4=)](https://forthebadge.com)

Микросервис для генерации токенов(Access и Refresh) по уникальному идентификатору пользовотеля(GUID) и их обновления.

Используемые технологии:
- **Go**
- **JWT**
- **MongoDB**
- **Fiber** (веб фреймворк)
- **Swagger** (для документации API)
- **Docker** (для запуска сервиса)

Сервис был написан с Clean Architecture, что позволяет легко расширять функционал сервиса и тестировать его. Также был реализован Graceful Shutdown для корректного завершения работы сервиса.

# Getting start
Опциональлно можно изменить `.env` файл, есть такие параметры:
- CONN_URI - uri для подключения к mongoDB. По умолчанию: `mongodb://mongo:27017`.
- MG_DB - имя базы данных mongoDB. По умолчанию: `test_task`.
- SRV_PORT - Порт fiber сервера. По умолчанию: `3003`.
- JWT_SIGN_KEY - ключ для подписи JWT. Обязателен.
- JWT_TOKEN_TTL - время жизни Access токена. Обязателен. В файле задано 2 минуты.
- REFRESH_TOKEN_TTL - время жизни Rerfresh токена. Обязателен. В файле задано 10 минут.
- HASH_COST - число для хеширования в bcrypt. По умолчанию: 12.

# Usage
#### Запуск сервиса
Команда `make compose-up` - запускает сервис
#### Просмотр документации
Документацию можно посмотреть после запуска сервиса, перейдя по адресу - `http://localhost:3003/swagger/` с портом 3003 по дефолту
#### Остановка сервиса
Команда `make compose-down` - останавливает сервис
# Examples
#### Генерация токенов
В query кладем guid, на выходе получаем accessToken в виде объекта и в куках будет лежать refreshToken с max-age равнная RefreshTokenTTL в .env
**Пример запроса:**
```bash
curl --location 'http://127.0.0.1:3003/auth/generateTokens?guid=db654f5d2c354b118cccaaa1fb6cc81a'
```

**Пример ответа:**
```json
{
	"auth_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY1MDg1MjkuMTkzMDkxMiwiaWF0IjoxNjk2NTA4NDA5LjE5MzA5MywiZ3VpZCI6ImRiNjU0ZjVkMmMzNTRiMTE4Y2NjYWFhMWZiNmNjODFhIn0.F2O6E39oKTO9sKZ3XN7n1KRuosa1E_K6y89GAxlWydYUGsfxGif0M4Iuib9iiVQCwAlKlP2hGc95hBLlTOe2ww"
}
```
#### Рефреш
После выполнения запроса обновятся токены: в ответе будет новый access, в куках новый рефреш. В случае если рефреш не валидный или его не будет в базе, вернет ошибку 401.
**Пример запроса с валидным токеном:**
```bash
curl --location 'http://127.0.0.1:3003/auth/refresh' \
--header 'Cookie: RefreshToken=JDJhJDEyJDRJaXBtdDQ3UjhMUFB2amU0UjV2ZC5BNVJ2MWNuQ2NWdk5oalNoRTQxMWl3T2wvY2ZORGFp'
```
в куках лежит рефреш токен сгенерированный ранее

**Пример ответа:**
```json
{
	"auth_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY1MDg3ODkuNzAxNDA4LCJpYXQiOjE2OTY1MDg2NjkuNzAxNDEsImd1aWQiOiJkYjY1NGY1ZDJjMzU0YjExOGNjY2FhYTFmYjZjYzgxYSJ9.uJgA8PGaEZHHHbNFac_Yu0DB32Td57UnGYsBA2MbDxgk9MLC6Ib8WIuTI71FO-pMILsZuHgW_nsmFsYpbF0PQA"
}
```

**Пример запроса с не валидным токеном:**
```bash
curl --location 'http://127.0.0.1:3003/auth/refresh' \
--header 'Cookie: RefreshToken=JDJhJDEyJDF4cEwvaUdnSWkvMHRqZXExRldTN09lLkJKb3ZRUXM2SjNqVEduUXRhQm1pUS9BdVZnWi9T'
```
в куках лежит измененный рефреш токен

**Пример ответа:**
```json
{
	"error" : "refresh token not valid"
}
```