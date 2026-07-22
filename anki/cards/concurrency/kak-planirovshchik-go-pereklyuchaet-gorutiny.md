---
id: concurrency_kak_planirovshchik_go_pereklyuchaet
deck: Go
tags:
  - concurrency
---

# Front

Как планировщик Go переключает горутины?

# Back

Кооперативно: проверка при вызове функции (function prologue).

Вытесняюще (с Go 1.14+): async preemption через сигнал каждые ~10мс.
