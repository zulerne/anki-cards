---
id: docker_chto_takoe_kubernetes_pod
deck: Go
tags:
  - docker
---

# Front

Что такое Kubernetes pod, service, deployment?

# Back

Pod — минимальная единица деплоя, один или несколько контейнеров с общей сетью и storage.

Service — стабильный сетевой адрес для группы подов (load balancing, DNS-имя).

Deployment — управляет ReplicaSet, обеспечивает rolling update, откат, масштабирование подов.
