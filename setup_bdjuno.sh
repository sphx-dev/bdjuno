#!/bin/bash

# Variables (set these according to your environment)
bdjuno_dir="/home/galen/code/bdjuno"
bigdipper_db="bdjuno"
bigdipper_db_user="galen"
bigdipper_db_password="Warhawks58!"
bigdipper_user_home="/home/galen"
bigdipper_user="galen" # bigdipper_user_galen
chain_url="http://localhost:26557"
chain_version="latest"
bdjuno_bin="bdjuno" ## /home/galen/code/bdjuno/bin/bdjuno-builder"
bdjuno_home="/home/galen/code/bdjuno/.bdjuno"

# Find bdjuno schema files
bdjuno_schemas=$(find $bdjuno_dir/database/schema -type f)
echo "Schema files found:"
echo "$bdjuno_schemas" | sort

# Run postgresql bdjuno setup scripts
if [ "${remove_db:-false}" = "true" ]; then
  for schema in $bdjuno_schemas; do
    PGPASSWORD=$bigdipper_db_password psql -h "127.0.0.1" -U $bigdipper_db_user -d $bigdipper_db -f "$schema"
  done
fi

echo "done postgres step"

# Create bdjuno_bin directory
mkdir -p $bigdipper_user_home/go/bin
chown $bigdipper_user:$bigdipper_user $bigdipper_user_home/go/bin
chmod 0700 $bigdipper_user_home/go/bin

echo "A"
# Init BDJuno
#bash -c "$bdjuno_bin init --home $bdjuno_home"
echo "B"
# Install bdjuno config.yaml
#cp bdjuno-config.yaml.j2 $bdjuno_home/config.yaml
ls $bdjuno_home/config.yaml
chown $bigdipper_user:$bigdipper_user $bdjuno_home/config.yaml
chmod 0600 $bdjuno_home/config.yaml

# Install Genesis
#cp genesis.j2 $bdjuno_home/genesis.json
ls $bdjuno_home/genesis.json
chown $bigdipper_user:$bigdipper_user $bdjuno_home/genesis.json
chmod 0600 $bdjuno_home/genesis.json

# Parse genesis file with bdjuno
bash -c "$bdjuno_bin parse genesis-file --home $bdjuno_home --genesis-file-path $bdjuno_home/genesis.json"
