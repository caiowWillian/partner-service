# partner-service


### Stack
- [golang](https://golang.org/)
- [mongodb](https://www.mongodb.com/pt-br)
- [docker](https://www.docker.com/)

### how to run the solution
see below how to run the solution locally and by docker image

### Local
- execute ``go mod download``in root path
- boot mongo db. Access the path ``./dev/`` and run ``docker-compose up`` 
- add the environment variable to your local system ``mongo_uri`` ex: ``mongodb://localhost:27017/teste``
- access ``./cmd/``and run ``go run main.go``

### Docker

In root RUN ``docker-compose up``

### Endpoints

Get near partner
```
curl --location --request GET 'http://localhost:5555/partner?lat=-43.24067&long=-22.99381' \
--data-raw ''
```

Create a new partner
```
curl --location --request POST 'http://localhost:5555/partner?lat=-43.24067&long=-22.99381' \
--header 'Content-Type: application/json' \
--data-raw '{
  "tradingName": "Adega da Cerveja - Pinheiros",
  "ownerName": "Zé da Silva",
  "document": "1432132123891/30001",
  "coverageArea": { 
    "type": "MultiPolygon", 
    "coordinates": [
      [[[30, 20], [45, 40], [10, 40], [30, 20]]], 
      [[[15, 5], [40, 10], [10, 20], [5, 10], [15, 5]]]
    ]
  },
  "address": { 
    "type": "Point",
    "coordinates": [-46.57421, -21.785741]
  }
}'
```

Get partner by id
```
curl --location --request GET 'http://localhost:5555/partner/61291c48fcd6803e412ceb9e' \
--data-raw ''
```


