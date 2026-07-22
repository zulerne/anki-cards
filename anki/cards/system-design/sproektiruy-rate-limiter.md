---
id: system-design_sproektiruy_rate_limiter
deck: Go
tags:
  - system-design
---

# Front

Спроектируй rate limiter.

# Back

Алгоритмы:

- **Token Bucket** — корзина с токенами, пополняется с постоянной скоростью. Запрос берёт токен. Позволяет burst.
- **Sliding Window** — считаем запросы за скользящее окно времени.

Хранение счётчиков в Redis (`INCR` + `EXPIRE`). Ключ: `user_id` + окно. Ответ 429 при превышении.
