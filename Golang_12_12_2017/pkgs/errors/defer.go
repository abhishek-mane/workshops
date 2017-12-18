package errors

import (
	"io"
	"os"
)

// defer
// - executes in LIFO(stack) order
func CopyFile(dstName, srcName string) (written int64, err error) {

	src, err := os.Open(srcName)
	// defer src.Close()
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	// defer dst.Close()
	if err != nil {
		return
	}

	return io.Copy(dst, src)
}
