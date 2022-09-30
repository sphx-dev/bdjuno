#!/bin/sh


export PGPASSWORD=$POSTGRES_PASSWORD
export PGHOST=$POSTGRES_HOST
export PGUSER=$POSTGRES_USER
export PGDATABASE=$POSTGRES_DATABASE

psql -h $POSTGRES_HOST -U $POSTGRES_USER -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$POSTGRES_DATABASE'" | grep -q 1 || psql -h $POSTGRES_HOST -U $POSTGRES_USER -d postgres -c "CREATE DATABASE $POSTGRES_DATABASE"
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/00-cosmos.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/01-auth.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/02-bank.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/03-staking.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/04-consensus.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/05-mint.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/06-distribution.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/07-pricefeed.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/09-modules.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/10-slashing.sql
echo $POSTGRES_PASSWORD | psql -h $POSTGRES_HOST -W -U $POSTGRES_USER -d $POSTGRES_DATABASE -f /var/lib/postgresql/schema/11-feegrant.sql
