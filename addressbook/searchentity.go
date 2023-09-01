package addressbook

type SearchEntity int64

const (
	FullName    SearchEntity = 0
	FirstName   SearchEntity = 1
	LastName    SearchEntity = 2
	PhoneNumber SearchEntity = 3
)

func getSearchEntityFileMap() map[SearchEntity]string {
	return map[SearchEntity]string{
		FullName:    BASEDIR,
		FirstName:   FIRSTNAMEINDEXDIR,
		LastName:    LASTNAMEINDEXDIR,
		PhoneNumber: PHONEINDEXDIR,
	}
}
