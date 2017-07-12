package format

// Formatter generic formatting piece
type Formatter interface {
	Clarify(string) (Formatter, error)
	Format(string) (string, error)
}
