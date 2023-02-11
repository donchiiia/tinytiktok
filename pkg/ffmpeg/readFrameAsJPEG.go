package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
)

func ReadFrameAsJpeg(inFileName string, frameNum int) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(buf)
	if err != nil {
		klog.Errorf("[]Byte can't decode to image. %v", err)
		return nil, err
	}

	frameBuf := new(bytes.Buffer)
	err = jpeg.Encode(frameBuf, img, nil)
	if err != nil {
		klog.Errorf("Cover image can't encode to JPEG. %v", err)
		return nil, err
	}
	return frameBuf.Bytes(), nil
}
