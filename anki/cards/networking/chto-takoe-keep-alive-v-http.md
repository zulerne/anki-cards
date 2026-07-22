---
id: networking_chto_takoe_keepalive_v
deck: Go
tags:
  - networking
---

# Front

Что такое keep-alive в HTTP?

# Back

Переиспользование TCP-соединения для нескольких HTTP-запросов.

Без keep-alive: на каждый запрос — новое TCP + TLS соединение (дорого).

В HTTP/1.1 включено по умолчанию (`Connection: keep-alive`). В HTTP/2 — встроенное мультиплексирование.
