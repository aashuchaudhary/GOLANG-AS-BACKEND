package model

// whwnever use mongodb  it gives us the ids which monogo db generates automatically for us there are not just ids they are just underscore id type just bson structure  they are very similar to  be in visual aspect json but they have some added fields into them.
import (
	"go.mongodb.org/mongo-driver/bson/primitive" //we have to give primitive  that will give all the ids to it
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
