#!/bin/bash
WORKDIR=`pwd`
FILEDIR="${WORKDIR}/build"
WORLDAPI_DATA_PATH=$WORKDIR \
WORLDAPI_SAVE_TARGET=filesystem \
WORLDAPI_SAVE_DIRECTORY=$FILEDIR \
WORLDAPI_WEB_DOMAIN=ironarachne.test \
build/worldapi