package mongo

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joaomlucio/projeto/api/user/models"
)

var Collection *mgm.Collection
var Context context.Context

func init(){
	mgm.SetDefaultConfig(nil, "API", options.Client().ApplyURI("mongodb://eu:cafe-eh-vida@localhost:27017"))
	Collection = mgm.Coll(&models.User{})
	Context = mgm.Ctx()
}