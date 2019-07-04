package characterservice

import (
	"context"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//CharacterCollection Collection abstraction over the characters saved in the database.
type CharacterCollection struct {
	collection *mongo.Collection
}

//GetCharacter take a character filter and resturns the first character in the collection that matches.
func (c CharacterCollection) GetCharacter(characterFilter interface{}) (*models.Character, error) {
	var result *models.Character
	err := c.collection.FindOne(context.TODO(), characterFilter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
