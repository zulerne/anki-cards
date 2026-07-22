---
id: kafka_chem_kafka_otlichaetsya_ot
deck: Go
tags:
  - kafka
---

# Front

Чем Kafka отличается от RabbitMQ?

# Back

Kafka: распределённый лог, сообщения сохраняются на диске (retention), консьюмер читает по offset, можно перечитать. Высокий throughput.

RabbitMQ: классическая очередь (AMQP), сообщение удаляется после подтверждения, гибкий routing (exchange, binding). Лучше для задач с подтверждением, Kafka — для стриминга и event-driven.
