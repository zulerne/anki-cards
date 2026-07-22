---
id: system-design_sproektiruy_sistemu_chatov
deck: Go
tags:
  - system-design
---

# Front

Спроектируй систему чатов.

# Back

Компоненты: API Gateway, Chat Service, WebSocket-сервер, Message Store (БД), Presence Service.Доставка: WebSocket для real-time. Если получатель офлайн — сохранить в БД, доставить при подключении (push notification).Группы: fan-out на WebSocket-серверы через Redis pub/sub или Kafka.
