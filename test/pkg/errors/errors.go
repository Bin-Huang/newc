package errors

// NoFound ...
type NoFound struct{}

// String ...
func (NoFound) String() string {
	return "no found"
}
