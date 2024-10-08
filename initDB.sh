#!/bin/sh

if [ "$( docker exec -u postgres postgres psql -XtAc "SELECT 1 from pg_catalog.pg_tables WHERE tablename='games'" postgres postgres)" = '1' ]; then
  echo "\n Already migrated"
  exit 1
fi

docker exec -u postgres postgres psql -c "$(cat \
  migrations/begin.sql \
  migrations/links.sql \
  migrations/commit.sql \
)" postgres postgres