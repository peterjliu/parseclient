package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"parse.com/client"
)

type GameScore struct {
	Score      int    `json:"score"`
	Id         string `json:"objectId"`
	Created    string `json:"createdAt"`
	PlayerName string `json:"playerName"`
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
	var g GameScore
	err = c.GetObj("GameScore", "U05bpEC98w", &g)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%+v\n", g)

	m, err := c.GetObjMap("GameScore", "U05bpEC98w")
	if err != nil {
		fmt.Print(err)
	}
	// numbers are unmarshalled to float64
	score := m["score"].(float64)
	fmt.Println(score)
}
