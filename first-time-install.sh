#!/bin/bash

set -xeuf -o pipefail

if [[ -v CI ]]; then
    SUDO=""
    apt-get update
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata
    echo "en_US UTF-8" > /etc/locale.gen
else
    SUDO="sudo"
fi

# Install Go
${SUDO} apt update
${SUDO} apt-get install -y curl tar
${SUDO} apt-get update

${SUDO} rm -rf /usr/local/go
wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
${SUDO} tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
echo 'export GOPATH=$HOME/go' >> /etc/profile
echo 'export PATH=$PATH:$GOPATH/bin' >> /etc/profile
source /etc/profile
rm go1.20.2.linux-amd64.tar.gz

${SUDO} apt-get install -y libpq-dev

(cd app/src && go get github.com/lib/pq && go get github.com/golang-migrate/migrate/v4)

wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
rm migrate.linux-amd64.deb

sudo apt-get install -y postgresql-client

# Install Docker
if ! type "docker" > /dev/null; then
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | ${SUDO} gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | ${SUDO} tee /etc/apt/sources.list.d/docker.list > /dev/null

    ${SUDO} apt update

    apt-cache policy docker-ce

    if [[ -v CI ]]; then
        echo "Skipping creation of docker group"
    else
        ${SUDO} apt install -y docker-ce
        ${SUDO} usermod -aG docker ${USER}
    fi
else
    echo 'docker already installed'
fi

# Install Docker-Compose
if ! type "docker-compose" > /dev/null; then
    ${SUDO} curl -SL https://github.com/docker/compose/releases/download/v2.3.3/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
    ${SUDO} chmod +x /usr/local/bin/docker-compose
else
    echo 'docker-compose already installed'
fi