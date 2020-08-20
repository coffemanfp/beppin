#!/bin/bash

cd migrations
go build -o ../bin/migrations
cd ../

echo 'Starting migrations...'
./bin/migrations --with-examples
