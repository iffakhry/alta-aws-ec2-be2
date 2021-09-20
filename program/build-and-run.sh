#!/bin/bash
go build -o altastore
sudo setcap 'cap_net_bind_service=+ep' "$(realpath altastore)"
source prod.env
./altastore