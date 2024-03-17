package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/godbus/dbus/v5"
)

var (
	deviceName string
	deviceMAC  string
)

func init() {
	flag.StringVar(&deviceName, "name", "", "Name of the Bluetooth device")
	flag.StringVar(&deviceMAC, "mac_address", "", "MAC address of the Bluetooth device")
	flag.Parse()

	if deviceName == "" || deviceMAC == "" {
		fmt.Println("Both device name and MAC address must be specified using the -name and -mac_address flags.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	var action string

	fmt.Print("Enter the action to perform (play, pause, next, previous): ")
	fmt.Scanln(&action)

	deviceMACFormatted := strings.ToUpper(strings.Replace(deviceMAC, ":", "_", -1))
	mediaPlayerPath := fmt.Sprintf("/org/bluez/hci0/dev_%s/player0", deviceMACFormatted)

	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to SystemBus: %s\n", err)
		os.Exit(1)
	}

	switch action {
	case "play", "pause", "next", "previous":
		controlMedia(conn, mediaPlayerPath, strings.Title(action))
	default:
		fmt.Println("Invalid action. Use 'play', 'pause', 'next', or 'previous'")
	}
}

func controlMedia(conn *dbus.Conn, mediaPlayerPath, method string) {
	mediaPlayer := conn.Object("org.bluez", dbus.ObjectPath(mediaPlayerPath))
	call := mediaPlayer.Call("org.bluez.MediaPlayer1."+method, 0)
	if call.Err != nil {
		fmt.Fprintf(os.Stderr, "Failed to %s: %s\n", strings.ToLower(method), call.Err)
		return
	}
	fmt.Printf("%s action executed for %s\n", method, deviceName)
}