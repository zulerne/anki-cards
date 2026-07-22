---
id: architecture_chto_takoe_api_gateway
deck: Go
tags:
  - architecture
---

# Front

Что такое API Gateway? Зачем нужен?

# Back

Единая точка входа для всех клиентов. Маршрутизирует запросы к микросервисам. Функции: аутентификация, rate limiting, агрегация ответов, TLS termination, кэширование, трансформация запросов. Примеры: Kong, NGINX, AWS API Gateway.
