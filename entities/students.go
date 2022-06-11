package entities

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Students struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Nom    string             `bson:"nom"`
	Prenom string             `bson:"prenom"`
	Age    int                `bson:"age"`
}

func (s *Students) String() string {
	return fmt.Sprintf("Nom : %s, Prenom : %s, Age : %d", s.Nom, s.Prenom, s.Age)
}
