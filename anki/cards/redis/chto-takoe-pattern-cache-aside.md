---
id: redis_chto_takoe_pattern_cache
deck: Go
tags:
  - redis
---

# Front

Что такое паттерн cache aside (lazy loading)?

# Back

1) Читаем из кэша. Если есть (cache hit) — возвращаем.2) Если нет (cache miss) — читаем из БД, записываем в кэш, возвращаем.3) При записи в БД — инвалидируем кэш (delete, не update).Просто, кэшируются только запрашиваемые данные. Минус: первый запрос всегда медленный (cold start).
