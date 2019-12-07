#!/usr/bin/env bash
BUILD_DIR=$(cd $(dirname $0);pwd)
ROOT_DIR=${BUILD_DIR}/..
BIN_DIR=${ROOT_DIR}/bin

if [[ ! -d ${BIN_DIR} ]];then
    mkdir "${BIN_DIR}"
fi

set_env_arm() {
export GOOS=linux
export GOARCH=arm
# TODO: we should try to turn on vfp.see https://github.com/golang/go/issues/18483,
export GOARM=5
}

set_env_arm

#GOARCH=arm GOARM=7 GOOS=linux go build -ldflags="-s -w" -o cloud_cron .
go build -o cloud_cron .

mv cloud_cron "${BIN_DIR}"

