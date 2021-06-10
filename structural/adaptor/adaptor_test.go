package adaptor

import "testing"

func TestAdaptor(t *testing.T) {
	play := PlayerAdaptor{}
	play.play("mp3", "死了都要爱")
	play.play("wma", "嘀嗒")
}
