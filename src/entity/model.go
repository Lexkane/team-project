package entity

import (
	"reflect"
	"strings"
	"time"

	"../database"
	"github.com/satori/go.uuid"
)

//Flights model in DB
type Flight struct {
	ID             uuid.UUID
	Departure_city string
	Departure      time.Time
	Arrival_city   string
	Arrival        time.Time
	Price          int
}

//Train model  in DB
type Train struct {
	ID             uuid.UUID
	Departure      time.Time
	Arrival        time.Time
	Departure_city string
	Arrival_city   string
	Train_type     string
	Car_type       string
	Price          int
}

var AddToTrip = func(dataID uuid.UUID, tripID uuid.UUID, dataSource interface{}) error {
	_, err := database.DB.Exec(GenerateQueryAdd(dataSource), dataID, tripID)
	return err
}

func GenerateQueryAdd(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	var query = "INSERT INTO trips_" + strings.ToLower(dataType.Name()) + "s" + " (" + strings.ToLower(dataType.Name()) + "_id, trip_id) VALUES ($1, $2)"
	return query
}

func generateQueryGet(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	name := strings.ToLower(dataType.Name())
	pluralName := name + "s"
	var query = "SELECT " + pluralName + ".* FROM " + pluralName + " INNER JOIN trips_" + pluralName + " ON " + pluralName + ".id=trips_" + pluralName + "." + name + "_id AND trips_" + pluralName + ".trip_id=$1"
	return query
}
