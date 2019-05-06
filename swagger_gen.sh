#!/usr/bin/env bash
swagger generate spec -b . -o ./swagger/swagger.json -m

mv ./swagger/swagger.json ~/