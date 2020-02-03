package ee

type Transaction interface {
	Commit() error
	Rollback() error
}
