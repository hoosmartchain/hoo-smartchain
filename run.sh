#!/usr/bin/env bash
/data/hsc/geth-linux-amd64 \
--config /data/hsc/config.toml  \
--logpath /data/hsc/logs \
--syncmode full \
--gcmode archive \
--verbosity 3  >> /data/hsc/logs/systemd_chain_console.out 2>&1
