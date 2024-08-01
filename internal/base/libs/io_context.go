package libs

type IOContext interface {
}

type IOResult struct {
	V interface{}
}

type IOContextFactory interface {
	Transaction(f func(io IOContext) (*IOResult, error)) (IOResult, error)
	ReadOnly(f func(io IOContext) (*IOResult, error)) (IOResult, error)
}
