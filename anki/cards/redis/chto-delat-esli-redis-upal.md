---
id: redis_chto_delat_esli_redis
deck: Go
tags:
  - redis
---

# Front

Что делать если Redis упал? Как обеспечить отказоустойчивость?

# Back

Redis Sentinel — мониторинг, автоматический failover master->replica.Redis Cluster — шардирование + репликация, автофейловер.В приложении: fallback на БД при недоступности Redis (graceful degradation). Persistence: RDB (снимки) + AOF (журнал операций) для восстановления данных.
