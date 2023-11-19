FROM golang:1.21 as builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/bin/server .

FROM gcr.io/distroless/static-debian11

COPY --from=builder /usr/bin/ /app/bin/
COPY --from=builder /usr/src/app/did-configuration.json /opt/well-known/did-configuration.json

ENTRYPOINT [ "/app/bin/server" ]