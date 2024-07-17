package ffmpeg

const ResizeCmd = CommonCmd + `-vf "scale=iw*.5:ih*.5" %s`
