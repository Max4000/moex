package main

import (
	"fmt"
	"moex/deals"
)

func main() {
	countRecords := deals.GetDeals()
	fmt.Println(countRecords)
}
