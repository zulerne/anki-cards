# Reference practice

Эталонные, минимальные и идиоматичные примеры для тем из `anki/cards`.

Это не дневник решений и не production framework. Дневниковые решения находятся
в `practice/daily/<date>`. Каждый каталог показывает
один законченный паттерн, который можно запустить или протестировать:

```bash
go run ./practice/reference/concurrency/worker-pool
go test ./practice/reference/testing/table-driven
go test ./...
```

## Карта примеров

| Пример | Тема карточек |
| --- | --- |
| `concurrency/worker-pool` | worker pool, context cancellation, `WaitGroup.Go` |
| `concurrency/errgroup` | `errgroup.WithContext`, отмена при первой ошибке |
| `concurrency/fan-in` | fan-out/fan-in и закрытие выходного канала |
| `context/pipeline` | pipeline с `ctx.Done()` на каждой стадии |
| `http/server` | `ServeMux` Go 1.22+, request context, graceful shutdown |
| `io/copy` | `io.Reader`, `io.Writer`, `io.Copy` |
| `generics/collections` | type parameters, `slices`, `maps` |
| `errors/join` | `errors.Join`, `errors.Is` |
| `sync/waitgroup-go` | `sync.WaitGroup.Go` в Go 1.25+ |
| `testing/table-driven` | table-driven tests без `time.Sleep` |
| `design/builder` | builder с валидацией на этапе `Build` |
| `design/functional-options` | functional options для конструктора |
| `http/middleware` | middleware chain на `http.Handler` |
| `concurrency/semaphore` | bounded concurrency через buffered channel |
| `types/type-switch` | type switch и порядок `case` |
| `sync/once` | thread-safe lazy initialization через `sync.OnceValue` |

Для каждого примера желательно иметь соответствующую карточку. Если практика
требует написать новый код по теме, сначала добавь или уточни карточку, а затем
добавь эталонный вариант сюда.
