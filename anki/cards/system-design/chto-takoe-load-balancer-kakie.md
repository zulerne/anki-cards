---
id: system-design_chto_takoe_load_balancer
deck: Go
tags:
  - system-design
---

# Front

Что такое load balancer? Какие алгоритмы балансировки?

# Back

Распределяет трафик между серверами. Алгоритмы:

- **Round Robin** — по очереди.
- **Weighted Round Robin** — с учётом мощности сервера.
- **Least Connections** — на сервер с наименьшим числом соединений.
- **IP Hash** — привязка клиента к серверу.
- **Random** — случайный выбор.

L4 (TCP-уровень) vs L7 (HTTP-уровень).
