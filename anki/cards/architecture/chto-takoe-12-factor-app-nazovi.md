---
id: architecture_chto_takoe_12factor_app
deck: Go
tags:
  - architecture
---

# Front

Что такое 12-Factor App? Назови 3-4 фактора.

# Back

Методология для SaaS-приложений (Heroku, 2011). Основные факторы:

1) Codebase — один репозиторий, много деплоев.

2) Config — конфигурация через переменные окружения.

3) Dependencies — явно объявлены (`go.mod`).

4) Backing services — БД, кэш как подключаемые ресурсы.

5) Logs — поток событий в stdout.
