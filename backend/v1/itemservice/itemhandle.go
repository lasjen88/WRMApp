package itemservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	log "github.com/sirupsen/logrus"
)

const (
	//ContentTypeHeaderKey HTTP header key used to set the type of the body content
	ContentTypeHeaderKey = "Content-Type"
	//ContentTypeHeaderValue HTTP header value used to define the type of the body content
	ContentTypeHeaderValue = "application/json"
)

func mockReadAllDatabase() []models.Item {
	items := make([]models.Item, 0)
	items = append(items, models.Item{ItemName: "Adventurer’s Kit", ItemCost: 5})
	items = append(items, models.Item{ItemName: "Backpack", ItemCost: 4})
	items = append(items, models.Item{ItemName: "Cask of beer", ItemCost: 6})
	items = append(items, models.Item{ItemName: "Cask of wine", ItemCost: 9})
	items = append(items, models.Item{ItemName: "Donkey or mule", ItemCost: 25})
	return items
}

func mockWriteToDatabase(item models.Item) error {
	items := make([]models.Item, 0)
	items = append(items, item)
	return nil
}

//GetItems is a service handle for getting all items in the database
func GetItems(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	items := mockReadAllDatabase()
	json.NewEncoder(writer).Encode(items)
}

//CreateItem is a service handle inserting a new item into the database
func CreateItem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	var item models.Item
	decodeErr := json.NewDecoder(request.Body).Decode(&item)
	if decodeErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Could not parse request body"))
		log.Error(decodeErr)
		return
	}
	databaseErr := mockWriteToDatabase(item)
	if databaseErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not write to database"))
		log.Error(databaseErr)
		return
	}
	log.Info("Saved item [" + item.ItemName + "] to database.")
}
