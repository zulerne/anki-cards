---
id: docker_chto_takoe_volume_v
deck: Go
tags:
  - docker
---

# Front

Что такое volume в Docker? Зачем?

# Back

Механизм хранения данных вне writable-слоя контейнера. Writable-слой обычно
сохраняется при restart, но исчезает при удалении контейнера.

Named volume сохраняет данные между удалением и созданием контейнера. Виды
подключения: named volume (`docker volume create`) и bind mount (директория
хоста).

Используется для БД, логов, конфигов.
