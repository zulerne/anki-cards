---
id: system-design_sproektiruy_sistemu_notifikatsiy
deck: Go
tags:
  - system-design
---

# Front

Спроектируй систему нотификаций.

# Back

Компоненты: Notification Service (API), Router (выбор канала: email/push/SMS), Provider Adapters (SendGrid, FCM, Twilio), Queue (Kafka).

Поток: событие -> Kafka -> Router -> адаптер провайдера.

Retry с exponential backoff. Дедупликация по idempotency key. Preferences Service для настроек пользователя.
