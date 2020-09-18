package data

type DataRepository interface {
	Save(*Data) error
}
