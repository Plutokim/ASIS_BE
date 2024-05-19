#!/bin/bash

docker-compose up -d

docker-compose exec sport-backend ./air.sh
