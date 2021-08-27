# partner-service


### Stack
- [golang](https://golang.org/)
- [mongodb](https://www.mongodb.com/pt-br)
- [docker](https://www.docker.com/)

### how to run the solution
see below how to run the solution locally and by docker image

#### Local
- execute ``go mod download``in root path
- boot mongo db. Access the path ``./dev/`` and run ``docker-compose up`` 
- add the environment variable to your local system ``mongo_uri`` ex: ``mongodb://localhost:27017/teste``
- access ``./cmd/``and run ``go run main.go``

#### Docker

In root RUN ``docker-compose up``
