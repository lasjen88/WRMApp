package mongo

import "github.com/globalsign/mgo"

//WrmCollection holds the data collection
type WrmCollection struct {
	DatabaseName, CollectionName string
	Session                      *mgo.Session
	collection                   *mgo.Collection
}

//SetupCollection sets up the database collection
func (w *WrmCollection) SetupCollection() {
	w.collection = w.Session.DB(w.DatabaseName).C(w.CollectionName)
}
