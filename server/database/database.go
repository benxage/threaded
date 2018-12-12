package database

// Database represents a database instance. Database types conform to this
// interface to be used as the server's database.
type Database interface {
	Close()
	PrintInfo()
	Insert(interface{}) error
}
