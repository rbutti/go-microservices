FROM golang:1.13-alpine as build-env
WORKDIR /library-service

RUN apk update && apk add --no-cache gcc musl-dev git

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./main/app


# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /library-service/bin/app /library-service/

EXPOSE 9090
CMD ["/library-service/app"]