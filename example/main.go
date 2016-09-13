package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/cryptix/wav"
	"github.com/maxhawkins/go-webrtcvad"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("usage: example infile.wav")
	}

	filename := flag.Arg(0)

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	wavReader, err := wav.NewReader(file, info.Size())
	if err != nil {
		log.Fatal(err)
	}
	reader, err := wavReader.GetDumbReader()
	if err != nil {
		log.Fatal(err)
	}

	wavInfo := wavReader.GetFile()
	rate := int(wavInfo.SampleRate)
	if wavInfo.Channels != 1 {
		log.Fatal("expected mono file")
	}
	if rate != 32000 {
		log.Fatal("expected 32kHz file")
	}

	vad, err := webrtcvad.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := vad.SetMode(2); err != nil {
		log.Fatal(err)
	}

	frame := make([]byte, 320*2)

	if ok := vad.ValidRateAndFrameLength(rate, len(frame)); !ok {
		log.Fatal("invalid rate or frame length")
	}

	var isActive bool
	var offset int

	report := func() {
		t := time.Duration(offset) * time.Second / time.Duration(rate) / 2
		fmt.Printf("isActive = %v, t = %v\n", isActive, t)
	}

	for {
		_, err := io.ReadFull(reader, frame)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		frameActive, err := vad.Process(rate, frame)
		if err != nil {
			log.Fatal(err)
		}

		if isActive != frameActive || offset == 0 {
			isActive = frameActive
			report()
		}

		offset += len(frame)
	}

	report()
}
