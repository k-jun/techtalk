#!/bin/bash

# ----------BASIC SETUP----------
apt -y update
apt install -y make

# ----------DOCKER----------
apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88

add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt -y update
apt install -y docker-ce docker-ce-cli containerd.io
# to use docker command without sudo
usermod -aG docker ubuntu
apt install -y docker-compose

# ----------DEPLOY----------

git clone http://github.com/k-jun/techtalk
cd techtalk
docker volume create --name=mysql-volume
docker-compose up -d mysql redis
sleep 20
docker-compose up -d app web

# chown -R ubuntu /home/ubuntu/
# chgrp -R ubuntu /home/ubuntu/
