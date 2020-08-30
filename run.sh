#!/bin/bash

go build -o ./bin/beppin-server && ./bin/beppin-server --read-env-vars=false
