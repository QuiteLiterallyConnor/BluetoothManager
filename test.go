package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	fmt.Println("Hello, World!")
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to session bus")

	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/org/freedesktop/DBus"),
		dbus.WithMatchInterface("org.freedesktop.DBus"),
		dbus.WithMatchSender("org.freedesktop.DBus"),
	); err != nil {
		panic(err)
	}

	fmt,Println("Matched signal")

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for v := range c {
		fmt.Println(v)
	}
}