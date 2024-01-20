# syntax=docker/dockerfile:1

FROM alpine:latest AS deps
RUN apk add --no-cache mariadb-dev sqlite-dev build-base go

FROM deps AS builder
WORKDIR /app
COPY . .
RUN cp configs/common.dev configs/common.go
RUN go mod download && go mod verify
RUN CGO_ENABLED=1 go build -v -o bootstrap cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app
RUN apk add --no-cache sqlite-libs mariadb-connector-c libgcc
COPY --from=builder /app/bootstrap .

CMD ["sh", "-c", "./bootstrap"]