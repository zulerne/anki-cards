---
id: architecture_chem_mikroservisy_otlichayutsya_ot
deck: Go
tags:
  - architecture
---

# Front

Чем микросервисы отличаются от монолита? Минусы микросервисов.

# Back

Минусы: сетевая сложность (latency, отказы), распределённые транзакции, сложная отладка и трассировка, operational overhead (k8s, service mesh, мониторинг), eventual consistency вместо ACID.

Монолит проще на старте — начинай с монолита.
