package main

import (
  //"github.com/jasonlvhit/gocron"
  "github.com/nathan-osman/go-rpigpio"
  "log"
  "time"
)

func scheduleActions(events []Event, durations []Duration) {
  var openDur, closeDur int
  for _, dur := range durations {
    if dur.Action == "open" {
      openDur = dur.Duration
    } else if dur.Action == "close" {
      closeDur = dur.Duration
    }
  }

  //gocron.Clear()
  for _, event := range events {
    if event.Action == "open" {
      log.Println("opening at", event.Time, "for", openDur)
      //gocron.Every(1).Day().At(event.Time).Do(open, openDur)
    } else if event.Action == "close" {
      log.Println("closing at", event.Time, "for", closeDur)
      //gocron.Every(1).Day().At(event.Time).Do(close, closeDur)
    }
  }
  //<-gocron.Start()
}

func open(s int) {
  p, err := rpi.OpenPin(5, rpi.OUT)
  if err != nil {
    panic(err)
  }
  p.Write(rpi.HIGH)
  time.Sleep(time.Duration(s) * time.Second)
  p.Write(rpi.LOW)
  p.Close()
}

func close(s int) {
  p, err := rpi.OpenPin(12, rpi.OUT)
  if err != nil {
    panic(err)
  }
  p.Write(rpi.HIGH)
  time.Sleep(time.Duration(s) * time.Second)
  p.Write(rpi.LOW)
  p.Close()
}
