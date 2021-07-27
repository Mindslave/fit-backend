package error

type Error struct {
	Err		error
	Op 		string

	// Machine-readable error code
	Code 	string

	// Human-readable error message
	Message string
}