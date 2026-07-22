---
id: kafka_chto_garantiruet_atleastonce_atmostonce
deck: Go
tags:
  - kafka
---

# Front

Что гарантирует at-least-once, at-most-once и exactly-once delivery?

# Back

At-most-once: сообщение может потеряться, но не дублируется. Fire and forget.At-least-once: сообщение доставлено минимум раз, возможны дубли. Нужна идемпотентность обработчика.Exactly-once: ровно один раз. Сложно: требует транзакций (Kafka transactions, idempotent producer).
