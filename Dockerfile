FROM golang:1.23.0 as builder

ARG GOARCH

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY application/parser.go main.go
COPY common common
COPY publisher publisher
COPY domain domain
COPY usecase usecase
COPY application application
COPY in-memory-storage in-memory-storage
COPY json-rpc json-rpc
COPY rest rest

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -a -o app main.go

FROM gcr.io/distroless/static:nonroot
EXPOSE 8080
WORKDIR /
COPY --from=builder /workspace/app .
USER 65532:65532

ENTRYPOINT ["/app"]
