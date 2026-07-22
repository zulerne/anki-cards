---
id: architecture_chto_takoe_idempotentnost_pochemu
deck: Go
tags:
  - architecture
---

# Front

Что такое идемпотентность? Почему важна для API?

# Back

Повторный запрос даёт тот же результат. GET, PUT, DELETE — идемпотентны. POST — нет. Важно: сетевые retry, дубли из-за таймаутов. Реализация: idempotency key в заголовке, проверка "уже выполнено" перед действием.
