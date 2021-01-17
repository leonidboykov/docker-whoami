# Who Am I

[![Build Status](https://travis-ci.org/leonidboykov/docker-whoami.svg?branch=master)](https://travis-ci.org/leonidboykov/docker-whoami)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonidboykov/docker-whoami)](https://goreportcard.com/report/github.com/leonidboykov/docker-whoami)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/leonidboykov/whoami)](https://hub.docker.com/repository/docker/leonidboykov/whoami)
[![Layers](https://images.microbadger.com/badges/image/leonidboykov/whoami.svg)](https://microbadger.com/images/leonidboykov/whoami "Get your own image badge on microbadger.com")

<!-- see https://github.com/leonidboykov/docker-whoami for better markdown rendering -->

Super slim _who am i_ image based on `scratch`. No overweight base images or additional fuctionality.

### Run container

    sudo docker pull leonidboykov/whoami
    sudo docker run -d --name whoami -p 8080:8080 leonidboykov/whoami

Check hostname

    curl localhost:8080
