#!/bin/sh

docker build -t usql . && docker run -p 8080:8080 usql