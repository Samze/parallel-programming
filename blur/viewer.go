package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"os/exec"
)

//https://github.com/golang-samples/image/blob/master/draw/main.go
func ShowImage(image image.Image) {
	w, _ := os.Create("tmp.png")
	defer w.Close()
	png.Encode(w, image)
	Show(w.Name())
}

func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
