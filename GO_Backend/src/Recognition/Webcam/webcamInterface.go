package Webcam

import (
	Webcam "Recognition/Webcam/Consumer"

	"github.com/vee2xx/camtron"
)

func MainWebcam() {
	Webcam.StartForwardStreamConsumer()
	camtron.StartCam()
}
