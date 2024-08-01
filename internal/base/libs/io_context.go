package libs

type IOContext interface {
}

// FIXME: 型を安全に利用できるように修正してください.
type IOResult struct {
	V interface{}
}

type IOContextFactory interface {
	Transaction(f func(io IOContext) (*IOResult, error)) (IOResult, error)
	ReadOnly(f func(io IOContext) (*IOResult, error)) (IOResult, error)
}
