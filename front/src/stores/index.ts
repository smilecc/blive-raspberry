import { IPlaySong, IWebsocketMessage, IWebsocketNewSong } from "@/interfaces";
import { defineStore } from "pinia";

export const useCommonStore = defineStore("common", {});

export const usePlayerStore = defineStore("player", {
  state() {
    return {
      playerRef: undefined as HTMLAudioElement | undefined,
      ws: null as WebSocket | null,
      isPause: false,
      currentTime: 0,
      musicList: [] as IPlaySong[],
    };
  },
  getters: {
    currentSong(): IPlaySong | undefined {
      return this.musicList[0];
    },
  },
  actions: {
    connectWebsocket() {
      this.ws = new WebSocket("ws://localhost:18000/ws/connect");
      this.ws.onmessage = (event) => {
        console.log(event);
        if (event.data) {
          const data = JSON.parse(event.data) as IWebsocketMessage<IWebsocketNewSong>;
          if (data.type === "new_song") {
            const song = data.data;
            this.musicList.push({
              id: parseInt(song.id),
              name: song.name,
              album: song.albumName,
              coverImg: song.albumPicUrl,
              duration: song.duration,
              lyric: song.lrc,
              artists: [{ id: 0, name: song.singerName }],
              songUrl: `http://localhost:18000/music/${song.fileName}`,
            });
          }
        }
      };
    },
    closeWebsocket() {
      if (this.ws) {
        this.ws?.close();
        this.ws = null;
      }
    },
  },
});
