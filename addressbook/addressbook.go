package addressbook

import (
	"os"
	"strings"
)

type IAddressBook interface {
	AddContact(firstName *string, lastName *string, phoneNumber *string, address *string) bool
	SearchContact(identity *string, searchEntity SearchEntity) *[]Contact
}

type AddressBook struct {
	hasher Hasher
}

func (a AddressBook) AddContact(firstName *string, lastName *string, phoneNumber *string, address *string) bool {
	c := Contact{FirstName: *firstName, LastName: *lastName, PhoneNumber: *phoneNumber, Address: *address}

	hashValue := a.HashAndCreateFile(*firstName+*lastName, c.EncodeToString(), BASEDIR)
	a.HashAndLinkToFile(*firstName, hashValue, FIRSTNAMEINDEXDIR)
	a.HashAndLinkToFile(*lastName, hashValue, LASTNAMEINDEXDIR)
	a.HashAndLinkToFile(*phoneNumber, hashValue, PHONEINDEXDIR)
	return true
}

func (a AddressBook) SearchContact(identity *string, searchEntity SearchEntity) *[]Contact {
	if searchEntity == FullName {
		panic("err")
	}

	return a.SearchIndex(identity, getSearchEntityFileMap()[searchEntity])
}

func (a *AddressBook) SearchIndex(valueToSearch *string, directory string) *[]Contact {
	hashToSearch := a.hasher.Hash(*valueToSearch)
	values, _ := os.ReadFile(directory + hashToSearch)
	nameFileMap := strings.Split(string(values), "\n")

	filesToSearch := make(map[string]struct{})
	for _, nameFile := range nameFileMap {
		if len(nameFile) == 0 {
			continue
		}

		nameFileString := strings.Split(nameFile, SEPARATOR)
		if strings.ToLower(nameFileString[0]) == strings.ToLower(*valueToSearch) {
			var exists = struct{}{}
			_, exist := filesToSearch[nameFileString[1]]
			if !exist {
				filesToSearch[nameFileString[1]] = exists
			}
		}
	}

	var contacts []Contact = nil
	for key, _ := range filesToSearch {
		fileBytes, _ := os.ReadFile(BASEDIR + key)
		fileStrings := strings.Split(string(fileBytes), "\n")
		for _, fileLine := range fileStrings {
			if len(fileLine) == 0 {
				continue
			}

			c := *DecodeContact(fileLine)
			if strings.Contains(fileLine, *valueToSearch) {
				contacts = append(contacts, c)
			}
		}
	}

	return &contacts
}

func (a *AddressBook) HashAndCreateFile(valueToHash string, dataToWrite string, directory string) string {
	fileHash := a.hasher.Hash(valueToHash)
	file, err := os.OpenFile(directory+fileHash, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	if dataToWrite[len(dataToWrite)-1] != '\n' {
		dataToWrite += "\n"
	}

	_, err = file.WriteString(dataToWrite)
	if err != nil {
		panic(err)
	}

	return fileHash
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
