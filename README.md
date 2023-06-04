# spacetraders_sdk

A hand-written spacetraders SDK for Go


## Usage

```go
package main

import (
    api "github.com/ult-biffer/spacetraders_sdk/api"
)

func main() {
    // email is optional, use if you paid to reserve your Symbol
    email := "test@example.com"
    // Token is automagically stored in the api package for later use
    resp, err := api.Register("BADGER", "COSMIC", &email)

    if err != nil {
        panic(err)
    }

    // Hopefully do something with resp here
    ...

    ships, err := api.ListAllShips()

    if err != nil {
        panic(err)
    }

    for ship := range ships {
        fmt.Printf("%s %s %s", ship.Symbol, ship.Registration.FactionSymbol, ship.Modules[0].Symbol)
    }
}
```

## Notes

I am literally just gaming here. Please don't hate
