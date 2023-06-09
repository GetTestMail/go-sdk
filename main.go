package main

import (
	"context"
	"log"

	"github.com/GetTestMail/go-sdk/client"
)

func main() {
	c := client.NewGetTestMailClient("YOUR_API_KEY")

	// Create a new GetTestMail
	getTestMail, err := c.CreateNew(context.Background())
	if err != nil {
		panic(err)
	}

	log.Printf("%v", getTestMail)

	// Wait for a message to arrive
	getTestMail, err = c.WaitForMessage(context.Background(), getTestMail.ID)
	if err != nil {
		panic(err)
	}

	log.Print(getTestMail.Message.Text)
}
