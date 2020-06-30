#!/bin/bash
cd migrations

echo 'Starting migrations...'
go run migrations.go -with-examples