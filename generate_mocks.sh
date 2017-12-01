#!/bin/bash
set -euo pipefail

mockgen -package websocketwrapper -source websocketwrapper.go > mockwebsocketwrapper.go

