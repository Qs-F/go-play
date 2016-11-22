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
	// f := color.New(color.FgWhite).Add(color.BgGreen)
	// f.Printf("SUCCESS")
	// fmt.Println(" ", s)
	fmt.Fprintf(os.Stdout, s)
}

func OutputErr(s string) {
	f := color.New(color.FgWhite).Add(color.BgRed).Add(color.Bold)
	f.Printf(" ERROR ")
	fmt.Print(" ")
	// fmt.Println(" ", s)
	fmt.Fprintf(os.Stderr, s)
	fmt.Println()
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
	err = os.Mkdir(rand, 0755)
	if err != nil {
		return "", err
	}
	return srcpath + "/" + rand, nil
}

func GiveDescribe(path string) error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return errors.New("Not found $GOPATH")
	}
	srcpath := filepath.Join(gopath, "src", "localhost", "tmp")
	if srcpath == "" {
		return errors.New("Couldn't compose to change directory `$GOPATH/src/localhost/tmp` query.")
	}
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	_, err = os.Stat(".describe")
	if err == nil {
		return err
	}
	_, err = os.Create(".describe")
	if err != nil {
		return err
	}
	return nil
}
