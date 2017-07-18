// +build windows

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/moutend/go-wss"
)

type FilenameFlag struct {
	Value     string
	Extension string
}

func (f *FilenameFlag) Set(value string) (err error) {
	s := strings.Split(value, ".")
	f.Value = value
	f.Extension = s[len(s)-1]
	return
}

func (f *FilenameFlag) String() string {
	return f.Value
}

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
	return
}

func run(args []string) (err error) {
	var inputFlag FilenameFlag
	var outputFlag FilenameFlag
	var voiceNumberFlag int
	var file []byte
	var stream []byte

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&inputFlag, "i", "input file name (*.txt or *.ssml)")
	f.Var(&outputFlag, "o", "output file name (default: voice.wav)")
	f.IntVar(&voiceNumberFlag, "n", 0, "voice number")
	f.Parse(args[1:])

	if inputFlag.Value == "" {
		err = ListVoices()
		return
	}
	if file, err = ioutil.ReadFile(inputFlag.Value); err != nil {
		return
	}
	if outputFlag.Value == "" {
		outputFlag.Value = "voice.wav"
	}
	if stream, err = CreateSpeechStream(voiceNumberFlag, inputFlag.Extension, string(file[:])); err != nil {
		return
	}
	return ioutil.WriteFile(outputFlag.Value, stream, 0644)
}

func CreateSpeechStream(index int, format, input string) ([]byte, error) {
	if format == "ssml" {
		return wss.SsmlToStream(index, input)
	} else {
		return wss.TextToStream(index, input)
	}
}

func ListVoices() (err error) {
	var vs []wss.VoiceInformation
	if vs, err = wss.GetVoices(); err != nil {
		return
	}
	for i, v := range vs {
		fmt.Printf("%v. (%v) %v\n", i+1, v.Language, v.Name)
	}
	return
}
