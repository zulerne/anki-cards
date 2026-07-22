---
id: networking_chto_takoe_protobuf_chem
deck: Go
tags:
  - networking
---

# Front

Что такое Protobuf? Чем лучше JSON для межсервисного взаимодействия?

# Back

Protocol Buffers — бинарный формат сериализации от Google. Схема в .proto файлах, кодогенерация для любого языка. Компактнее JSON (в 3-10x), быстрее сериализация/десериализация, строгая типизация, обратная совместимость. JSON лучше для публичных API и дебага.
