#!/bin/bash

# 启动区块链网络、创建通道
./network.sh up createChannel
# 部署链码，使用basic链码
./network.sh deployCC -ccn creditfile -ccp ../chaincode -ccl go