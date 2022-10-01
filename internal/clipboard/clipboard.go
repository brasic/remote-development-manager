package clipboard

type Clipboard interface {
	Copy([]byte)
	Paste() []byte
}
