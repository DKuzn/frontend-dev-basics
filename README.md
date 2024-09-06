# Основы frontend-разработки

Данный репозиторий содержит примеры для выполнения практических работ по дисциплине "Основы frontend-разработки".

## PR1

В папке PR1 содержатся материалы для выполнения практической работы №1. 

## PR2-3

В папке PR2-3 содержатся материалы для выполнения практических работ №2 и №3.

## Документация к API из практической работы №3

Для запуска обслуживания API необходимо запустить файл [api.exe](https://github.com/DKuzn/frontend-dev-basics/releases) внутри директории PR2-3 этого репозитория.

```powershell
.\api.exe
```

Сайт и API будут доступны по адресу [localhost:8080](http://localhost:8080).

### POST /api/event

Маршрут добавляет новую запись о мероприятии в базу данных. Принимает данные в формате JSON. Ничего не возвращает в ответ.


**Пример запроса**:

```json
{
    "event_name": "Название мероприятия в произвольном текством формате",
    "event_plate": "Название площадки в произвольном текством формате",
    "event_type": "Тип мероприятия в произвольном текством формате",
    "event_date": "01.01.1970",
    "event_time": "23:59"
}
```

### GET /api/event

Маршрут возвращает список ID всех мероприятий в базе данных в формате JSON.

**Пример ответа**:

```json
[1,2,3]
```

### GET /api/event/:id

Маршрут возвращает мероприятие в базе данных по ID в формате JSON.

**Пример ответа**:

```json
{
    "id": 1,
    "event_name": "Название мероприятия",
    "event_plate": "Название площадки",
    "event_type": "Тип мероприятия",
    "event_date": "01.01.1970",
    "event_time": "23:59"
}
```

### PUT /api/event/:id

Маршрут полностью обновляет существующую запись в базе данных по ID. Принимает данные в формате JSON. Ничего не возвращает.

**Пример запроса**:

```json
{
    "event_name": "Название мероприятия в произвольном текством формате",
    "event_plate": "Название площадки в произвольном текством формате",
    "event_type": "Тип мероприятия в произвольном текством формате",
    "event_date": "01.01.1970",
    "event_time": "23:59"
}
```

### DELETE /api/event/:id

Маршрут удаляет существующую запись в базе данных по ID. Ничего не возвщает.
