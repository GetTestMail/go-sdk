# go-sdk

Go client for interacting with the GetTestMail API, which provides a simple way to create temporary email addresses and receive emails sent to them.


## Usage

To create a new GetTestMail API client, you need to instantiate the GetTestMail class with your API key. To get an API key, sign up for a free [account](https://gettestmail.com).

```go
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

	// Wait for a message to arrive
	getTestMail, err = c.WaitForMessage(context.Background(), getTestMail.ID)
	if err != nil {
		panic(err)
	}

	log.Print(getTestMail.Message.Text)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details