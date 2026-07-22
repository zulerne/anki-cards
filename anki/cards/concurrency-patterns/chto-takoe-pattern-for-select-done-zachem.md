---
id: concurrency-patterns_chto_takoe_pattern_forselectdone
deck: Go
tags:
  - concurrency-patterns
---

# Front

Что такое паттерн "for-select-done"? Зачем done-канал?

# Back

```go
for {
    select {
    case <-done:
        return
    case v := <-ch:
        process(v)
    }
}
```

done-канал (или `ctx.Done()`) позволяет корректно завершить горутину извне. Без него горутина блокируется навечно — утечка.
