---
id: stdlib_kak_realizovat_middleware_v
deck: Go
tags:
  - stdlib
---

# Front

Как реализовать middleware в net/http?

# Back

Функция, принимающая http.Handler и возвращающая http.Handler:func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL)
        next.ServeHTTP(w, r)
    })
}Оборачивается цепочкой: logging(auth(handler)).
