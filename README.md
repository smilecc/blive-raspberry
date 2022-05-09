# B站直播点歌台

`blive-raspberry`是一个为B站直播编写的点歌台，可运行在云主机或VPS上。

2.0版本更换了更稳定的网易云API，新增了Web UI，废弃了原有使用ffmpeg编码推流方式，改用为obs推流。


![](https://i.imgur.com/XazH42m.jpg)

本项目灵感来自[晨旭的点歌台](https://github.com/chenxuuu/24h-raspberry-live-on-bilibili)，1.0版本功能已经失效，旧代码见`master`分支。

## 原由

原本我计划用Go复刻1.0的版本，让功能更稳定，但经过一番尝试，1.0项目推流使用的参数是将图片和音乐成每秒的3帧的视频，但是随着B站功能更新，B站的网页版直播拉流策略为每5秒钟最低100帧触发拉流动作，否则就会无限转圈圈。

以树莓派羸弱的性能（我只有树莓派3b），每秒3、4帧已经是极限，根本达不到每秒20帧的视频压制，所以2.0换了思路，主要给云主机和VPS适配，放弃了对树莓派的适配。

所以新版本的直播功能不再通过ffmpeg压制，而是转为使用docker在主机上运行一个obs studio来进行推流，UI方面使用了Web来绘制，歌曲和数据处理使用Golang来运算。

目前2.0在我的云主机（腾讯云 2核4G）上运行良好，能跑满24/s帧，程序说来说目前还有一些不稳定，后续还要近一步开发。

## 使用

### 安装与运行

release中提供了linux和windows的amd64版本的可执行文件，下载后直接打开即可使用。

程序默认端口18000，打开应用后在浏览器中访问`http://localhost:18000`即可看到界面。

### 配置

新版本提供UI来配置应用。

![](https://i.imgur.com/GyUmdtW.jpg)

### 推流

目前您可以使用项目目录下的Dockerfile构建Docker镜像，稍后会提供官方的Docker镜像。

容器运行起来以后，会看到一个VNC链接，默认端口5901，使用该链接即可操作，界面如下。

![](https://i.imgur.com/ge9vUTE.jpg)

## Thanks
- [晨旭的点歌台](https://github.com/chenxuuu/24h-raspberry-live-on-bilibili)
- [FFmpeg](http://ffmpeg.org/)