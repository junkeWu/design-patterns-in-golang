package adaptor

import "fmt"

// 音乐播放新接口
type MusicPlayer interface {
	play(fileType string, fileName string)
}
// 音乐播放旧接口
type ExistPlayer struct {}

func (*ExistPlayer) playMp3(fileName string)  {
	fmt.Println("play mp3: ", fileName)
}
func (*ExistPlayer) playWma(fileName string)  {
	fmt.Println("play wma: ", fileName)
}

type PlayerAdaptor struct {
	// 持有一个旧接口
	existPlayer ExistPlayer
}

// 实现新接口
func (p *PlayerAdaptor) play(fileType string, fileName string)  {
	switch fileType {
	case "mp3":
		p.existPlayer.playMp3(fileName)
	case "wma":
		p.existPlayer.playWma(fileType)
	default:
		fmt.Println("暂时不支持此类型文件播放")
	}
}
