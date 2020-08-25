package blop

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type Sound struct {
	path     string
	streamer beep.StreamSeekCloser
	buffer   *beep.Buffer
	format   beep.Format
}

func (s *Sound) load() {
	f, err := os.Open(s.path)
	if err != nil {
		log.Fatal(err)
	}

	extension := filepath.Ext(s.path)
	switch extension {
	case ".wav":
		s.streamer, s.format, err = wav.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
	case ".mp3":
		s.streamer, s.format, err = mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
	case ".flac":
		s.streamer, s.format, err = flac.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("unknown extension ", extension)
	}

	err = speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	s.buffer = beep.NewBuffer(s.format)
	s.buffer.Append(s.streamer)
	err = s.streamer.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var sounds map[string]*Sound

func LoadSound(key, path string) {
	if sounds == nil {
		sounds = make(map[string]*Sound)
	}

	s := Sound{
		path: path,
	}
	s.load()

	sounds[key] = &s
}

func Play(key string) {
	if _, ok := sounds[key]; !ok {
		fmt.Println("sound.Play", key, "doesn't exist")
		return
	}
	s := sounds[key]

	streamer := s.buffer.Streamer(0, s.buffer.Len())
	speaker.Play(streamer)
}
