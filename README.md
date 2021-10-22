# kafka-golang-example
Proof of concept of kafka consumer/producer messages in golang

## Getting started

This project contains 2 docker compose files

`docker-compose.pre.yml` for spinning up kafka and zookeper official images

`docker-compose.yml`kafka consumer and producer services implmentations in golang

## Configuration seup

First run `docker-compose -f docker-compose.pre.yml up --build`

Second run `docker-compose -f docker-compose.yml up --build`
