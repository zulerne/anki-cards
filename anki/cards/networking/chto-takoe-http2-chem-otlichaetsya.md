---
id: networking_chto_takoe_http2_chem
deck: Go
tags:
  - networking
---

# Front

Что такое HTTP/2? Чем отличается от HTTP/1.1?

# Back

HTTP/2: бинарный протокол (вместо текстового), мультиплексирование (несколько запросов через одно соединение без head-of-line blocking), сжатие заголовков (HPACK), server push.

HTTP/1.1: один запрос за раз на соединение, или keep-alive с очередью.
