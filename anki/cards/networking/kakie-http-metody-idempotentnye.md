---
id: networking_kakie_httpmetody_idempotentnye
deck: Go
tags:
  - networking
---

# Front

Какие HTTP-методы идемпотентные?

# Back

Идемпотентные по семантике HTTP (повторение имеет тот же intended effect): GET,
PUT, DELETE, HEAD, OPTIONS.

POST обычно неидемпотентен. PATCH может быть идемпотентным или нет в зависимости
от операции.

Безопасные (не меняют состояние): GET, HEAD, OPTIONS.
