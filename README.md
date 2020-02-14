# go-microservices-goldcopy

A Reference project to implement microservices using go programming language. This project acts as a template/goldcopy to setup and implement microservices using Golang within short period of time.

This application has the following features:

* A library microservices that exposes APIs to store, retrieve, update and delete list of books in database
* Provide docker and docker-compose tempates to quickly build docker images, run containers and push the containers to docker registry
* Custom logging to print all request and response from handlers
* Demonstrates the use of interceptors to inject headers into the response
* Provides a health monitoring API to check the service and DB health status

### Tools and Technologies

This reference project uses the following tools and technologies:

| Technology | Version | Description |
| ------ | ------ |------ |
| Golang | 1.13 | Programming language|
| go-chi | 2.1.0 | Router Framework |
| mysql | 8.0 | Database |
| gorm | 1.9.12 | ORM framework |
| zerolog | 1.17.2 | Logging framework |
| Docker | 18.09.2 | Container Creator |
| git | 2.20.1 | Soure Code Management |
| Goclipse | 0.16.1| Integrated Developement Environment |

### Deliverables

The deliverables for the application can be found at the below location
[Library-Service/deliverables](https://github.com/rbutti/go-microservices-goldcopy/tree/master/library-service/deliverables)

| Deliverable | Summary |
| ------ | ------ |
| Swagger.yaml | Swagger documentation for library-service API |
| ProjectStructure.PNG |Diagram depicting the project structure |

### Assumptions

* Docker is installed in your system following the [guide](https://docs.docker.com/install/)
* Golang is installed in your system following the [guide](https://golang.org/doc/install#install) and the $GOROOT and $GOPATH are set correctly
* The code is expected to use Goclipse project structure for GoProjects


### Execution
An docker image of the code is currently pushed into public docker hub repository. Follow the below steps once docker is installed in your local machine
* Open terminal and pull the docker image
```sh
$ docker pull ravikiran763/library-service
```

* run the following command to start the library services
```sh
$ docker run -d -p 9090:9090 ravikiran763/library-service /bin/bash
```

* go to your browser and check for http://localhost:9090  you should see "Hello World!!" printed on the browser

##### Example
Go to browser and hit the following URL : http://localhost:9090/api/v1/library/books


Expected output

```json
[
   {
      "id":1,
      "title":"I Too Had a Love Story",
      "author":"Ravinder Singh",
      "published_date":"2008-01-01",
      "image_url":"https://images-na.ssl-images-amazon.com/images/I/81phwRtlzCL.jpg",
      "description":"I Too Had a Love Story is an English autobiographical novel written by Ravinder Singh"
   }
]

```


### Technical Design

##### Project Structure

![Project Structure](https://github.com/rbutti/go-microservices-goldcopy/blob/master/library-service/deliverables/ProjectStructure.png "Project Structure")

##### Package Structure

| Package | Summary |
| ------ | ------ |
| library-service | code containing library microservices |
| library-service/src | golang code of microser  |
| library-service/src/config | configuration files  |
| library-service/src/main | entrypoint to the application |
| library-service/src/model| models for library service|
| library-service/src/repository | classes to interact with database |
| library-service/src/server| contains routers and handlers of the application  |
| library-service/src/setup |setup files used during build |
| library-service/src/server| utility and constant files  |

##### Model

| Class | Summary |
| ------ | ------ |
| book.go| Domain object representing a book |
| bookDto.go| Data transfer object for book domain object |
| bookForm.go| presentation  object for book domain object |

##### Class Design

| Class | Summary |
| ------ | ------ |
| libraryApp.go | Entrypoint to the Application consisting main() method|
| appConfig.go | Contains configuration related to applction like server details etc |
| ormConfig.go | Contains configuration related to database and  gorm |
| bookRepo.go | Repository class to perform CRUD operation on Book table |
| handler.go |Application request handler |
| handlerWrapper.go | Wrapper for request handler and logging |
| rootHander.go | handles request to '/' |
| libraryHander.go | handles request to '/api/vi/library-service' |
| libraryHander.go | handles request to '/health' |
| router.go | routes the request to necessary handler |
| responseInterceptor.go | intercepts router reponse |
| constants.go | contains application constants |
| logger.go | wrapper to zerolog framework |
| logEntry.go | log entry formatter |

### Future Enhancements

* Implement swagger generator for the router and handlers
* Implement service-discovery, load-balance and gateway
* Deploy in kubernates and implement liveness and readiness probe to monitor health
* Deploy in cloud
* Create another microservice and demonstrate communication between services
* implement OAuth security


### Contact Information

**Author** : Ravikiran Butti,
**Email Id** : (ravikiran763@gmail.com)

