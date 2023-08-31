package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	addressbook := GetAddressBookInstance("prod")
	startTime := time.Now()

	if len(os.Args) < 4 || len(os.Args) > 6 {
		fmt.Println("Please specify the arguments as required")
	}

	if os.Args[1] == ADD && len(os.Args) == 6 {
		addressbook.AddContact(&os.Args[2], &os.Args[3], &os.Args[4], &os.Args[5])
	} else if os.Args[1] == SEARCH && len(os.Args) == 4 {
		searchEntity := os.Args[3]
		if os.Args[2] == NAME {
			contacts := addressbook.SearchName(&searchEntity)
			if contacts != nil {
				// fmt.Println((*contacts)[0])
			}
		} else if os.Args[2] == PHONE {
			contacts := addressbook.SearchPhone(&searchEntity)
			if contacts != nil {
				// fmt.Println((*contacts)[0])
			}
		}
	}

	latency := time.Now().Sub(startTime)
	fmt.Println(latency)
}
