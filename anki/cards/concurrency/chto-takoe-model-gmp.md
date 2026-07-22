---
id: concurrency_chto_takoe_model_gmp
deck: Go
tags:
  - concurrency
---

# Front

Что такое модель GMP?

# Back

G = Goroutine (горутина). M = Machine (поток ОС). P = Processor (логический процессор, от 1 до GOMAXPROCS). P связывает G с M: горутина выполняется на M только через P. Каждый P имеет локальную очередь горутин.
