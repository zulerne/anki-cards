---
id: networking_chto_proiskhodit_pri_vvode
deck: Go
tags:
  - networking
---

# Front

Что происходит при вводе URL в браузер?

# Back

1) DNS-резолвинг — домен в IP-адрес.2) TCP-соединение (3-way handshake).3) TLS handshake (если HTTPS).4) HTTP-запрос (GET /).5) Сервер обрабатывает, отправляет ответ.6) Браузер парсит HTML, загружает CSS/JS/изображения.7) Рендеринг страницы (DOM + CSSOM -> layout -> paint).
