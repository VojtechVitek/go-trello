Golang Trello API client
------------------------
go-trello is a [Go](http://golang.org/) client package for accessing the [Trello](http://www.trello.com/) [API](http://trello.com/api).

<a href="http://golang.org"><img alt="Go package" src="https://golang.org/doc/gopher/pencil/gopherhat.jpg" width="20%" /></a>
<a href="http://trello.com"><img src="https://d2k1ftgv7pobq7.cloudfront.net/meta/p/res/images/c13d1cd96a2cff30f0460a5e1860c5ea/header-logo-blue.svg" style="height: 80px; margin-bottom: 2em;"></a>

[![GoDoc](https://godoc.org/github.com/VojtechVitek/go-trello?status.png)](https://godoc.org/github.com/VojtechVitek/go-trello)
[![Travis](https://travis-ci.org/VojtechVitek/go-trello.svg?branch=master)](https://travis-ci.org/VojtechVitek/go-trello)

Example
-------

```
package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
)

func main() {
	// New Trello Client
	trello, err := trello.NewAuthClient("application-key", "token")
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member("trello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trello Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
```

prints

```
Trello
* How to Use Trello for Android (https://trello.com/b/9dnaRkNt)
   -  Getting Started
      +  Welcome to Trello! This is a card.
      +  Tap on a card to open it up.
      +  Color-coded labels can be used to classify cards.
      +  Create as many cards as you want. We've got an unlimited supply!
      +  Here is a picture of Taco, our Siberian Husky.
   -  Diving In
      +  Tap and hold this card to drag it to another list.
      +  Tap on the board name to view other sections.
      +  Make as many lists and boards as you need. We'll make more!
   -  Mastering Trello
      +  Finished with a card? Drag it to the top of the board to archive it.
      +  You can reorder lists too.
      +  Invite team members to collaborate on this board.
   -  More Info
      +  Want updates on new features?
      +  You can also view your boards on trello.com
```

Influenced by
-------------
- [fsouza/go-dockerclient](https://github.com/fsouza/go-dockerclient)
- [jeremytregunna/ruby-trello](https://github.com/jeremytregunna/ruby-trello)

License
-------
Licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
