package clipboard

import (
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

type macosClipboard struct{}

func (m macosClipboard) Copy(input []byte) {
	clipboard.Write(clipboard.FmtText, input)
}

func (m macosClipboard) Paste() []byte {
	return clipboard.Read(clipboard.FmtText)
}

var MacosClipboard Clipboard = macosClipboard{}
