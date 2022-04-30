FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o /shakemon cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /shakemon /shakemon

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/shakemon"]