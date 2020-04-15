docker volume create servervol
docker run --network="host" --mount source=servervol,target=/serverdata arbaazkhan/server:v1
