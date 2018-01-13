package format

type nopeFormatter struct{}

// Clarify ...
func (n nopeFormatter) Clarify(string) (Formatter, error) {
	return n, nil
}

// Format ...
func (nopeFormatter) Format(string) (string, error) {
	return "", nil
}
