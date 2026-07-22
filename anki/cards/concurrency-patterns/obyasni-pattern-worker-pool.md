---
id: concurrency-patterns_obyasni_pattern_worker_pool
deck: Go
tags:
  - concurrency-patterns
---

# Front

Объясни паттерн Worker Pool.

# Back

N горутин-воркеров читают задачи из общего канала. Ограничивает параллелизм, переиспользует горутины. Используется при обработке большого количества однотипных задач (HTTP-запросы, файлы, записи БД).

```go
jobs := make(chan Job)
for range n { go worker(jobs) }
for _, j := range tasks { jobs <- j }
```
