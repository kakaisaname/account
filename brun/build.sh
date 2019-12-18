#!/usr/bin/env bash

#cpath=`pwd`
#PROJECT_PATH=${cpath%src*}
#echo $PROJECT_PATH
#export GOPATH=$GOPATH:${PROJECT_PATH}

SOURCE_FILE_NAME=main
TARGET_FILE_NAME=reskd

rm -fr ${TARGET_FILE_NAME}*

build(){
    echo $GOOS $GOARCH                                      #打印系统和系统架构
    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}       #目标文件名称命名
    #编译
    env GOOS=$GOOS GOARCH=$GOARCH \                         #设置系统和架构
    go build -o ${tname} \
    -v ${SOURCE_FILE_NAME}.go
    #添加可执行权限
    chmod +x ${tname}
    mv ${tname} ${TARGET_FILE_NAME}${EXT}
    if [ ${GOOS} == "windows" ];then
        zip ${tname}.zip ${TARGET_FILE_NAME}${EXT} config.ini ../public/
    else
        tar --exclude=*.gz  --exclude=*.zip  --exclude=*.git -czvf ${tname}.tar.gz ${TARGET_FILE_NAME}${EXT} config.ini *.sh ../public/ -C ./ .
    fi
    mv ${TARGET_FILE_NAME}${EXT} ${tname}

}
CGO_ENABLED=0
#mac os 64
GOOS=darwin
GOARCH=amd64
build

#linux 64
GOOS=linux
GOARCH=amd64
build

#windows
#64
GOOS=windows
GOARCH=amd64
build

GOARCH=386
build