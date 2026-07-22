---
id: concurrency_kak_s_pomoshchyu_kanala
deck: Go
tags:
  - concurrency
---

# Front

Как с помощью канала реализовать семафор на N горутин?

# Back

sem := make(chan struct{}, N)Перед работой: sem <- struct{}{} (занять слот). После: <-sem (освободить). Буфер размером N ограничивает параллелизм.
