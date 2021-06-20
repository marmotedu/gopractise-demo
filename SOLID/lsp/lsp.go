package lsp

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func Write(w Writer, p []byte) (int, error) {
	return w.Write(p)
}
