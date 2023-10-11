FROM golang:1.21-alpine3.18 AS build

WORKDIR /app

COPY . .

# install dependencies
RUN go mod tidy

# build binary
RUN go build -o ./bin/binary ./cmd/http/main.go



FROM alpine:3.18 AS final

WORKDIR /app

# copy binary
COPY --from=build /app/bin/binary ./

EXPOSE 80

ENTRYPOINT [ "/app/binary" ]