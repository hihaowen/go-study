#!/bin/bash

GOMAXPROCS=2 go test -timeout 30s $(go list ./...)

