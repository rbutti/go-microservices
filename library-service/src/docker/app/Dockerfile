#Docker Multi-Stage Build
#BUILD  - using golang alpine
FROM golang:1.13-alpine as build-env

#Setup working directory in the image
WORKDIR /library-service

#Run apk update
RUN apk update && apk add --no-cache gcc musl-dev git

#Copy from current location to image working director
COPY . .

#Run the go build command
RUN go build -ldflags '-w -s' -a -o ./bin/app ./main/app


# DEPLOYMENT
FROM alpine

#Run apk update
RUN apk update && apk add --no-cache bash

#Copy from build enviroment to deployment image
COPY --from=build-env /library-service/bin/app /library-service/

# CONFIGURATION
EXPOSE 9090
CMD ["/library-service/app"]