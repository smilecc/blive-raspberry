package services

import (
	"fmt"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
)

func Test() {

	ps := rtmp.NewPushSession(func(option *rtmp.PushSessionOption) {
		option.PushTimeoutMs = 5000
		option.WriteAvTimeoutMs = 10000
		option.WriteBufSize = 0
		option.WriteChanSize = 0
	})

	err := ps.Push("rtmp://live-push.bilivideo.com/live-bvc/?streamname=live_927293_332_c521e483&key=71e3bf02c3502e40750dec010e364bee&schedule=rtmp&pflag=1")
	if err != nil {
		return
	}

	go func() {
		list := [...]string{"D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv", "D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv", "D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv", "D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv", "D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv", "D:\\Temp\\test\\1.flv", "D:\\Temp\\test\\2.flv"}
		for _, url := range list {
			fmt.Println(url)
			flvFilePump := httpflv.NewFlvFilePump(func(option *httpflv.FlvFilePumpOption) {
				option.IsRecursive = false
			})

			tags, _ := httpflv.ReadAllTagsFromFlvFile(url)
			flvFilePump.PumpWithTags(tags, func(tag httpflv.Tag) bool {
				chunks := remux.FlvTag2RtmpChunks(tag)
				err := ps.Write(chunks)
				if err != nil {
					return false
				}
				return true
			})
		}

	}()

	_ = <-ps.WaitChan()
}
