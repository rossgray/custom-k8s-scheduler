FROM golang:1.21-alpine3.18 as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /bin/custom-scheduler

FROM alpine:3.18

COPY --from=builder /bin/custom-scheduler /bin/custom-scheduler

CMD ["/bin/custom-scheduler"]