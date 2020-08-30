#!/bin/bash

cd migrations
go build -o ../bin/migrations
cd ../

echo 'Starting migrations...'
./bin/migrations --read-env-vars=false