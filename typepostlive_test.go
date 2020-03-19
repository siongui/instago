package instago

import (
	"testing"
)

const testDashManifest = `<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" minBufferTime="PT1.500S" type="static" mediaPresentationDuration="PT0H13M5.000S" maxSegmentDuration="PT0H0M2.999S" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011,http://dashif.org/guidelines/dash264"><Period duration="PT0H13M5.000S"><AdaptationSet segmentAlignment="true" maxWidth="504" maxHeight="896" maxFrameRate="16000/534" par="216:384" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17915875321397632vd" mimeType="video/mp4" codecs="avc1.4D401F" width="216" height="384" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="259189" FBQualityClass="sd" FBQualityLabel="216w"><BaseURL>https://instagram.frmq3-1.fna.fbcdn.net/v/t72.12950-16/10000000_202927840937946_175935913577152512_n.mp4?_nc_ht=instagram.frmq3-1.fna.fbcdn.net&amp;_nc_cat=106&amp;_nc_ohc=RE61iyJJkDgAX9BuTix&amp;oh=0459096b0ffb527e5b75f4424f0d9a94&amp;oe=5E76491E</BaseURL><SegmentBase indexRangeExact="true" indexRange="907-5618" FBFirstSegmentRange="5619-89287" FBSecondSegmentRange="89288-159897"><Initialization range="0-906"/></SegmentBase></Representation><Representation id="17915875330397632v" mimeType="video/mp4" codecs="avc1.4D401F" width="504" height="896" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="1016931" FBQualityClass="sd" FBQualityLabel="504w"><BaseURL>https://instagram.frmq3-1.fna.fbcdn.net/v/t72.12950-16/10000000_659534454857835_2411504066598273024_n.mp4?_nc_ht=instagram.frmq3-1.fna.fbcdn.net&amp;_nc_cat=103&amp;_nc_ohc=C2ZIXMkHaJAAX8JJnaV&amp;oh=f581a49e3bd86ee541f7a384637d69d2&amp;oe=5E766404</BaseURL><SegmentBase indexRangeExact="true" indexRange="907-5618" FBFirstSegmentRange="5619-330851" FBSecondSegmentRange="330852-609737"><Initialization range="0-906"/></SegmentBase></Representation><Representation id="17915875327397632v" mimeType="video/mp4" codecs="avc1.4D401F" width="396" height="704" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="748080" FBQualityClass="sd" FBQualityLabel="396w"><BaseURL>https://instagram.frmq3-1.fna.fbcdn.net/v/t72.12950-16/10000000_1125942734407542_3417473802043392000_n.mp4?_nc_ht=instagram.frmq3-1.fna.fbcdn.net&amp;_nc_cat=106&amp;_nc_ohc=fabLSbux9bUAX-xkIk8&amp;oh=e237ce780352bfb6c8afb1f2ec1db3c7&amp;oe=5E7670FB</BaseURL><SegmentBase indexRangeExact="true" indexRange="907-5618" FBFirstSegmentRange="5619-225821" FBSecondSegmentRange="225822-440103"><Initialization range="0-906"/></SegmentBase></Representation><Representation id="17915875324397632v" mimeType="video/mp4" codecs="avc1.4D401F" width="324" height="576" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="499810" FBQualityClass="sd" FBQualityLabel="324w"><BaseURL>https://instagram.frmq3-1.fna.fbcdn.net/v/t72.12950-16/10000000_497545760920551_7158980222027563008_n.mp4?_nc_ht=instagram.frmq3-1.fna.fbcdn.net&amp;_nc_cat=103&amp;_nc_ohc=yAPEpV5EAFwAX81_S7O&amp;oh=f79a073149bf2cb54c4ec4e4bae6dd27&amp;oe=5E765146</BaseURL><SegmentBase indexRangeExact="true" indexRange="906-5617" FBFirstSegmentRange="5618-169097" FBSecondSegmentRange="169098-306450"><Initialization range="0-905"/></SegmentBase></Representation></AdaptationSet><AdaptationSet segmentAlignment="true" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17915875327397632ad" mimeType="audio/mp4" codecs="mp4a.40.2" audioSamplingRate="44100" startWithSAP="1" bandwidth="50069"><AudioChannelConfiguration schemeIdUri="urn:mpeg:dash:23003:3:audio_channel_configuration:2011" value="1"/><BaseURL>https://instagram.frmq3-1.fna.fbcdn.net/v/t72.12950-16/74825693_1246519412209779_3697639277585235968_n.mp4?_nc_ht=instagram.frmq3-1.fna.fbcdn.net&amp;_nc_cat=103&amp;_nc_ohc=2cBwxyjyL44AX95k11P&amp;oh=46c0b77f220ff767b2ef71d90409b309&amp;oe=5E763AB1</BaseURL><SegmentBase indexRangeExact="true" indexRange="869-5616" FBFirstSegmentRange="5617-18582" FBSecondSegmentRange="18583-31097"><Initialization range="0-868"/></SegmentBase></Representation></AdaptationSet></Period></MPD>`

func ExampleDashManifest(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
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
