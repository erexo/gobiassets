package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/erexo/gobiassets/in"
	"github.com/erexo/gobiassets/out"
)

func main() {
	fmt.Println("HI")

	prices := in.ReadPrices()
	items := out.SaveItems(prices)
	out.SaveMonsters(items)

	out.SaveSageTree()

	verifyPrices(items, prices)

	fmt.Println("BYE")
}

func verifyPrices(items []*out.Item, prices in.Prices) {
	desired := make(map[uint16]struct{})
	for id := range prices {
		desired[uint16(id)] = struct{}{}
	}
	for _, item := range items {
		delete(desired, item.ServerId)
	}
	if len(desired) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString("Omitted Price item ids:\n")
	for id := range desired {
		fmt.Fprintf(&sb, "%d, ", id)
	}
	log.Println(sb.String())
}
