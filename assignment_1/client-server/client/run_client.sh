docker volume create clientvol
docker run --network="host" --mount source=clientvol,target=/clientdata arbaazkhan/client:v1
