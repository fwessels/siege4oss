package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("Initial delay")
	time.Sleep(40 * time.Second)

	for i := 20; i <= 40; i += 1 {
		t := time.Now()
		fmt.Println("Dropping cache...")
		ioutil.WriteFile("/proc/sys/vm/drop_caches", []byte("3"), 0644)
		fmt.Println("Caches dropped.")
		tAfter := time.Now()
		elapsed := tAfter.Sub(t)
		fmt.Println(elapsed)
		sleepTime := 60 * time.Second - elapsed
		fmt.Println("Sleeping for ...", sleepTime)
		time.Sleep(sleepTime)
	}
}
