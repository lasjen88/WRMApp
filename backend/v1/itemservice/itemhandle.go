package itemservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	log "github.com/sirupsen/logrus"
)

//ItemHandle Rest handle for items
type ItemHandle struct {
	ItemCollection mongo.ItemCollection
}

//GetItems is a service handle for getting all items in the database
func (handle *ItemHandle) GetItems(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	items, err := handle.ItemCollection.GetAllItems()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the items"))
		log.Error(err)
		return
	}
	json.NewEncoder(writer).Encode(items)
}

//CreateItem is a service handle inserting a new item into the database
func (handle *ItemHandle) CreateItem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	var item models.Item
	decodeErr := json.NewDecoder(request.Body).Decode(&item)
	if decodeErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Could not parse request body"))
		log.Error(decodeErr)
		return
	}
	databaseErr := handle.ItemCollection.PutItem(item)
	if databaseErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not write to database"))
		log.Error(databaseErr)
		return
	}
	log.Info("Saved item [" + item.ItemName + "] to database.")
}
