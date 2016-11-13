package play

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Qs-F/gondom"
)

func Compose() (string, error) {
	gopath := os.Getenv("GOPATH")
	srcpath := filepath.Join(gopath, "src", "localhost", "tmp")
	err := os.Chdir(srcpath)
	if err != nil {
		return "", err
	}
IfExist:
	rand := gondom.Make(5, time.Now().UnixNano())
	_, err = os.Stat(rand)
	if err != nil {
		goto IfExist
	}
	err = os.Mkdir(rand, 700)
	if err != nil {
		return "", err
	}
	return srcpath + "/" + rand, nil
}
