package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CtrlAltGiri/go-address-book/addressbook"
)

func init() {
	os.MkdirAll(addressbook.INDEXDIR, os.ModePerm)
	os.MkdirAll(addressbook.FIRSTNAMEINDEXDIR, os.ModePerm)
	os.MkdirAll(addressbook.PHONEINDEXDIR, os.ModePerm)
	os.MkdirAll(addressbook.LASTNAMEINDEXDIR, os.ModePerm)
}

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

			// Search by full name if it is split by a space
			if len(strings.Split(searchEntity, " ")) == 2 {
				contacts = abInstance.SearchContact(&searchEntity, addressbook.FullName)
			}

			// Search only if results were not found
			if contacts == nil || len(*contacts) == 0 {
				contacts = abInstance.SearchContact(&searchEntity, addressbook.FirstName)
				if len(*contacts) == 0 {
					contacts = abInstance.SearchContact(&searchEntity, addressbook.LastName)
				}
			}
		} else if os.Args[2] == addressbook.PHONE {
			contacts = abInstance.SearchContact(&searchEntity, addressbook.PhoneNumber)
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
