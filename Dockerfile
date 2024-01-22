# build stage
FROM golang:1.21-alpine AS builder

# set working directory
WORKDIR /app

# copy source code
COPY . .

# install dependencies
RUN go mod download

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/EWallet ./cmd/http/main.go

# final stage
FROM alpine AS final

# set working directory
WORKDIR /app

# copy binary
COPY --from=builder /app/bin/EWallet ./

COPY ./.env ./

ENTRYPOINT [ "./EWallet" ]