/*
Package cached provides a client for using the Cachet API (https://cachethq.io/).

Construct a new Cachet client, then use the various services on the client to
access different parts of the Cachet API. For example:

	instance := "https://demo.cachethq.io/"
	client, err := cachet.NewClient(instance, nil)

	// Get all components
	components, resp, err := client.Components.GetAll()

The services of a client divide the API into logical chunks and correspond to
the structure of the Cachet API documentation at https://docs.cachethq.io/docs/.

Authentication

The cachet library supports various methods to support the authentication.
This methods are combined in the AuthenticationService that is available at client.Authentication.

One way is an authentication via HTTP BasicAuth:

	instance := "https://demo.cachethq.io/"
	client, err := cachet.NewClient(instance, nil)

	client.Authentication.SetBasicAuth("test@test.com", "test123")

	component := &cachet.Component{
		Name:        "Beer Fridge",
		Description: "Status of the beer fridge in the kitchen",
		Status:      cachet.ComponentStatusOperational,
	}
	newComponent, resp, err := client.Components.Create(component)

	fmt.Printf("Result: %s\n", newComponent.Name)
	// Result: Beer Fridge

The other way is the API Token by Cachet:

	instance := "https://demo.cachethq.io/"
	client, err := cachet.NewClient(instance, nil)

	client.Authentication.SetTokenAuth("MY-SECRET-TOKEN")

	// ... your action here

Additionally when creating a new client, pass an http.Client that supports further actions for you.
For more information regarding authentication have a look at the Cachet documentation:
https://docs.cachethq.io/docs/api-authentication

*/
package cachet
