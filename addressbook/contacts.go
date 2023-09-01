package addressbook

import "strings"

type Contact struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
}

func (c *Contact) EncodeToString() string {
	return strings.Join([]string{c.FirstName, c.LastName, c.PhoneNumber, c.Address}, SEPARATOR)
}

func DecodeContact(s string) *Contact {
	splits := strings.Split(s, SEPARATOR)
	c := Contact{FirstName: splits[0], LastName: splits[1], PhoneNumber: splits[2], Address: splits[3]}
	return &c
}

func (c *Contact) Equals(value string, searchEntity SearchEntity) bool {
	if searchEntity == FirstName {
		return c.FirstName == value
	}

	if searchEntity == LastName {
		return c.LastName == value
	}

	if searchEntity == PhoneNumber {
		return c.PhoneNumber == value
	}

	if searchEntity == FullName {
		return value == (c.FirstName+c.LastName) ||
			value == (strings.Join([]string{c.FirstName, c.LastName}, " "))
	}

	return false
}
