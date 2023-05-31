#!/bin/sh

docker buildx build --push --platform linux/arm64 -t dinhlockt02/cs_chat_app_server:dev-bulleye-slim .
docker buildx build --push --platform linux/amd64 -t dinhlockt02/cs_chat_app_server:dev-bulleye-slim-amd .
