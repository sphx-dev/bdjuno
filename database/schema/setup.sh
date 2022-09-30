#!/bin/sh


export PGPASSWORD=$POSTGRES_PASSWORD
export PGHOST=$POSTGRES_HOST
export PGUSER=$POSTGRES_USER
export PGDATABASE=$POSTGRES_DATABASE

psql -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$POSTGRES_DATABASE'" | grep -q 1 || psql -d postgres -c "CREATE DATABASE $POSTGRES_DATABASE"
psql -f /var/lib/postgresql/schema/00-cosmos.sql
psql -f /var/lib/postgresql/schema/01-auth.sql
psql -f /var/lib/postgresql/schema/02-bank.sql
psql -f /var/lib/postgresql/schema/03-staking.sql
psql -f /var/lib/postgresql/schema/04-consensus.sql
psql -f /var/lib/postgresql/schema/05-mint.sql
psql -f /var/lib/postgresql/schema/06-distribution.sql
psql -f /var/lib/postgresql/schema/07-pricefeed.sql
psql -f /var/lib/postgresql/schema/09-modules.sql
psql -f /var/lib/postgresql/schema/10-slashing.sql
psql -f /var/lib/postgresql/schema/11-feegrant.sql
