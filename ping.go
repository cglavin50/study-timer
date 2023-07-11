package main

import (
	"os"
	"time"

	"github.com/hajimehoshi/oto/v2"

	"github.com/hajimehoshi/go-mp3"
)

func ping() error {
	f, err := os.Open("./ping.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	for {
		time.Sleep(time.Millisecond)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
