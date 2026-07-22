---
id: networking_kakie_httpmetody_idempotentnye
deck: Go
tags:
  - networking
---

# Front

Какие HTTP-методы идемпотентные?

# Back

Идемпотентные (повторный запрос = тот же результат): GET, PUT, DELETE, HEAD, OPTIONS.

НЕ идемпотентные: POST, PATCH.

Безопасные (не меняют состояние): GET, HEAD, OPTIONS.
