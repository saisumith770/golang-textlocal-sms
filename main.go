package main

func main() {
	sms_service := Generate_SMSService()
	go sms_service.run()

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
