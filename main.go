package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

type Sms struct {
	ClientID  string
	OneToOne  int
	Bulk      int
	Scheduled int
	Month     string
}

func CreateSmsPack_SingleUser(ds *datastore.Client, ctx context.Context, clientID string) {
	key := datastore.NameKey("Sms", clientID, nil)
	time := time.Now()
	sms := Sms{
		ClientID:  clientID,
		Month:     time.Format("01-2006"),
		OneToOne:  0,
		Bulk:      0,
		Scheduled: 0,
	}

	data, err := ds.Put(ctx, key, &sms)
	if err != nil {
		log.Printf("DATASTORE::CREATE: failed to create a new document in datastore")
		log.Print(err)
	} else {
		log.Print(data)
	}
}

func GetSmsPack_SingleUser(ds *datastore.Client, ctx context.Context, clientID string) {
	key := datastore.NameKey("Sms", clientID, nil)
	smsEntity := new(Sms)

	if err := ds.Get(ctx, key, smsEntity); err != nil {
		log.Printf("DATASTORE::GET: failed to get the sms pack for client:%v", clientID)
	} else {
		log.Printf("DATASTORE::GET: successfully retrieved sms pack for client:%v details:%v", clientID, smsEntity)
	}
}

func UpdateSmsPack_SingleUser(ds *datastore.Client, ctx context.Context, clientID string, increment_pack string) {
	key := datastore.NameKey("Sms", clientID, nil)
	smsEntity := Sms{
		ClientID:  clientID,
		OneToOne:  0,
		Bulk:      0,
		Scheduled: 0,
	}

	switch increment_pack {
	case "one-to-one":
		smsEntity.OneToOne++
	case "bulk":
		smsEntity.Bulk++
	case "scheduled":
		smsEntity.Scheduled++
	}

	mutation := datastore.NewUpdate(key, &smsEntity)
	data, err := ds.Mutate(ctx, mutation)
	if err != nil {
		log.Printf("DATASTORE::UPDATE: failed to update the sms pack for client:%v", clientID)
	} else {
		log.Printf("DATASTORE::UPDATE: successfully updated the sms pack for client:%v data:%v", clientID, data)
	}
}

func DeleteSmsPack_SingleUser(ds *datastore.Client, ctx context.Context, clientID string) {
	key := datastore.NameKey("Sms", clientID, nil)
	if err := ds.Delete(ctx, key); err != nil {
		log.Printf("DATASTORE::DELETE: failed to delete the sms pack for client:%v", clientID)
	}
}

func main() {
	sms_service := Generate_SMSService()
	go sms_service.run()

	ctx := context.Background()
	ds, err := datastore.NewClient(ctx, "ontiotechnologies")
	if err != nil {
		log.Fatal("DATASTORE::CONNECTION: failed to connect to datastore", err)
	}

	//release control before quitting
	defer func() {
		ds.Close()
	}()

	// CreateSmsPack_SingleUser(ds, ctx, "1234")
	GetSmsPack_SingleUser(ds, ctx, "1234")
	// UpdateSmsPack_SingleUser(ds, ctx, "1234", "one-to-one")
	// DeleteSmsPack_SingleUser(ds, ctx, "1234")

	// sms_service.broadcast <- BroadCastMessage{
	// 	category:  "one-to-one",
	// 	sender:    "67800",
	// 	receivers: []string{"9180045566"},
	// 	message:   "sup bruv",
	// 	delay:     0,
	// }
	// sms_service.broadcast <- BroadCastMessage{
	// 	category:  "bulk",
	// 	sender:    "67800",
	// 	receivers: []string{"9180045566","91955400026"},
	// 	message:   "sup bruv",
	// 	delay:     0,
	// }
	// sms_service.broadcast <- BroadCastMessage{
	// 	category:  "scheduled",
	// 	sender:    "67800",
	// 	receivers: []string{"9180045566","91363636366"},
	// 	message:   "sup bruv",
	// 	delay:     10,
	// }
}
