package mongo

import (
	"context"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joaomlucio/projeto/api/user/models"
)

var Collection *mgm.Collection
var Context context.Context

func init(){
	err := mgm.SetDefaultConfig(nil, "api", options.Client().ApplyURI("mongodb://user:password@mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	Collection = mgm.Coll(&models.User{})
	Context = mgm.Ctx()
}
