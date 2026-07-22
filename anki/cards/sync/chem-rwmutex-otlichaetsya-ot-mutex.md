---
id: sync_chem_rwmutex_otlichaetsya_ot
deck: Go
tags:
  - sync
---

# Front

Чем RWMutex отличается от Mutex?

# Back

RWMutex допускает несколько одновременных читателей (RLock), но только одного писателя (Lock). Используй когда чтений значительно больше записей. Mutex блокирует всех — и читателей, и писателей.
