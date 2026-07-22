---
id: networking_chto_takoe_websocket_chem
deck: Go
tags:
  - networking
---

# Front

Что такое WebSocket? Чем отличается от HTTP?

# Back

WebSocket — протокол полнодуплексной связи поверх TCP.

HTTP — запрос-ответ (клиент инициирует). WebSocket: после upgrade-хендшейка (HTTP 101) — постоянное соединение, сервер и клиент могут отправлять сообщения в любой момент.

Для чатов, real-time данных, игр.
