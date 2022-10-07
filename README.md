# Go_Sample

## GoLang Проект-заготовка для быстрой протипизации микросервисов

### Запустить проект

```
make up-project
```

### Структура

1. [cmd](cmd/app/main.go): Основные приложения проекта
    - [main.go](cmd/app/main.go): Точка входа приложения
2. [internal](internal): Внутренний код приложения и библиотек
    - [config](internal/config): Конфигурация приложения
    - [delivery](internal/delivery): Слой определения протоколов общения
    - [entity](internal/entity): Сущности-структуры
    - [repository](internal/repository): Репозитории работы с данными
    - [service](internal/service): Сервисы
3. [pkg](pkg): Код библиотек, пригодных для использования в сторонних приложениях

### Тесты

Тесты следует писать рядом с файлом, который вы тестируете
[greeting_test.go](internal/service/greeting_test.go)
