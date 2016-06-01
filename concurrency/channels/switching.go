package main

import "fmt"

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// msg := Message{
	// 	To:      []string{"frodo@something.com"},
	// 	From:    "gandalf@toogreat.com",
	// 	Content: "Nerding out",
	// }
	//
	// failedMessage := FailedMessage{
	// 	ErrorMessage:    "Message intercepted by black rider",
	// 	OriginalMessage: Message{},
	// }
	//
	// msgCh <- msg
	// errCh <- failedMessage

	// The first message to be true will be executed
	// so the channels with higher priority should be eveluated first
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("Nothing to do here")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
