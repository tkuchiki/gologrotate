package gologrotate

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Rotate(w *bufio.Writer, f *os.File, filename, suffix string, flag int, perm os.FileMode) (newWriter *bufio.Writer, err error) {
	oldWriter := w
	oldFile := f

	if suffix == "" {
		suffix = time.Now().Format("20060102")
	}

	err = os.Rename(filename, fmt.Sprintf("%s-%s", filename, suffix))
	if err != nil {
		return oldWriter, err
	}

	f, err = os.OpenFile(filename, flag, perm)
	newWriter = bufio.NewWriter(f)

	oldWriter.Flush()
	oldFile.Close()

	return newWriter, err
}
