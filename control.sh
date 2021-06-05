#! /bin/bash

set -e

CONF_NAME="service.conf"
SERVICE_NAME="goldtalkAPI"

CUR_DIR=$(dirname "$0")
echo "CUR_DIR: $CUR_DIR"
MY_PATH=$(cd "$CUR_DIR" && pwd -P)
echo "MY_PATH: $MY_PATH"
BASE_CONF_NAME="$MY_PATH/conf/$CONF_NAME"


function setConfigFile() {
    echo "pwd: $MY_PATH"
    echo "NODE_ENV: $NODE_ENV"
    if [ "$NODE_ENV" ]; then
        echo "change to conf"
        rm -f "$BASE_CONF_NAME"
        ln -s "$BASE_CONF_NAME.$NODE_ENV" "$BASE_CONF_NAME"
    fi
}

function start() {
	setConfigFile
	echo "bin$MY_PATH/bin/$SERVICE_NAME"
	echo "$MY_PATH/conf/$CONF_NAME"
	./bin/goldtalkAPI
}

function stop() {
    echo "todo stop"
}

function restart() {
    stop
    sleep 1
    start
}

function usage() {
    echo "Usage: $0 {start|stop|restart}"
    exit 1
}

if [ $# != 1 ]; then
    usage
fi

case "$1" in
    start|stop|restart)
        $1
        ;;
    *)
        usage
esac
