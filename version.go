package main

//import "gopkg.in/h2non/bimg.v0"
import "github.com/twilkensboopsie/bimg"

const Version = "0.1.22"

type Versions struct {
	ImaginaryVersion string `json:"imaginary"`
	BimgVersion      string `json:"bimg"`
	VipsVersion      string `json:"libvips"`
}

var CurrentVersions = Versions{Version, bimg.Version, bimg.VipsVersion}
