# 关闭区块链浏览器
docker compose -f explorer/docker-compose.yaml down -v 
# 关闭区块链网络
./network.sh down
# 删除organizations
rm -rf explorer/organizations