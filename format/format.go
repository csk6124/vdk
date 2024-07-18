package format

import (
	"github.com/csk6124/vdk/av/avutil"
	"github.com/csk6124/vdk/format/aac"
	"github.com/csk6124/vdk/format/flv"
	"github.com/csk6124/vdk/format/mp4"
	"github.com/csk6124/vdk/format/rtmp"
	"github.com/csk6124/vdk/format/rtsp"
	"github.com/csk6124/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
