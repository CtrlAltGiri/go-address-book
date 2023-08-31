package main

import (
	"fmt"
	"os"
	"time"

	"github.com/CtrlAltGiri/go-address-book/addressbook"
)

func main() {
	abInstance := addressbook.GetAddressBookInstance("prod")
	startTime := time.Now()
	latency := time.Now().Sub(startTime)

	if len(os.Args) < 4 || len(os.Args) > 6 {
		fmt.Println("Please specify the arguments as required")
	}

	if os.Args[1] == addressbook.ADD && len(os.Args) == 6 {
		abInstance.AddContact(&os.Args[2], &os.Args[3], &os.Args[4], &os.Args[5])
		latency = time.Now().Sub(startTime)
	} else if os.Args[1] == addressbook.SEARCH && len(os.Args) == 4 {
		searchEntity := os.Args[3]
		var contacts *[]addressbook.Contact = nil
		if os.Args[2] == addressbook.NAME {
			contacts = abInstance.SearchName(&searchEntity)

		} else if os.Args[2] == addressbook.PHONE {
			contacts = abInstance.SearchPhone(&searchEntity)
		}

		latency = time.Now().Sub(startTime)
		if contacts != nil {
			for _, contact := range *contacts {
				fmt.Println(contact)
			}
		}
	}

	fmt.Println(latency)
}
