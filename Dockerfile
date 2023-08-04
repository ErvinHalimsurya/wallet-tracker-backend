## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# to change accoding to main.go location
RUN CGO_ENABLED=0 go build cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main ./
COPY config.yml ./




EXPOSE 8080

RUN chmod +x ./main

CMD [ "./main" ]
