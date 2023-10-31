#!/bin/bash

docker image build -f Dockerfile -t ascii .

docker container run -p 8080:8080 --detach --name my_container ascii
