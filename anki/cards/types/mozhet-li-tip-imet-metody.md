---
id: types_mozhet_li_tip_imet
deck: Go
tags:
  - types
---

# Front

Может ли тип иметь методы и с pointer receiver, и с value receiver?

# Back

Да. На одном типе можно определить часть методов с `*T` и часть с `T`.

Но конкретный метод может иметь только один тип receiver.
