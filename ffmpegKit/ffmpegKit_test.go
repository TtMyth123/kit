package ffmpegKit

import (
	"fmt"
	"github.com/disintegration/imaging"
	"testing"
)

func TestExampleReadFrameAsJpeg(t *testing.T) {
	reader, e := ExampleReadFrameAsJpeg("./testData/aaa2.mp4", 50)
	if e != nil {
		fmt.Println(e)
	}
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}
	err = imaging.Save(img, "./testData/out1.jpeg")
	if err != nil {
		fmt.Println(err)
	}
}
