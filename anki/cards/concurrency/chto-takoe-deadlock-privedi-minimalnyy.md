---
id: concurrency_chto_takoe_deadlock_privedi
deck: Go
tags:
  - concurrency
---

# Front

Что такое deadlock? Приведи минимальный пример на Go.

# Back

Состояние, когда две или более горутины ждут друг друга и ни одна не может продолжить.Минимальный пример:ch := make(chan int)
ch <- 1 // main заблокирован навсегдаfatal error: all goroutines are asleep - deadlock!
