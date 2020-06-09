package main

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() (err error)
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWrite interface {
	Write(buf []byte) (n int, err error)
}

type IClose interface {
	Close() (err error)
}
