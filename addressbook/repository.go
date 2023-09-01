package addressbook

import (
	"os"
	"strings"
)

type TieredCache interface {
	GetData(index *string, hash *string, entity SearchEntity) *[]Contact
	SetData(contact *Contact, hash *string)
	SetLink(linkValue *string, linkHash string, hashValue string, entity SearchEntity)
	GetNextTier() TieredCache
}

type FileHandler struct {
}

func (f FileHandler) GetData(identity *string, hash *string, entity SearchEntity) *[]Contact {

	directory := getSearchEntityFileMap()[entity]
	if entity == FullName {
		val := make(map[string]struct{})
		val[*hash] = struct{}{}
		return f.SearchFiles(val, identity, entity)
	}

	return f.SearchIndex(identity, hash, &directory, entity)
}

func (f FileHandler) SetData(contact *Contact, hash *string) {
	dataToWrite := contact.EncodeToString()
	file, err := os.OpenFile(BASEDIR+*hash, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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
}

func (f FileHandler) SetLink(linkValue *string, fileHash string, hashValue string, entity SearchEntity) {
	file, err := os.OpenFile(getSearchEntityFileMap()[entity]+fileHash, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	if hashValue[len(hashValue)-1] != '\n' {
		hashValue += "\n"
	}

	_, err = file.WriteString(hashValue)
	if err != nil {
		panic(err)
	}
}

func (f FileHandler) GetNextTier() TieredCache {
	return nil
}

func (f *FileHandler) SearchIndex(value *string, hash *string, directory *string, entity SearchEntity) *[]Contact {
	values, _ := os.ReadFile(*directory + *hash)
	nameFileMap := strings.Split(string(values), "\n")

	filesToSearch := make(map[string]struct{})
	for _, nameFile := range nameFileMap {
		if len(nameFile) == 0 {
			continue
		}

		nameFileString := strings.Split(nameFile, SEPARATOR)
		if strings.ToLower(nameFileString[0]) == strings.ToLower(*value) {
			var exists = struct{}{}
			_, exist := filesToSearch[nameFileString[1]]
			if !exist {
				filesToSearch[nameFileString[1]] = exists
			}
		}
	}

	return f.SearchFiles(filesToSearch, value, entity)
}

func (f *FileHandler) SearchFiles(filesToSearch map[string]struct{}, valueToSearch *string, entity SearchEntity) *[]Contact {
	var contacts []Contact = nil
	for key, _ := range filesToSearch {
		fileBytes, _ := os.ReadFile(BASEDIR + key)
		fileStrings := strings.Split(string(fileBytes), "\n")
		for _, fileLine := range fileStrings {
			if len(fileLine) == 0 {
				continue
			}

			c := *DecodeContact(fileLine)
			if c.Equals(*valueToSearch, entity) {
				contacts = append(contacts, c)
			}
		}
	}

	return &contacts
}
