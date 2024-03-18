package main

import (
	"fmt"

	bt "github.com/QuiteLiterallyConnor/BluetoothController"
)

var ScannerListener = func(device bt.Device) {
	fmt.Printf("Name: %s, Device: %+v\n", device.Name, device)
}

func scanner() {
	bt.EnableDebugging()

	bs, err := bt.NewBluetoothScanner(ScannerListener)
	if err != nil {
		fmt.Println("Error initializing Bluetooth Scanner:", err)
		return
	}

	if err = bs.GetManagedDevices(); err != nil {
		fmt.Printf("err getting managed devices: %s\n", err)
	}

	bs.StartScanner()

	for {
	}

}

func main() {
	scanner()
}
