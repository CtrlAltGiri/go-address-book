package addressbook

import (
	"os"
)

type IAddressBook interface {
	AddContact(firstName *string, lastName *string, phoneNumber *string, address *string) bool
	SearchContact(identity *string, searchEntity SearchEntity) *[]Contact
}

type AddressBook struct {
	hasher Hasher
	cache  TieredCache
}

func (a AddressBook) AddContact(firstName *string, lastName *string, phoneNumber *string, address *string) bool {
	c := Contact{FirstName: *firstName, LastName: *lastName, PhoneNumber: *phoneNumber, Address: *address}
	hashValue := a.hasher.Hash(*firstName + *lastName)

	a.cache.SetData(&c, &hashValue)
	a.cache.SetLink(firstName, a.hasher.Hash(*firstName), *firstName+SEPARATOR+hashValue, FirstName)
	a.cache.SetLink(lastName, a.hasher.Hash(*lastName), *lastName+SEPARATOR+hashValue, LastName)
	a.cache.SetLink(phoneNumber, a.hasher.Hash(*phoneNumber), *phoneNumber+SEPARATOR+hashValue, PhoneNumber)

	return true
}

func (a AddressBook) SearchContact(identity *string, searchEntity SearchEntity) *[]Contact {
	hashToSearch := a.hasher.Hash(*identity)
	return a.cache.GetData(identity, &hashToSearch, searchEntity)
}

func (a *AddressBook) HashAndLinkToFile(key string, value string, directory string) {
	fileHash := a.hasher.Hash(key)
	file, err := os.OpenFile(directory+fileHash, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	if value[len(value)-1] != '\n' {
		value += "\n"
	}

	_, err = file.WriteString(key + SEPARATOR + value)
	if err != nil {
		panic(err)
	}
}
