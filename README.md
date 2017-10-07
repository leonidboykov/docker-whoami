# Who Am I

[![Build Status](https://travis-ci.org/leonidboykov/whoami.svg?branch=master)](https://travis-ci.org/leonidboykov/whoami)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonidboykov/whoami)](https://goreportcard.com/report/github.com/leonidboykov/whoami)
[![Layers](https://images.microbadger.com/badges/image/leonidboykov/whoami.svg)](https://microbadger.com/images/leonidboykov/whoami "Get your own image badge on microbadger.com")

<!-- see https://github.com/leonidboykov/whoami for better Markdown rendering -->

Super slim _who am i_ image based on `scratch`. No overweight base images or
additional fuctionality.

### Run container

    sudo docker pull leonidboykov/whoami
    sudo docker run -d --name whoami -p 8080:8080 leonidboykov/whoami

Check hostname

    curl localhost:8080
