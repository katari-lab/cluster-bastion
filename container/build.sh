docker-compose build

docker tag katari/bastion localhost:32000/katari/bastion
docker push localhost:32000/katari/bastion
