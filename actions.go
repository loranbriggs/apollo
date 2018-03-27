package main

import (
  "fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/nathan-osman/go-rpigpio"
	"log"
  "os/exec"
	"time"
)

func scheduleActions(events []Event) {
	gocron.Clear()
	for _, event := range events {
		if event.Action == "open" {
			log.Println("opening at", event.Time, "for", event.Minutes)
			gocron.Every(1).Day().At(event.Time).Do(open, event.Minutes)
		} else if event.Action == "close" {
			log.Println("closing at", event.Time, "for", event.Minutes)
			gocron.Every(1).Day().At(event.Time).Do(close, event.Minutes)
		}
	}
	<-gocron.Start()
}

func open(s int) {
  log.Println("opening!")
	p, err := rpi.OpenPin(5, rpi.OUT)
	if err != nil {
		panic(err)
	}
	p.Write(rpi.HIGH)
	time.Sleep(time.Duration(s) * time.Minute)
	p.Write(rpi.LOW)
	p.Close()
}

func close(s int) {
  log.Println("closing!")
	p, err := rpi.OpenPin(12, rpi.OUT)
	if err != nil {
		panic(err)
	}
	p.Write(rpi.HIGH)
	time.Sleep(time.Duration(s) * time.Minute)
	p.Write(rpi.LOW)
	p.Close()
}

func setDate(t string, d string) {
  cmd := fmt.Sprintf("sudo date -s '%s %s'", d, t)
  log.Println("executing: ", cmd)
  out, err := exec.Command("/bin/bash", "-c", cmd).Output()
  if err != nil {
    checkErr(err)
  }
  log.Printf("%s", out)
}
