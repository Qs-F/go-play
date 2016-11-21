package play

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Qs-F/gondom"
	"github.com/fatih/color"
)

func Output(s string) {
	f := color.New(color.FgWhite).Add(color.BgGreen)
	f.Printf("SUCCESS")
	fmt.Println(" ", s)
}

func OutputErr(s string) {
	f := color.New(color.FgWhite).Add(color.BgRed)
	f.Printf("ERROR")
	fmt.Println(" ", s)
}

func Compose() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", errors.New("Not found $GOPATH")
	}
	srcpath := filepath.Join(gopath, "src", "localhost", "tmp")
	if srcpath == "" {
		return "", errors.New("Couldn't compose to change directory `$GOPATH/src/localhost/tmp` query.")
	}
	err := os.Chdir(srcpath)
	if err != nil {
		return "", err
	}
IfExist:
	rand := gondom.Make(5, time.Now().UnixNano())
	_, err = os.Stat(rand)
	if err == nil {
		goto IfExist
	}
	err = os.Mkdir(rand, 700)
	if err != nil {
		return "", err
	}
	return srcpath + "/" + rand, nil
}
