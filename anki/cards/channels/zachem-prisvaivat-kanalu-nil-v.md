---
id: channels_zachem_prisvaivat_kanalu_nil
deck: Go
tags:
  - channels
---

# Front

Зачем присваивать каналу nil в select?

# Back

Чтобы отключить ветку `select` — `nil`-канал навсегда блокируется, и case перестаёт срабатывать.
