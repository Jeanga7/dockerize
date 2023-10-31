#!/bin/bash
image_name="ascii"

port_local="8080"

docker build -t $image_name .

docker run -p $port_local:8080 $image_name
