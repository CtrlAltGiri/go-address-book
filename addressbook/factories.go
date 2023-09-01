package addressbook

type AddressBookFactory struct {
	addressbookinstance IAddressBook
	hasherInstance      Hasher
}

func (f *AddressBookFactory) GetAddressBookInstance() IAddressBook {
	if f.addressbookinstance != nil {
		return f.addressbookinstance
	}

	return AddressBook{hasher: f.GetHasher(), cache: FileHandler{}}
}

func (f *AddressBookFactory) GetHasher() Hasher {
	if f.hasherInstance != nil {
		return f.hasherInstance
	}

	return &SHA1Hasher{}
}
