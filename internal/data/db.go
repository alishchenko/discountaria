package data

type DB interface {
	New() DB
	Close()
}
