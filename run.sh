#!/bin/bash
WORLD_DATA_PATH=`pwd` \
WORLD_SAVE_TARGET=filesystem \
WORLD_SAVE_DIRECTORY=`pwd`/dist \
WORLD_WEB_DOMAIN=ironarachne.test \
dist/world $1
