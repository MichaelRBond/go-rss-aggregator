package errors

// Handle Simple error handling
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}
