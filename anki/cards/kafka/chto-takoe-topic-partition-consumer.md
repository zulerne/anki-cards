---
id: kafka_chto_takoe_topic_partition
deck: Go
tags:
  - kafka
---

# Front

Что такое topic, partition, consumer group, offset в Kafka?

# Back

Topic — именованный поток сообщений (лог).Partition — часть топика, обеспечивает параллелизм. Сообщения внутри партиции упорядочены.Consumer group — группа консьюмеров, каждая партиция читается одним консьюмером в группе.Offset — позиция сообщения в партиции, консьюмер коммитит offset.
