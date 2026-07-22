---
id: context_kak_pravilno_obrabatyvat_ctxdone
deck: Go
tags:
  - context
---

# Front

Как правильно обрабатывать ctx.Done() в горутине?

# Back

Через select: одна ветка — рабочий канал, другая — <-ctx.Done() для завершения.При получении сигнала — освободить ресурсы и выйти.Ошибку можно получить через ctx.Err(): context.Canceled или context.DeadlineExceeded.
