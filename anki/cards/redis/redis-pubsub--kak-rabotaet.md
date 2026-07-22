---
id: redis_redis_pubsub__kak
deck: Go
tags:
  - redis
---

# Front

Redis pub/sub — как работает? Чем отличается от Kafka?

# Back

Pub/sub: publisher отправляет в канал, все подписчики получают сообщение в реальном времени. Fire-and-forget — если подписчика нет, сообщение теряется.Kafka: сообщения сохраняются на диске, consumer читает по offset, можно перечитать. Kafka — надёжная очередь, Redis pub/sub — real-time нотификации.
