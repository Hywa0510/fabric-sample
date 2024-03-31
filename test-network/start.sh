#!/bin/bash

./stop.sh
# 启动区块链网络、创建通道
./network.sh up createChannel
# 部署链码，使用basic链码
./network.sh deployCC -ccn creditfile -ccp ../chaincode -ccl go

# 启动区块链浏览器
cp -r organizations explorer/
docker compose -f explorer/docker-compose.yaml up -d