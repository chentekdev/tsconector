package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chentekdev/tsconector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InventarioSincronizaAll(productos []models.Producto) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	db := MongoCN.Database("tsconector")
	col := db.Collection("productos")

	modelos := []mongo.WriteModel{}

	for i := 0; i < len(productos); i++ {
		operacion := mongo.NewUpdateOneModel()
		operacion.SetFilter(bson.M{"idProveedor": productos[i].IDP})
		operacion.SetUpdate(bson.M{
			"$set": productos[i],
		})
		operacion.SetUpsert(true)
		modelos = append(modelos, operacion)
	}

	bulkOptions := options.BulkWriteOptions{}
	bulkOptions.SetOrdered(true)

	res, err := col.BulkWrite(ctx, modelos, &bulkOptions)

	if err != nil {
		return "Error" + err.Error(), false, err
	}

	return fmt.Sprint("Exito ", res.UpsertedCount, " registros insertados, ", res.ModifiedCount, " registros actualizados"), true, nil
}
