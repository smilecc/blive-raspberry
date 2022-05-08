package music

import (
	"blive/src/database"
	"encoding/json"
	"fmt"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/imroc/req/v3"
	"os"
	"strconv"
)

type NeteaseMusicService struct {
	ApiHost string
	cookie  string
}

type (
	NeteaseMusicData[T interface{}] struct {
		Result T   `json:"result"`
		Data   T   `json:"data"`
		Code   int `json:"code"`
	}
	NeteaseMusicSong struct {
		Id  int    `json:"id"`
		Url string `json:"url"`
	}
	NeteaseMusicLrc struct {
		Code int `json:"code"`
		Lrc  struct {
			Version int    `json:"version"`
			Lyric   string `json:"lyric"`
		} `json:"lrc"`
	}
	NeteaseMusicSearchSong struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	NeteaseMusicSearch struct {
		SongCount int                      `json:"songCount"`
		Songs     []NeteaseMusicSearchSong `json:"songs"`
	}
)

func NewNeteaseMusicService() *NeteaseMusicService {
	service := &NeteaseMusicService{}
	var configs []database.SysConfig
	database.DB.Where("name = 'netease_api_host'").Find(&configs)
	if len(configs) == 0 {
		service.ApiHost = "https://netease-cloud-music-api-ochre-one.vercel.app"
	} else {
		service.ApiHost = configs[0].Value
	}

	database.DB.Where("name = 'netease_cookie'").Find(&configs)
	if len(configs) == 0 {
		service.cookie = ""
	} else {
		service.cookie = configs[0].Value
	}

	return service
}

func (n *NeteaseMusicService) GetSongById(id string) (*SongDetail, error) {
	return n.getMusic(id)
}

func (n *NeteaseMusicService) GetSongByName(name string) (*SongDetail, error) {
	return n.getMusic(name)
}

//getMusic 通过ID或名称查询音乐
func (n *NeteaseMusicService) getMusic(name string) (*SongDetail, error) {
	searchResult := n.searchMusic(name)
	if searchResult.Result.SongCount > 0 {
		firstSong := searchResult.Result.Songs[0]
		song, err := n.getMusicById(firstSong.Id, &firstSong)

		if err != nil {
			return nil, err
		}

		song.Name = firstSong.Name
		return song, err
	}

	return nil, fmt.Errorf("歌曲不存在")
}

//searchMusic 通过关键词搜索音乐
func (n *NeteaseMusicService) searchMusic(keyword string) *NeteaseMusicData[NeteaseMusicSearch] {
	result := &NeteaseMusicData[NeteaseMusicSearch]{}
	resp, err := req.R().
		SetResult(&result).
		SetQueryParam("cookie", n.cookie).
		SetQueryParam("keywords", keyword).
		Get(fmt.Sprintf("%s/search", n.ApiHost))

	if err != nil && resp != nil {
		return nil
	}

	listJson, _ := json.Marshal(result.Result)
	log.Infof("搜索到音乐信息 keyword: %s 列表信息: %s", keyword, listJson)
	return result
}

//getMusicById 通过ID查询音乐
func (n *NeteaseMusicService) getMusicById(id int, searchSong *NeteaseMusicSearchSong) (*SongDetail, error) {
	// 如果没有查询过音乐 则先查询
	if searchSong == nil {
		searchResult := n.searchMusic(strconv.Itoa(id))
		if searchResult.Result.SongCount > 0 {
			searchSong = &searchResult.Result.Songs[0]
		}
	}

	// 通过ID查询音乐链接
	result := &NeteaseMusicData[[]NeteaseMusicSong]{}
	_, err := req.R().
		SetResult(&result).
		SetQueryParam("id", strconv.Itoa(id)).
		SetQueryParam("br", "320000").
		SetQueryParam("cookie", n.cookie).
		Get(fmt.Sprintf("%s/song/url", n.ApiHost))
	if err != nil {
		return nil, err
	}

	listJson, _ := json.Marshal(result.Data)
	log.Infof("获取到音乐信息 Id: %d 列表信息: %s", id, listJson)
	if len(result.Data) == 0 {
		return nil, err
	}

	lrcResult := &NeteaseMusicLrc{}
	_, err = req.R().
		SetResult(&lrcResult).
		SetQueryParam("id", strconv.Itoa(id)).
		SetQueryParam("cookie", n.cookie).
		Get(fmt.Sprintf("%s/lyric", n.ApiHost))

	lrc := ""
	if lrcResult.Code == 200 {
		lrc = lrcResult.Lrc.Lyric
		lrcJson, _ := json.Marshal(lrcResult.Lrc)
		log.Infof("获取到音乐歌词 Id: %d Lrc: %s", id, lrcJson)
	}

	music := result.Data[0]
	dir := os.TempDir()
	fileName := fmt.Sprintf("/music/%d.mp3", id)

	log.Infof("开始下载音乐 Id: %d Path: %s%s", id, dir, fileName)
	// 下载音乐文件
	client := req.C().SetOutputDirectory(dir)
	_, err = client.R().SetOutputFile(fileName).Get(music.Url)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Infof("音乐下载完毕 Id: %d", id)
	return &SongDetail{
		Id:        strconv.Itoa(id),
		Url:       music.Url,
		LocalPath: fmt.Sprintf("%s%s", dir, fileName),
		Lrc:       lrc,
	}, nil
}
