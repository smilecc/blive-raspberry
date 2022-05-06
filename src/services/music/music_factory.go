package music

type (
	SongDetail struct {
		Id        string
		Name      string
		Url       string
		LocalPath string
	}
)

const (
	NeteaseService = iota
)

type IMusicService interface {
	GetSongById(id string) (*SongDetail, error)
	GetSongByName(name string) (*SongDetail, error)
}

func GetMusicService(service int) IMusicService {
	switch service {
	case NeteaseService:
		return &NeteaseMusicService{}
	default:
		return nil
	}
}
