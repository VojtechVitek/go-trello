/*
Copyright 2014 go-trello authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"os"
	"testing"

	"github.com/VojtechVitek/go-trello"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

var Client *trello.Client
var Board *trello.Board
var err error

func TestManyThings(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("many things", func() {
		g.It("should create a client", func() {
			key := os.Getenv("API_KEY")
			token := os.Getenv("API_TOKEN")
			Client, err = trello.NewAuthClient(key, &token)
			Expect(err).To(BeNil())
		})

		g.It("should get a board", func() {
			Board, err = Client.Board("iZVEfBeQ")
			Expect(err).To(BeNil())
			Expect(Board.Name).To(Equal("go-trello test board"))
		})

		g.It("should list the lists", func() {
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			Expect(lists[0].Name).To(Equal("meta"))
			Expect(lists[1].Name).To(Equal("a list"))
		})

		g.It("should get a card using two different methods", func() {
			card, err := Board.Card("56cdb3e0f7f4609c2b6f15e4")
			Expect(err).To(BeNil())
			Expect(card.Name).To(Equal("a card"))
			sameCard, err := Client.Card("8sB7wile")
			Expect(err).To(BeNil())
			Expect(sameCard.Name).To(Equal("a card"))
			Expect(sameCard.Desc).To(Equal(card.Desc))
		})
	})

}
