import { IPlaySong, IWebsocketDanmuCommand, IWebsocketMessage, IWebsocketNewSong } from "@/interfaces";
import { last } from "lodash";
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
      onDanmuCommand: null as null | ((command: IWebsocketDanmuCommand) => void),
    };
  },
  getters: {
    currentSong(): IPlaySong | undefined {
      return this.musicList[0];
    },
  },
  actions: {
    connectWebsocket() {
      this.ws = new WebSocket(`ws://${import.meta.env.DEV ? "localhost:18000" : window.location.host}/ws/connect`);
      this.ws.onmessage = (event) => {
        console.log(event);
        if (event.data) {
          const data = JSON.parse(event.data) as IWebsocketMessage<any>;
          if (data.type === "new_song") {
            const song = data.data as IWebsocketNewSong;
            const lastSong = last(this.musicList);
            if (lastSong && lastSong.id == parseInt(song.id)) {
              return;
            }

            this.musicList.push({
              id: parseInt(song.id),
              name: song.name,
              album: song.albumName,
              coverImg: song.albumPicUrl,
              duration: song.duration,
              lyric: song.lrc,
              artists: [{ id: 0, name: song.singerName }],
              songUrl: `${import.meta.env.DEV ? "http://localhost:18000" : ""}/music/${song.fileName}`,
            });
          } else if (data.type == "danmu_command") {
            this.onDanmuCommand?.(data.data);
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
