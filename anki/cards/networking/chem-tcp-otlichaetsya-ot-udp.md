---
id: networking_chem_tcp_otlichaetsya_ot
deck: Go
tags:
  - networking
---

# Front

Чем TCP отличается от UDP? Когда что использовать?

# Back

TCP: надёжный, с установкой соединения (3-way handshake), гарантия порядка и доставки, flow/congestion control. Для HTTP, БД, файлов.

UDP: без соединения, без гарантий доставки/порядка, минимальные задержки. Для DNS, видео/аудио стриминга, игр, метрик.
