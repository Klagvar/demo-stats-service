# demo-stats-service

Маленький HTTP-сервис для демонстрации работы [testgen-agent](https://github.com/Klagvar/testgen-agent).

## Что внутри

- `core/` — функции расчёта статистик: `Mean`, `Min`, `Max`, `Round`.
- `storage/` — интерфейс `Store` + in-memory реализация.
- `handler/` — HTTP-handler `POST /stats`, который принимает серию значений, сохраняет в store, возвращает агрегаты.
- `main.go` — bootstrap сервера на `:8080`.

## Запуск

```bash
go run .
```

## Тесты

```bash
go test ./...
```

## CI

В `.github/workflows/testgen.yml` настроен запуск `testgen-agent` на каждый pull request, затрагивающий `*.go` файлы. Агент анализирует diff, генерирует тесты для изменённых функций, валидирует их через `go test`, считает метрики качества и оставляет сводный комментарий в PR.
