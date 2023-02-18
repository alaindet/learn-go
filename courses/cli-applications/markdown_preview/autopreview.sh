#! /bin/bash

# Run mdp tool every 5 seconds if the file changes
# Usage
# ./autopreview.sh ./data/README.md

FILEHASH=`md5sum $1`
while true; do
  NEWFILEHASH=`md5sum $1`
  if [ "$NEWFILEHASH" != "$FILEHASH" ]; then
    ./mdp -file $1
    FILEHASH=$NEWFILEHASH
  fi
  sleep 5
done
