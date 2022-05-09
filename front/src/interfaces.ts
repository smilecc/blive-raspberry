export interface ILyric {
  lyric: string /** 歌词（包含歌词和每句歌词的时间） */;
  version: number;
}

export interface IFormatLyric {
  time: number /** 歌词播放时间，格式：00:00.123, 分:秒:毫秒 */;
  text: string /** 歌词内容（可能为空字符串） */;
  transText?: string /** 翻译歌词内容（可能为空字符串） */;
}

export interface IlyricUser {
  id: number /** 记录id */;
  nickname: string /** 歌词贡献者名 */;
  uptime: number /** 歌词上传时间（时间戳） */;
  userid: number /** 歌词贡献者id */;
}

/** 歌词接口定义 */
export interface ILyricResponse {
  lrc: ILyric /** 歌词 */;
  lyricUser: IlyricUser /** 歌词贡献者 */;
  tlyric: ILyric /** 翻译后的歌词 */;
  transUser: IlyricUser /** 翻译歌词贡献者 */;
}

export interface ILyricState {
  lyric: string /** 歌词 */;
  lyricUser?: IlyricUser /** 歌词贡献者  */;
  transLyric?: string /** 翻译歌词 */;
  transLyricUser?: IlyricUser /** 翻译歌词贡献者 */;
}

export interface IPlaySong {
  id: number /** 歌曲id */;
  coverImg: string /** 歌曲封面图片 */;
  name: string /** 歌曲名 */;
  subName?: string /** 歌曲副标题 */;
  album?: string /** 歌曲专辑 */;
  artists: IArtist[] /** 歌手 */;
  duration?: number /** 歌曲时长（秒：s） */;
  songUrl: string /** 歌曲的播放地址 */;
  hasCollected?: boolean /** 歌曲是否已收藏 */;
  lyric?: string /** 歌词 */;
}

export interface IArtist {
  id: number;
  name: string /** 歌手名 */;
  picUrl?: string /** 歌手封面图片 */;
  cover?: string /** 歌手封面 */;
  musicSize?: number /** 单曲数 */;
  mvSize?: number /** 专辑数 */;
  videoCount?: number /** MV数 */;
  alias?: string[] /** 歌手别名（英文名｜中文名） */;
}

export interface IWebsocketMessage<T = any> {
  type: string;
  data: T;
}

export interface IWebsocketNewSong {
  id: string;
  name: string;
  fileName: string;
  url: string;
  localPath: string;
  lrc: string;
  singerName: string;
  duration: number;
  albumName: string;
  albumPicUrl: string;
}
