package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ragul28/geo-tracking-redis/redis"
)

// Error messages
var (
	ErrExpired  = errors.New("request expired")
	ErrCanceled = errors.New("request canceled")
)

type RequestDriverTask struct {
	ID        string
	UserID    string
	Lat, Long float64
	DriverID  string
}

func NewReqDriverTask(id, userID string, lat, long float64) *RequestDriverTask {
	return &RequestDriverTask{
		ID:     id,
		UserID: userID,
		Lat:    lat,
		Long:   long,
	}
}

func (r *RequestDriverTask) Run() {

	// Run driver search 30s duration
	ticker := time.NewTicker(time.Second * 30)

	done := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			err := r.validateRequest()
			switch err {
			case nil:
				log.Println(fmt.Sprintf("Search Driver - Request %s for Lat: %f and Lng: %f", r.ID, r.Lat, r.Long))
				go r.doSearch(done)
			case ErrExpired:
				// Notify to user that the request expired.
				sendInfo(r, "Sorry, we did not find any driver.")
				return
			case ErrCanceled:
				log.Printf("Request %s has been canceled. ", r.ID)
				return
			default: // defensive programming: expected the unexpected
				log.Printf("unexpected error: %v", err)
				return
			}

		case _, ok := <-done:
			if !ok {
				sendInfo(r, fmt.Sprintf("Driver %s found", r.DriverID))
				ticker.Stop()
				return
			}
		}
	}
}

func (r *RequestDriverTask) validateRequest() error {
	rClient := redis.GetRedisClient()
	keyValue, err := rClient.Get(r.ID).Result()
	if err != nil {
		// Request has been expired.
		return ErrExpired
	}

	isActive, _ := strconv.ParseBool(keyValue)
	if !isActive {
		// Request has been canceled.
		return ErrCanceled
	}

	return nil
}

func (r *RequestDriverTask) doSearch(done chan struct{}) {
	rClient := redis.GetRedisClient()
	drivers := rClient.SearchDrivers(1, r.Lat, r.Long, 5)
	if len(drivers) == 1 {
		// Driver found
		// Remove driver location, we can send a message to the driver for that it does not send again its location to this service.
		rClient.RemoveDriverLoc(drivers[0].Name)
		r.DriverID = drivers[0].Name
		close(done)
	}

	return
}

// sendInfo this func is only example, you can use another services, websocket or push notification for send data to user.
func sendInfo(r *RequestDriverTask, message string) {
	log.Println("Message to user:", r.UserID)
	log.Println(message)
}
