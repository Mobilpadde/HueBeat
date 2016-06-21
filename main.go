package main

import (
	"fmt"
	"regexp"
	"time"
	"math"

	"github.com/tarm/serial"
	"github.com/Collinux/GoHue"
)

var (
	ser     *serial.Port
	light   hue.Light
	err 	error
	highest int
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func findLast(buf []byte) int {
	n := len(buf)
	str := string(buf[:n])

	r, err := regexp.Compile(`(\d+)\D*$`)
	checkErr(err)

	str = r.FindString(str)

	var val int
	fmt.Sscan(str, &val)

	return val
}

func reader() {
	buf := make([]byte, 10)
	_, err := ser.Read(buf)
	checkErr(err)

	val := findLast(buf)
	percent := int(math.Min(math.Max(float64(val) / float64(highest) * float64(100.0), 1), 100))
	fmt.Println(val, highest, percent)

	if val > highest {
		highest = val
	}

	err = light.SetBrightness(percent)

	time.Sleep(1010 * time.Millisecond)
}

func init() {
	bridges, err := hue.FindBridges()
	checkErr(err)

	bridge := bridges[0]

	bridge.Login("jkEgNmQTV4PhCYhvZ2nNblvKT7Brte8r6LtpVdtJ") // Yay, laziness

	light, err = bridge.GetLightByName("Avl :3")
	checkErr(err)

	highest = 1

	light.On()
}

func main() {
	cfg := &serial.Config{Name: "COM3", Baud: 9600}

	ser, err = serial.OpenPort(cfg)
	checkErr(err)

	defer func() {
		light.Off()
	}()

	for {
		reader()
	}
}
