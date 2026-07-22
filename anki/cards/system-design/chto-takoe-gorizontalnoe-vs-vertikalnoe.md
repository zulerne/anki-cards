---
id: system-design_chto_takoe_gorizontalnoe_vs
deck: Go
tags:
  - system-design
---

# Front

Что такое горизонтальное vs вертикальное масштабирование?

# Back

**Вертикальное (scale up)** — мощнее сервер (CPU, RAM). Просто, но есть предел и single point of failure.

**Горизонтальное (scale out)** — больше серверов за load balancer. Сложнее (stateless, shared storage), но без предела и с отказоустойчивостью.

Для веб-сервисов обычно горизонтальное.
