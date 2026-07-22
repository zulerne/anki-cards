---
id: concurrency_kak_realizovat_semafor
deck: Go
tags:
  - concurrency
  - channels
---

# Front

Как реализовать семафор для ограничения числа одновременно выполняющихся задач?

# Back

Используй буферизированный канал ёмкостью `N`:

```go
sem := make(chan struct{}, N)
sem <- struct{}{} // acquire
defer func() { <-sem }() // release
```

Если нужен только лимит параллелизма, важно освобождать слот через `defer`
после успешного acquire. Для ожидания набора задач можно совместить семафор с
`sync.WaitGroup` или `errgroup.Group`.
