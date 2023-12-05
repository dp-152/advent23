package fs

import (
	"bufio"
	"os"
)

func init() {

}

func ReadLines(path string, cancel <-chan struct{}) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		fileHandle, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer fileHandle.Close()

		scanner := bufio.NewScanner(fileHandle)

		for scanner.Scan() {
			select {
			case ch <- scanner.Text():
			case <-cancel:
				return
			}
		}
	}()

	return ch
}

func ReadFile(path string) string {
	contents, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(contents)
}
