package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"parse.com/client"
)

type GameScore struct {
	Score      int    `json:"score,omitempty"`
	ObjectId   string `json:"objectId,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	PlayerName string `json:"playerName,omitempty"`
}

// Assume parse app id and key are in a file separated by a new line
var keyfile = flag.String("parsekeyfile", "parse-keys",
	"location of file containing parse app id and api key")

func main() {
	d, err := ioutil.ReadFile(*keyfile)
	parsekeys := strings.Split(string(d), "\n")

	if len(parsekeys) < 2 {
		fmt.Printf("length %d", len(parsekeys))
		fmt.Printf("%q\n", parsekeys)
		panic("Couldn't get parse keys")
	}

	c := client.New(parsekeys[0], parsekeys[1])

	fmt.Println("Create new object:")
	var r *client.CreateResp
	r, err = c.CreateObj("GameScore", GameScore{PlayerName: "Peter", Score: 100})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *r)

	fmt.Print("\nGet Object\n")
	var g GameScore
	err = c.GetObj("GameScore", r.ObjectId, &g)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", g)

	m, err := c.GetObjMap("GameScore", r.ObjectId)
	if err != nil {
		fmt.Print(err)
	}
	// numbers are unmarshalled to float64
	score := m["score"].(float64)
	fmt.Printf("\nGet object field (score) without specifying struct: %d\n", score)

	// Update object we just created
	var r2 *client.UpdateResp
	r2, err = c.UpdateObj("GameScore", r.ObjectId, GameScore{PlayerName: "Henry"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *r2)

	fmt.Print("\nGet Updated Object\n")
	err = c.GetObj("GameScore", r.ObjectId, &g)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", g)

	err = c.DeleteObj("GameScore", r.ObjectId, GameScore{PlayerName: "Henry"})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted object.")
	}
}
