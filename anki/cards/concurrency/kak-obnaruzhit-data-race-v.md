---
id: concurrency_kak_obnaruzhit_data_race
deck: Go
tags:
  - concurrency
---

# Front

Как обнаружить data race в Go?

# Back

`go test -race ./...` или `go run -race main.go`. Race detector работает в runtime.
