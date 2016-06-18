package main

import (
    "fmt"
    "time"
    "regexp"

    "github.com/tarm/serial"
    "github.com/Collinux/GoHue"
)

var (
    ser *serial.Port
    err error
    light hue.Light
)

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}

func findLast(buf []byte) int{
    n := len(buf)
    str := string(buf[:n])

    r, err := regexp.Compile(`(\d+)\D*$`);
    checkErr(err)

    str = r.FindString(str)

    var val int
    fmt.Sscan(str, &val)

    return val
}

func reader(){
    buf := make([]byte, 10)
    _, err := ser.Read(buf)
    checkErr(err)
    
    val := findLast(buf)
    fmt.Println(val)

    err = light.SetBrightness(val)
    fmt.Println(err)

    time.Sleep(1010 * time.Millisecond)
}

func init(){
    bridges, err := hue.FindBridges()
    checkErr(err)

    bridge := bridges[0]

    bridge.Login("jkEgNmQTV4PhCYhvZ2nNblvKT7Brte8r6LtpVdtJ") // Yay, laziness

    light, err = bridge.GetLightByName("Avl :3")
    checkErr(err)

    light.On()
}

func main(){
    cfg := &serial.Config{Name: "COM3", Baud: 9600}

    ser, err = serial.OpenPort(cfg)
    checkErr(err)

    defer func(){
        light.Off()
    }()

    for {
        reader()
    }
}