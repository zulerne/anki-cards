---
id: concurrency_chto_takoe_data_race
deck: Go
tags:
  - concurrency
---

# Front

Что такое data race?

# Back

Две или более горутины одновременно обращаются к одной области памяти, и хотя бы одна из них пишет — без синхронизации.

Пример:

```go
counter := 0
for range 100 {
    go func() { counter++ }()
}
```

Обнаружение: `go test -race ./...` или `go run -race .`

Защита: `sync.Mutex`, `sync/atomic`, каналы.
