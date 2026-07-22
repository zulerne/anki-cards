---
id: system-design_chto_takoe_capteorema
deck: Go
tags:
  - system-design
---

# Front

Что такое CAP-теорема?

# Back

В распределённой системе можно гарантировать только 2 из 3:

- **Consistency** — все узлы видят одни данные одновременно.
- **Availability** — каждый запрос получает ответ.
- **Partition tolerance** — система работает при разрыве сети.

Partition tolerance обязателен -> выбор между CP (консистентность) и AP (доступность).
