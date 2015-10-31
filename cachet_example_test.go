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

func ExampleComponentsService_Create() {
	client, err := cachet.NewClient("https://demo.cachethq.io/", nil)
	if err != nil {
		panic(err)
	}

	client.Authentication.SetBasicAuth("test@test.com", "test123")

	component := &cachet.Component{
		Name:        "Beer Fridge",
		Description: "Status of the beer fridge in the kitchen",
		Status:      cachet.ComponentStatusOperational,
	}
	newComponent, resp, err := client.Components.Create(component)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", newComponent.Name)
	if newComponent.ID > 0 {
		fmt.Println("ID > 0!")
	}
	fmt.Printf("Status: %s\n", resp.Status)

	// Output: Result: Beer Fridge
	// ID > 0!
	// Status: 200 OK
}
