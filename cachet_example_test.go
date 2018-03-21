package cachet_test

import (
	"fmt"

	"github.com/andygrunwald/cachet"
)

func ExampleGeneralService_Ping() {
	client, err := cachet.NewClient("https://demo.cachethq.io/", nil)
	if err != nil {
		panic(err)
	}

	pong, resp, err := client.General.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", pong)
	fmt.Printf("Status: %s\n", resp.Status)

	// Output: Result: Pong!
	// Status: 200 OK
}

func ExampleComponentsService_Get() {
	client, err := cachet.NewClient("https://demo.cachethq.io/", nil)
	if err != nil {
		panic(err)
	}

	comp, resp, err := client.Components.Get(1)
	if err != nil {
		panic(resp)
	}

	fmt.Printf("Result: %s (ID: %d)\n", comp.Name, comp.ID)
	fmt.Printf("Status: %s\n", resp.Status)

	// Output: Result: API (ID: 1)
	// Status: 200 OK
}
