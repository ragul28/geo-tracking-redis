package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ragul28/geo-tracking-redis/model"
	"github.com/ragul28/geo-tracking-redis/redis"
)

// HealthEndpoint
func health(w http.ResponseWriter, req *http.Request) {
	_ = redis.GetRedisClient()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service Healthy"))
}

// Tracking driver
func tracking(w http.ResponseWriter, r *http.Request) {

	var drivers []model.Driver

	rClient := redis.GetRedisClient()

	_ = jsonDecoder(r.Body, w, &drivers)

	log.Println(len(drivers))
	// add drivers location to redis
	for i := 0; i < len(drivers); i++ {
		rClient.AddDriverLoc(drivers[i].Long, drivers[i].Lat, drivers[i].ID)
	}

	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal(&model.MassageStatus{Status: "Drivers Location saved successfully"})
	w.Write(msg)
}

// simple Search drivers frm req point
func search(w http.ResponseWriter, r *http.Request) {

	rClient := redis.GetRedisClient()

	var body model.ReqDriver

	_ = jsonDecoder(r.Body, w, &body)

	// Search in redis
	drivers := rClient.SearchDrivers(body.Limit, body.Lat, body.Long, 15)
	data, err := json.Marshal(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func jsonDecoder(req io.Reader, w http.ResponseWriter, body interface{}) error {

	err := json.NewDecoder(req).Decode(body)
	if err != nil {
		log.Printf("Req Decode Failed: %v", err)
		http.Error(w, "Req Decode Failed", http.StatusInternalServerError)
		return err
	}
	return nil
}
