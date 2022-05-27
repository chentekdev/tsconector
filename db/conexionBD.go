package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de coneccion a la base de datos*/
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://chentek_admin:IHFM4UDE3wqiFx2H@cluster.xtecw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD conectar a la bd de mongo*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Conexion exitosa con la BD")

	return client
}

/*ChecarConnection es el ping de prueba a la bd*/
func ChecarConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}
