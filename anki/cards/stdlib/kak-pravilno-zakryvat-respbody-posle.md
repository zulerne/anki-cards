---
id: stdlib_kak_pravilno_zakryvat_respbody
deck: Go
tags:
  - stdlib
---

# Front

Как правильно закрывать resp.Body после http-запроса? Зачем?

# Back

resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()Если не закрыть — TCP-соединение не вернётся в пул, утечка файловых дескрипторов и горутин. Закрывать нужно даже если тело не читается.
