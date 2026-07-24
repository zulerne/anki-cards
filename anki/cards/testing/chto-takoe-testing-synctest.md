---
id: testing_chto_takoe_testing_synctest
deck: Go
tags:
  - testing
  - concurrency
---

# Front

Для чего нужен `testing/synctest`?

# Back

`testing/synctest` помогает детерминированно тестировать конкурентный код в «bubble»: горутины и время управляются тестовым окружением.

`synctest.Test` и `synctest.Wait` добавлены в Go 1.25. `Wait` ждёт, пока горутины bubble не станут устойчиво заблокированными, поэтому тесты не должны использовать `time.Sleep` для синхронизации.
