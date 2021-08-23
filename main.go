package main

import (
	"encoding/json"
	"fmt"

	"github.com/caiowWillian/partner-service/models"
	"github.com/caiowWillian/partner-service/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	var partners []models.Partner

	op1 := mongo.MongoOperation{
		Database:   "test",
		Collection: "partner",
	}

	//search by id
	query := bson.D{{"id", "2"}}
	mongo.Repo().Get(op1, query, &partners)

	var p models.Partner

	strJson := `{
		"id": 1002,
		"tradingName": "Adega Osasco",
		"ownerName": "Ze da Ambev",
		"document": "02.453.716/000170",
		"coverageArea": {
		   "type": "MultiPolygon",
		   "coordinates": [
			  [
				 [
					[
					   -43.36556,
					   -22.99669
					],
					[
					   -43.36539,
					   -23.01928
					],
					[
					   -43.26583,
					   -23.01802
					],
					[
					   -43.25724,
					   -23.00649
					],
					[
					   -43.23355,
					   -23.00127
					],
					[
					   -43.2381,
					   -22.99716
					],
					[
					   -43.23866,
					   -22.99649
					],
					[
					   -43.24063,
					   -22.99756
					],
					[
					   -43.24634,
					   -22.99736
					],
					[
					   -43.24677,
					   -22.99606
					],
					[
					   -43.24067,
					   -22.99381
					],
					[
					   -43.24886,
					   -22.99121
					],
					[
					   -43.25617,
					   -22.99456
					],
					[
					   -43.25625,
					   -22.99203
					],
					[
					   -43.25346,
					   -22.99065
					],
					[
					   -43.29599,
					   -22.98283
					],
					[
					   -43.3262,
					   -22.96481
					],
					[
					   -43.33427,
					   -22.96402
					],
					[
					   -43.33616,
					   -22.96829
					],
					[
					   -43.342,
					   -22.98157
					],
					[
					   -43.34817,
					   -22.97967
					],
					[
					   -43.35142,
					   -22.98062
					],
					[
					   -43.3573,
					   -22.98084
					],
					[
					   -43.36522,
					   -22.98032
					],
					[
					   -43.36696,
					   -22.98422
					],
					[
					   -43.36717,
					   -22.98855
					],
					[
					   -43.36636,
					   -22.99351
					],
					[
					   -43.36556,
					   -22.99669
					]
				 ]
			  ]
		   ]
		},
		"address": {
		   "type": "Point",
		   "coordinates": [
			  -43.297337,
			  -23.013538
		   ]
		}
	 }`
	json.Unmarshal([]byte(strJson), &p)

	//mongo.Repo().Insert(op, p)

	queryM := bson.M{
		"$and": []bson.M{
			{
				"coverageArea": bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": []float64{-43.3636, -22.99351},
						},
					},
				},
			},
			{
				"address": bson.M{
					"$near": bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": []float64{-43.3636, -22.99351},
						},
					},
				},
			},
		},
	}

	//mongo.Repo().GetAll(op1, &partners)
	mongo.Repo().Get(op1, queryM, &partners)

	fmt.Println("=========================================")
	for _, element := range partners {

		fmt.Println(element)
		fmt.Println("=========================================")

	}
}
