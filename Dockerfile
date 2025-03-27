FROM docker.repo.frg.tech/golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .


FROM docker.repo.frg.tech/alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
