package ns

//go:generate counterfeiter --fake-name Handle . Handle
type Handle interface {
	Close() error
	IsOpen() bool
	Fd() uintptr
}

//go:generate counterfeiter --fake-name NetworkNamespacer . NetworkNamespacer
type NetworkNamespacer interface {
	GetFromPath(string) (Handle, error)
	Set(Handle) error
}
