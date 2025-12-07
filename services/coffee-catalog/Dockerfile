FROM golang:1.20-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/coffee-catalog .

FROM alpine:3.18
COPY --from=build /bin/coffee-catalog /bin/coffee-catalog
EXPOSE 8081
ENTRYPOINT ["/bin/coffee-catalog"]
