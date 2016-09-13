A quick n' dirty Go port of [py-webrtcvad](https://github.com/wiseman/py-webrtcvad) Voice Activity Detector (VAD).

A VAD classifies a piece of audio data as being voiced or unvoiced. It can be useful for telephony and speech recognition.

The VAD that Google developed for the WebRTC project is reportedly one of the best available, being fast, modern and free.

Usage
-----

Go-get the package. You don't need to have webrtc installed.

```
go get github.com/maxhawkins/go-webrtcvad
```

Feed raw audio samples into the VAD:

```go
reader, err := wav.NewReader("test.wav")
if err != nil {
    log.Fatal(err)
}

vad, err := webrtcvad.New()
if err != nil {
    log.Fatal(err)
}

if err := vad.SetMode(2); err != nil {
    log.Fatal(err)
}

rate := 32000 // kHz
frame := make([]byte, 320*2)

if ok := vad.ValidRateAndFrameLength(rate, len(frame)); !ok {
    log.Fatal("invalid rate or frame length")
}
for {
    _, err := io.ReadFull(reader, frame)
    if err == io.EOF || err == io.ErrUnexpectedEOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }

    active, err := vad.Process(rate, frame)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(active)
}
```
