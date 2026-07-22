---
id: concurrency-patterns_kak_realizovat_rate_limiter
deck: Go
tags:
  - concurrency-patterns
---

# Front

Как реализовать rate limiter с помощью time.Ticker и канала?

# Back

```go
limiter := time.NewTicker(100 * time.Millisecond)
defer limiter.Stop()
for req := range requests {
    <-limiter.C // ждём тик
    process(req)
}
```

Тикер пропускает одно событие за интервал. Для burst — буферизированный канал, заполненный заранее.
