#!/bin/bash

if [[ -d /tmp/entrypoint.d ]]; then
  for f in /tmp/entrypoint.d/*.sh ;do
    source $f
  done
fi

/root/vltreplicator -config /etc/vltreplicator.yml
