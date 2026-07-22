---
id: docker_chto_takoe_volume_v
deck: Go
tags:
  - docker
---

# Front

Что такое volume в Docker? Зачем?

# Back

Механизм хранения данных вне writable-слоя контейнера. Данные в контейнере теряются при удалении.

Volume сохраняет данные между перезапусками. Типы: named volume (`docker volume create`), bind mount (директория хоста).

Используется для БД, логов, конфигов.
