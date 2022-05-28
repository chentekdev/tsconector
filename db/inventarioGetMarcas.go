package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InventarioGetMarcas() (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("tsconector")
	col := db.Collection("productos")

	projectStage := primitive.D{
		{Key: "$project", Value: primitive.D{
			{Key: "marca", Value: 1},
		}},
	}

	groupbyStage := primitive.D{
		{Key: "$group", Value: primitive.D{
			{Key: "_id", Value: "$marca"},
		}}}

	res, err := col.Aggregate(ctx, mongo.Pipeline{projectStage, groupbyStage})

	if err != nil {
		return "No se pudo ejecutar" + err.Error(), false, err
	}

	type MarcaStruct struct {
		Nombre string `json:"nombre" bson:"_id"`
	}

	resultados := []MarcaStruct{}

	err = res.All(ctx, &resultados)

	if err != nil {
		return err.Error(), false, err
	}

	modelos := []mongo.WriteModel{}

	for i := 0; i < len(resultados); i++ {
		operacion := mongo.NewUpdateOneModel()
		operacion.SetFilter(bson.M{"nombre": resultados[i].Nombre})
		operacion.SetUpdate(bson.M{
			"$set": resultados[i],
		})
		operacion.SetUpsert(true)
		modelos = append(modelos, operacion)
	}

	bulkOptions := options.BulkWriteOptions{}
	bulkOptions.SetOrdered(true)

	col = db.Collection("marcas")
	resBulk, err := col.BulkWrite(ctx, modelos, &bulkOptions)

	if err != nil {
		return "Error:" + err.Error(), false, err
	}
	return fmt.Sprint("Exito se añadieron ", resBulk.UpsertedCount, " Marcas y se actualizaron ", resBulk.ModifiedCount, " marcas"), true, nil
}
