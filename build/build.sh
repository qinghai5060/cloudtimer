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
export GOARM=7
}

set_env_arm

#GOARCH=arm GOARM=7 GOOS=linux go build -ldflags="-s -w" -o cloud_cron .
GOARCH=arm GOARM=7 GOOS=linux go build -o cloud_cron .

mv cloud_cron "${BIN_DIR}"

