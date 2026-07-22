---
id: system-design_chto_takoe_consistent_hashing
deck: Go
tags:
  - system-design
---

# Front

Что такое consistent hashing?

# Back

Алгоритм распределения ключей по узлам. Узлы и ключи размещаются на кольце (hash ring).

Ключ попадает на ближайший узел по часовой стрелке. При добавлении/удалении узла перемещается минимум ключей (~1/N).

Виртуальные узлы улучшают равномерность.

Используется в кэшах, Cassandra, DynamoDB.
