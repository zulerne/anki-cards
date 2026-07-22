---
id: architecture_v_chem_raznitsa_mezhdu
deck: Go
tags:
  - architecture
---

# Front

В чём разница между sync и async коммуникацией между сервисами?

# Back

Sync (HTTP/gRPC): вызывающий ждёт ответ. Простая логика, но coupling, cascading failures при недоступности.

Async (Kafka, RabbitMQ): сообщение в очередь, ответ не ждём. Слабая связанность, отказоустойчивость, но сложнее отладка, eventual consistency.
