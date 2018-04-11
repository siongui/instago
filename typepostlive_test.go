package instago

import (
	"os"
	"testing"
)

const testDashManifest = `<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" minBufferTime="PT1.500S" type="static" mediaPresentationDuration="PT0H0M9.468S" maxSegmentDuration="PT0H0M2.000S" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011,http://dashif.org/guidelines/dash264"><Period duration="PT0H0M9.468S"><AdaptationSet segmentAlignment="true" maxWidth="396" maxHeight="746" maxFrameRate="16000/528" par="396:746" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17924879014006631v" mimeType="video/mp4" codecs="avc1.4d401f" width="396" height="746" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="836675" FBQualityClass="sd" FBQualityLabel="396w"><BaseURL>https://instagram.fkhh1-1.fna.fbcdn.net/vp/4bc800aa48ed9f3bcc763d3e5d4e48fe/5A8F0D5A/t72.12950-16/27465973_336693050159819_3428588455850934272_n.mp4</BaseURL><SegmentBase indexRangeExact="true" indexRange="899-1026"><Initialization range="0-898"/></SegmentBase></Representation></AdaptationSet><AdaptationSet segmentAlignment="true" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17924879014006631a" mimeType="audio/mp4" codecs="mp4a.40.2" audioSamplingRate="44100" startWithSAP="1" bandwidth="51679"><AudioChannelConfiguration schemeIdUri="urn:mpeg:dash:23003:3:audio_channel_configuration:2011" value="2"/><BaseURL>https://instagram.fkhh1-1.fna.fbcdn.net/vp/c3198573a9e6375f6bd01da3c8312dc4/5A8F0753/t72.12950-16/27486167_191963884733826_1789728450089582592_n.mp4</BaseURL><SegmentBase indexRangeExact="true" indexRange="835-926"><Initialization range="0-834"/></SegmentBase></Representation></AdaptationSet></Period></MPD>`

func ExampleDashManifest(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	rt, err := mgr.GetReelsTray()
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range rt.PostLive.PostLiveItems {
		println(item.GetUsername())
		for _, broadcast := range item.GetBroadcasts() {
			//println(broadcast.GetDashManifest())
			urls, err := broadcast.GetBaseUrls()
			if err != nil {
				t.Error(err)
				return
			}
			println(broadcast.GetPublishedTime())
			for _, url := range urls {
				println(url)
			}
		}
	}
}
