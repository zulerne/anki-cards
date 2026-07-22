---
id: concurrency_chto_takoe_synconce_potokobezopasna
deck: Go
tags:
  - concurrency
---

# Front

Что такое sync.Once? Потокобезопасна ли она?

# Back

Примитив для однократного выполнения функции, вне зависимости от числа вызовов once.Do(fn).Потокобезопасна — если несколько горутин вызывают Do одновременно, только одна выполнит fn, остальные заблокируются до завершения.В Go 1.21+: sync.OnceFunc(fn) / sync.OnceValue(fn) — удобнее.
