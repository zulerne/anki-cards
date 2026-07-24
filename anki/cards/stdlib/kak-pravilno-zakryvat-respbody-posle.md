---
id: stdlib_kak_pravilno_zakryvat_respbody
deck: Go
tags:
  - stdlib
---

# Front

Как правильно закрывать `resp.Body` после http-запроса? Зачем?

# Back

```go
resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()
```

Закрывать body нужно всегда. Чтобы HTTP transport мог переиспользовать соединение, обычно также следует прочитать body до EOF; простого `Close` может быть недостаточно для reuse. Это не следует называть гарантированной утечкой горутины.
