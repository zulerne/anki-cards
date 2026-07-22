---
id: docker_chto_takoe_multistage_build
deck: Go
tags:
  - docker
---

# Front

Что такое multi-stage build? Зачем?

# Back

Несколько FROM в Dockerfile. Первый этап — сборка (компилятор, зависимости), второй — только бинарник. Итоговый образ маленький (без компилятора, исходников). Для Go:FROM golang AS build
RUN go build -o app .
FROM alpine
COPY --from=build /app/app /app
