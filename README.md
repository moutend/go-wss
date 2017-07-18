# go-wss

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/moutend/go-wca/blob/master/LICENSE

Go bindings for Windows SpeechSynthesizer (WinRT) API without using CGO.

This package allows you to generate WAV file which contains synthesized voice from plain text or SSML.

## Prerequisites

- Go 1.8 or later
- `go-ole` ([github.com/go-ole/go-ole](https://github.com/go-ole/go-ole))
- `speechsynthesizer.dll` ([github.com/moutend/speechsynthesizer](https://github.com/moutend/speechsynthesizer))

## Example

The example is located in `example` directory. You can download [executable](https://github.com/moutend/go-wca/releases) or build by yourself. For more information, please read the README.md in `example` directory.

## Documentation

`go-wss` provides three functions.

### TextToStream

Generates WAV file which contains synthesized voice from plain text.

### SsmlToStream

Generates WAV file which contains synthesized voice from SSML.

### GetVoices

Returns available voices.

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
