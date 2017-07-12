package format

// Format function
func Format(format string, context Context) (string, error) {
	splitter := NewSplitter(format, context)
	res := ""
	for splitter.Split() {
		res += splitter.Text()
	}
	return res, splitter.Err()
}
