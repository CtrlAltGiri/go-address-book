package addressbook

func GetAddressBookInstance(config string) IAddressBook {
	if config == "test" {
		return nil
	}

	return AddressBook{hasher: GetHasher(config)}
}

func GetHasher(config string) Hasher {
	if config == "test" {
		return nil
	}

	return &SHA1Hasher{}
}
