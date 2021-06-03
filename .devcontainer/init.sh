#!/bin/bash

echo "Initilizing project for the first time"

echo "Tidying go module (go mod tidy)"
go mod tidy

echo "Installing dependencies (npm install)"
cd interface/web/pwa
npm install

echo "Creating empty dist directory, if it does not exist"
mkdir -p ./dist
touch ./dist/ignore-empty

echo "Project Ready!, for more details refer to the README.md file"