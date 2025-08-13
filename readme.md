# Time Tracker

## Как запустить

- Установить нужные пакеты `go mod tidy`

- Все переменные окружения устанавливаются в `./configs/.env`
файле (обязательно указать все)

- Запустить:
    1. Run - `go run ./cmd/main.go`
    2. Build - `go build ...`
    3. Docker - `docker run ...`

## Примечания

- `./docs` - свагер документация

- `./migrations` - файлы миграции бд

## Техническое Задание

<b>1. Выставить REST методы</b>

    1.Получение данных пользователей:
        - Фильтрация по всем полям.
        - Пагинация.
    2. Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
    3. Начать отсчет времени по задаче для пользователя
    4. Закончить отсчет времени по задаче для пользователя
    5. Удаление пользователя
    6. Изменение данных пользователя
    7. Добавление нового пользователя в формате:

```json
{
	"passportNumber": "1234 567890" // серия и номер паспорта пользователя
}
```

<b>2. При добавлении сделать запрос в АПИ, описанного сваггером</b>

```yaml
openapi: 3.0.3
info:
  title: People info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: passportSerie
          in: query
          required: true
          schema:
            type: integer
        - name: passportNumber
          in: query
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/People'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    People:
      required:
        - surname
        - name
        - address
      type: object
      properties:
        surname:
          type: string
          example: Иванов
        name:
          type: string
          example: Иван
        patronymic:
          type: string
          example: Иванович
        address:
          type: string
          example: г. Москва, ул. Ленина, д. 5, кв. 1
```

<b>3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)</b>

<b>4. Покрыть код debug- и info-логами</b>

<b>5. Вынести конфигурационные данные в .env-файл</b>

<b>6. Сгенерировать сваггер на реализованное АПИ</b>


## Контакты

- Telegram: [@dvskchan](https://t.me/dvskchan)

- Email: duskchancellor@gmail.com
