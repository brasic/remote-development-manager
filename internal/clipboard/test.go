package clipboard

type TestClipboard struct {
	Buffer []byte
}

func (tc *TestClipboard) Copy(input []byte) {
	tc.Buffer = input
}

func (tc *TestClipboard) Paste() []byte {
	return tc.Buffer
}

func NewTestClipboard() *TestClipboard {
	return &TestClipboard{}
}

var _ Clipboard = (*TestClipboard)(nil)
