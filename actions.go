package main

import (
  "fmt"
	"github.com/gljubojevic/gocron"
	"github.com/nathan-osman/go-rpigpio"
	"log"
  "os/exec"
	"time"
)

func scheduleActions(events []Event) {
  log.Println("SCHEDULE: Clearing existing schedule.")
	gocron.Clear()
  log.Println("Scheduling:")
	for _, event := range events {
		if event.Action == "open" {
			log.Printf("opening at %s for %d.", event.Time, event.Minutes)
			gocron.Every(1).Day().At(event.Time).Do(open, event.Minutes)
		} else if event.Action == "close" {
			log.Printf("closing at %s for %d.", event.Time, event.Minutes)
			gocron.Every(1).Day().At(event.Time).Do(close, event.Minutes)
		}
	}
  log.Println("SCHEDULE: Saving schedule.")
	<-gocron.Start()
}

func open(mins int) {
  log.Printf("OPENING: Starting to open for %d minutes.", mins)
	p, err := rpi.OpenPin(20, rpi.OUT)
	if err != nil {
		panic(err)
	}
	p.Write(rpi.LOW)
	time.Sleep(time.Duration(mins) * time.Minute)
	p.Write(rpi.HIGH)
	p.Close()
  log.Println("OPENING: Finished")
}

func close(mins int) {
  log.Printf("CLOSING: Starting to close for %d minutes.", mins)
	p, err := rpi.OpenPin(21, rpi.OUT)
	if err != nil {
		panic(err)
	}
	p.Write(rpi.LOW)
	time.Sleep(time.Duration(mins) * time.Minute)
	p.Write(rpi.HIGH)
	p.Close()
  log.Println("CLOSING: Finished")
}

func setDate(t string, d string) {
  log.Printf("TIME: Setting the time to %s %s.", t, d)
  cmd := fmt.Sprintf("sudo date -s '%s %s'", d, t)
  log.Println("executing: ", cmd)
  out, err := exec.Command("/bin/bash", "-c", cmd).Output()
  if err != nil {
    checkErr(err)
  }
  log.Printf("TIME: The time has been set to: %s", out)
}
