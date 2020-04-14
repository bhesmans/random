Read auvio video with your player ?


Usage:

```
$ go build ./rtbf.go
# Get the link of you want to see e.g.:https://www.rtbf.be/auvio/detail_edition-speciale?id=2623899

$ ./rtbf https://www.rtbf.be/auvio/detail_edition-speciale?id=2623899
(...) json-like datas that you can reuse

# next I use jq (https://stedolan.github.io/jq/) and mpv (https://mpv.io/)

$ mpv $(./rtbf https://www.rtbf.be/auvio/detail_edition-speciale?id=2623899 | jq -r '.urlHlsAes128')
     Video --vid=1 (h264 416x234 25.000fps) (348 kbps)
     Video --vid=2 (h264 854x480 25.000fps) (1364 kbps)
 (+) Video --vid=3 (h264 1280x720 25.000fps) (2295 kbps)
     Audio --aid=1 --alang=en (*) (aac 2ch 48000Hz) (348 kbps)
     Audio --aid=2 --alang=en (*) (aac 2ch 48000Hz) (1364 kbps)
 (+) Audio --aid=3 --alang=en (*) (aac 2ch 48000Hz) (2295 kbps)
AO: [pulse] 48000Hz stereo 2ch float
VO: [gpu] 1280x720 yuv420p
AV: 00:00:04 / 00:20:07 (0%) A-V:  0.000 Cache: 85s/25MB

(...)
