---
id: architecture_chto_takoe_service_discovery
deck: Go
tags:
  - architecture
---

# Front

Что такое service discovery?

# Back

Механизм, позволяющий сервисам находить друг друга. Вместо хардкода адресов — сервис регистрируется при старте, другие запрашивают реестр. Client-side (приложение запрашивает реестр: Consul, etcd) или server-side (балансировщик знает адреса: Kubernetes DNS).
