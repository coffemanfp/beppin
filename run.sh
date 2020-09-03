#!/bin/bash

go build -o ./bin/beppin && ./bin/beppin --read-env-vars=false
