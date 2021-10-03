package format

import (
	"github.com/ogsts/vdk/av/avutil"
	"github.com/ogsts/vdk/format/aac"
	"github.com/ogsts/vdk/format/flv"
	"github.com/ogsts/vdk/format/mp4"
	"github.com/ogsts/vdk/format/rtmp"
	"github.com/ogsts/vdk/format/rtsp"
	"github.com/ogsts/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
