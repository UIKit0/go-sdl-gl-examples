package main

import "flag"

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "1":
		drawWithSdl()
	case "2":
		drawWithSdlGlImmediate()
	default:
		drawWithSdlGlRetained()
	}
}
