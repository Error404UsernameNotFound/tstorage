package tstorage

type walOperation byte

const (
	operationInsert walOperation = iota
)

// wal is a write-ahead log.
// See more: https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html
type wal interface {
	append(entry walEntry) error
}

// walEntry is an entry in the write-ahead log.
type walEntry struct {
	operation walOperation
	rows      []Row
}
