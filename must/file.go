package must

import (
	"log"
	"os"
	"path"
)

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Open(fn string) *os.File {
	f, err := os.Open(path.Clean(fn))
	Check(err)
	return f
}
