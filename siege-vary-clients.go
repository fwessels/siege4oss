package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strconv"
	"time"
)

func Siege(clients int, homeDir string) error {

	app := "siege"

	arg1 := "-c"
	arg2 := strconv.Itoa(clients)
	arg3 := "-q"
	arg4 := "-i"
	arg5 := "-b"
	arg6 := "--time=30s"
	arg7 := `--file=xl-32-30mb.txt` // specify file with presigned-urls here
	arg8 := `--log=siege.log`

	cmd := exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	cmd.Dir = homeDir
	cmb, err := cmd.CombinedOutput()
	fmt.Println(string(cmb))
	if err != nil {
		return errors.New("error")
	}
	return nil
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	for i := 20; i <= 40; i += 1 {
		fmt.Println("Invoking with clients:", i)
		Siege(i, usr.HomeDir)
		fmt.Println("Sleeping ...")
		time.Sleep(30 * time.Second)
	}
}
