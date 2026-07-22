---
id: docker_chto_takoe_cicd_iz
deck: Go
tags:
  - docker
---

# Front

Что такое CI/CD? Из каких этапов состоит pipeline?

# Back

CI (Continuous Integration): автоматическая сборка и тесты при каждом коммите.CD (Continuous Delivery/Deployment): автоматический деплой после прохождения тестов.Этапы: lint -> build -> unit tests -> integration tests -> build image -> deploy staging -> deploy production.
