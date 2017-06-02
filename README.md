# Description
Go: Check access by IP address

# Install and Usage

Install the package with:

```bash
go get github.com/snowirbis/ipauth
```

Import it with:

```go
import "github.com/snowirbis/ipauth"
```

and use `ipauth` as the package name inside the code.

# Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/snowirbis/ipauth"
)

func main() {

	var nets = []string{"10.0.1.0/24", "127.0.0.1/32", "10.0.0.2/32"}

	// allowed
	var client string = "10.0.0.2"

	// denied
	// var client string = "10.0.0.2"

	auth, err := ipauth.IPAuth(nets)
	if err != nil {
		fmt.Println(err)
	}

	allowed, err := auth.Allowed(client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if allowed == true {
		fmt.Println(client, "allowed")

	} else {
		fmt.Println(client, "denied")
	}

}
```

# Based on 
Technique "clientipauth_httpserver.go" Copyright (c) 2016 Mark LaPerriere (MIT License)
https://gist.github.com/marklap/71856411231463a8a16f9d55abc37330
