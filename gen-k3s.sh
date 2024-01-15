#!/bin/bash

mkdir k3s
cd k3s
kompose convert --file ../docker-compose.yml --volumes hostPath
kompose convert --file ../docker-compose.yml --volumes hostPath --out all
