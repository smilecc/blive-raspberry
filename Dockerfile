FROM bandi13/gui-docker:1.2
USER root
COPY ./docker/sources.list /etc/apt/sources.list
RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update -y \
    && apt-get install -y software-properties-common \
    && add-apt-repository ppa:obsproject/obs-studio \
    && apt-get update -y \
    && apt-get install -y obs-studio wget curl \
    && apt-get install -y language-pack-ja language-pack-zh* language-pack-ko \
    && apt-get clean -y

RUN echo "?package(bash):needs=\"X11\" section=\"DockerCustom\" title=\"OBS Screencast\" command=\"obs\"" >> /usr/share/menu/custom-docker && update-menus

WORKDIR /root/blive
RUN wget https://github.com/smilecc/blive-raspberry/releases/download/v2.0.1/blive_linux_amd64
RUN mv blive_linux_amd64 blive


