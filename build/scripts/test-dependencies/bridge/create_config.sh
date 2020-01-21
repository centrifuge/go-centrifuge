#!/usr/bin/env bash

config_dir=$1
asset_address=$2

eth_config="[[chains]]
id = 0
endpoint = \"ws://geth:9546\"
emitter = \"0x1fA38b0EfccA4228EB9e15112D4d98B0CEe3c600\"
receiver = \"$asset_address\"
from = \"021a7faa177621e1a6bb2e2794d77bcd1040221bcf5f0b83180f368eaab7848551\""

cent_config='[[chains]]
id = 1
endpoint = "ws://cc:9944"
emitter = "0x1fA38b0EfccA4228EB9e15112D4d98B0CEe3c600"
receiver = "0x290f41e61374c715C1127974bf08a3993afd0145"
from = "021a7faa177621e1a6bb2e2794d77bcd1040221bcf5f0b83180f368eaab7848551"'

echo "$eth_config" > "$config_dir"/config.toml
echo "" >> "$config_dir"/config.toml
echo -n "$cent_config" >> "$config_dir"/config.toml
