---
id: system-design_sproektiruy_url_shortener
deck: Go
tags:
  - system-design
---

# Front

Спроектируй URL shortener.

# Back

Компоненты: API-сервис, БД (хранение long->short URL), кэш (Redis).Генерация: Base62 от auto-increment ID или хеш (MD5/SHA256, первые 7 символов).Запись: POST /shorten -> генерация ID -> сохранение в БД.Чтение: GET /:short -> проверка кэша -> БД -> 301 редирект.Масштабирование: несколько инстансов, Redis, CDN.
