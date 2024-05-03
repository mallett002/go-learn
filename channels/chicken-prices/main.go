package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string) // holds sale we found websites
	var tofuChannel = make(chan string) // holds sale we found websites

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel) // puts sales in chickenChannel
		go checkTofuPrices(websites[i], tofuChannel) // puts sales in tofuChannel
	}

	sendMessage(chickenChannel, tofuChannel)
}

// adds website to channel if its chicken price is lower or equal to max chicken price
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)

		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)

		var tofuPrice = rand.Float32() * 20

		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Printf("\nFound a deal on chicken at %s", <- chickenChannel)
	// select: if receives value from chickenChannel, will set it to var website and run the first case
	select {
		case website := <- chickenChannel:
			fmt.Printf("\nFound a deal on chicken at %s", website)
		case website := <- tofuChannel:
			fmt.Printf("\nFound a deal on tofu at %s", website)
	}
}
