FROM golang:1.20-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/coffee-orders .

FROM alpine:3.18
COPY --from=build /bin/coffee-orders /bin/coffee-orders
EXPOSE 8082
ENTRYPOINT ["/bin/coffee-orders"]
