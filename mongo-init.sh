#!/bin/bash
set -e

mongosh <<EOF
use $MONGO_DATABASE
db.createUser(
  {
    user: "$MONGO_USERNAME",
    pwd: "$MONGO_PASSWORD",
    roles: [ "readWrite" ]
  }
)
EOF