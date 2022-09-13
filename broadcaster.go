package main

type BroadCastMessage struct {
	category  string //one-to-one or bulk or scheduled
	sender    string
	receivers []string
	message   string
	delay     int32
}

type SMSBroadCaster struct {
	broadcast chan BroadCastMessage
}

func Generate_SMSService() *SMSBroadCaster {
	return &SMSBroadCaster{
		broadcast: make(chan BroadCastMessage),
	}
}

func (b *SMSBroadCaster) run() {
	for {
		msg := <-b.broadcast

		switch msg.category {
		case "one-to-one":
			SendOneToOneMessage(msg.sender, msg.receivers[0], msg.message)
		case "bulk":
			SendBulkMessage(msg.sender, msg.receivers, msg.message)
		case "scheduled":
			SendScheduledMessage(msg.sender, msg.receivers, msg.message, msg.delay)
		}
	}
}
