# cachet

[![Build Status](https://travis-ci.org/andygrunwald/cachet.svg?branch=master)](https://travis-ci.org/andygrunwald/cachet)
[![GoDoc](https://godoc.org/github.com/andygrunwald/cachet?status.svg)](https://godoc.org/github.com/andygrunwald/cachet)

[Go(lang)](https://golang.org/) client library for [Cachet (open source status page system)](https://cachethq.io/).

## Features

* Full API support
	* Components
	* Incidents
	* Metrics
	* Subscribers
* Various authentification methods (Basic Auth and Token based)
* Fully tested

## Installation

It is go gettable

    $ go get github.com/andygrunwald/cachet

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/andygrunwald/cachet
    $ go test -v ./...

## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/cachet) for a detailed API description.

### Authentication

Cachet supports [two different ways](https://docs.cachethq.io/docs/api-authentication) for authentication: BasicAuth and API Token.
Both are supported by this library.

For BasicAuth you need to call the AuthenticationService and apply your email address and your password:

```go
client.Authentication.SetBasicAuth("test@test.com", "test123")
```

To use the API Token way, you do nearly the same but use the `SetTokenAuth` function:

```go
client.Authentication.SetTokenAuth("MY-SECRET-TOKEN")
```

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/cachet#pkg-examples).

### Ping

Call the [API test endpoint](https://docs.cachethq.io/docs/ping). Example without error handling.
Full example available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/cachet#pkg-examples).

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/cachet"
)

func main() {
	client, _ := cachet.NewClient("https://demo.cachethq.io/", nil)
	pong, resp, _ := client.General.Ping()

	fmt.Printf("Result: %s\n", pong)
	fmt.Printf("Status: %s\n", resp.Status)

	// Output: Result: Pong!
	// Status: 200 OK
}
```

### Create a new component

Calling [/components](https://docs.cachethq.io/docs/components). Example without error handling.
Full example available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/cachet#pkg-examples).

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/cachet"
)

func main() {
	client, _ := cachet.NewClient("https://demo.cachethq.io/", nil)
	client.Authentication.SetBasicAuth("test@test.com", "test123")

	component := &cachet.Component{
		Name:        "Beer Fridge",
		Description: "Status of the beer fridge in the kitchen",
		Status:      cachet.ComponentStatusOperational,
	}
	newComponent, resp, _ := client.Components.Create(component)

	fmt.Printf("Result: %s\n", newComponent.Name)
	if newComponent.ID > 0 {
		fmt.Println("ID > 0!")
	}
	fmt.Printf("Status: %s\n", resp.Status)

	// Output: Beer Fridge
	// ID > 0!
	// Status: 200 OK
}
```

## Supported versions

Tested with [v1.2.1](https://github.com/cachethq/Cachet/releases/tag/v1.2.1) of Cachet.
It may works with older and / or newer versions.
Newer versions will be supported. Older versions not.

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).

## Contribution and Contact

Every kind of contribution is welcome!

If you've found a bug, a typo, have a question or a want to request new feature, please [report it as a GitHub issue](https://github.com/andygrunwald/cachet/issues).

For other queries, i'm available on Twitter ([@andygrunwald](https://twitter.com/andygrunwald)).