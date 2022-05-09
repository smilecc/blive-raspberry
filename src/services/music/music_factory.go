package music

type (
	SongDetail struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		FileName    string `json:"fileName"`
		Url         string `json:"url"`
		LocalPath   string `json:"localPath"`
		Lrc         string `json:"lrc"`
		SingerName  string `json:"singerName"`
		Duration    int    `json:"duration"`
		AlbumName   string `json:"albumName"`
		AlbumPicUrl string `json:"albumPicUrl"`
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
		return NewNeteaseMusicService()
	default:
		return nil
	}
}
