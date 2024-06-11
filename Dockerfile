FROM golang:1.22 AS build

WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY ./app .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go 

FROM alpine:3.9
RUN apk --no-cache add curl
WORKDIR /
COPY --from=build /main /main
ENTRYPOINT ["/main"]
