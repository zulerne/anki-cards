---
id: stdlib_kak_zadat_timeout_http_client
deck: Go
tags:
  - http
  - context
---

# Front

Как ограничить время HTTP-запроса в Go?

# Back

Для общего ограничения используй `http.Client{Timeout: d}`.

Для дедлайна конкретного запроса передай context:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()
req = req.WithContext(ctx)
```

Оба варианта должны сопровождаться обработкой ошибки и закрытием `resp.Body`.
