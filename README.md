# partner-service


### Stack
- [golang](https://golang.org/)
- [mongodb](https://www.mongodb.com/pt-br)
- [docker](https://www.docker.com/)

### how to run the solution
see below how to run the solution locally and by docker image

#### Local
1. execute ``go mod download``in root path

2. boot mongo db. Access the path ``./dev/`` and run ``docker-compose up`` 

3. add the environment variable to your local system ``mongo_uri`` ex: ``mongodb://localhost:27017/teste``

4. access ``./cmd/``and run ``go run main.go``
