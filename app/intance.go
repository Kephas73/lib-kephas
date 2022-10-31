package app

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

var Instance InstanceApp

// InstanceApp type;
type InstanceApp interface {
	InitializeApp() error // load module
	RunApp() error        // listen port
	DestroyApp(signal syscall.Signal)
}

var ch = make(chan os.Signal, 1)

func StartApp(instance InstanceApp) {

	// Seeding for random
	rand.Seed(time.Now().UnixNano())
	// Recovery error handler
	defer func() {
		if f := recover(); f != nil {
			if err, ok := f.(error); ok {
				fmt.Println(err)
				debug.PrintStack()
			}
		}
	}()

	if instance == nil {
		err := fmt.Errorf("instance is nil")
		panic(err)
	}

	err := instance.InitializeApp()
	if err != nil {
		panic(err)
	}

	// Set global instance
	Instance = instance

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	go instance.RunApp()

	s2 := <-ch
	if i, ok := s2.(syscall.Signal); ok {
		instance.DestroyApp(i)
	} else {
		instance.DestroyApp(syscall.SIGQUIT)
	}

	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func StopApp() {
	ch <- syscall.SIGQUIT
}
