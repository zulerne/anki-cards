---
id: architecture_chto_takoe_health_check
deck: Go
tags:
  - architecture
---

# Front

Что такое Health Check endpoint? Зачем нужен?

# Back

Эндпоинт (обычно GET /health или /readyz), возвращающий статус сервиса. Используется: Kubernetes liveness/readiness probes, load balancer для исключения нездоровых инстансов, мониторинг. Liveness — жив ли процесс. Readiness — готов ли принимать трафик.
