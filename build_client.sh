# use docker container to build client
docker run -it -d --name node1 node:8.9.4-alpine;

docker exec -it node /bin/sh -c "mkdir /home/client";
docker cp ./client node1:/home/client;

docker exec -it node1 /bin/sh -c "cd /home/client && yarn install";
docker exec -it node1 /bin/sh -c "cd /home/client && yarn run build";
docker cp node1:/home/dist ./dist;

docker kill node1;
docker rm node1;

