package main

type Repository interface {
	GetData(index string, directory string) string
}

type FileHandler struct {
}

func (f *FileHandler) GetData(index string, directory string) string {
	return ""
}
