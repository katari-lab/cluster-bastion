docker-compose build

docker tag katari/bastion localhost:32000/katarilab/bastion
docker push localhost:32000/katarilab/bastion
