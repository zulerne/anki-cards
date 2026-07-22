---
id: architecture_chto_takoe_circuit_breaker
deck: Go
tags:
  - architecture
---

# Front

Что такое Circuit Breaker pattern?

# Back

Предохранитель для вызовов внешних сервисов. Три состояния:Closed — запросы проходят, считаются ошибки.Open — после N ошибок запросы отклоняются сразу (fail fast).Half-Open — пробный запрос. Успех — closed, ошибка — open.Предотвращает каскадный отказ. Библиотека: gobreaker.
