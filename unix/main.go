package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep(5 * time.Second)
	}
}

func handleSignal(singal os.Signal) {
	fmt.Println("handleSignal() caught:", singal)
}
