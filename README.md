Go Trello client
----------------
This [Go](http://golang.org/) package implements the [Trello](http://www.trello.com/) [API](http://trello.com/api).

Example
-------

```
package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
)

func main() {
	client, _ := trello.NewClient()

	user, _ := client.Member("user")
	
	fmt.Println("Full name of @user: " + user.FullName)
	fmt.Println("Avatar url of @user: " + user.AvatarUrl())
}
```

Influenced by
-------------
- [fsouza/go-dockerclient](https://github.com/fsouza/go-dockerclient)
- [jeremytregunna/ruby-trello](https://github.com/jeremytregunna/ruby-trello)

License
-------
Go Trello client is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).