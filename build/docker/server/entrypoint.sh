#!/bin/bash
set -o errexit

if [ ! -w /srv ]; then
    echo "FATAL: /srv is not writable by $(whoami) user." >&2
    echo "Please make sure that /srv is owned by uid $(id -u) and gid $(id -g) and try again." >&2
    echo "You can change ownership by running: sudo chown -R $(id -u):$(id -g) /srv" >&2
    exit 1
fi

# Initialize /srv if empty
INIT_FILE=/srv/initialized
if [ ! -f INIT_FILE ]; then
    echo "Initializing /srv..."
    mkdir -p /srv/{backup,clickhouse,grafana,logs,nginx,postgres14,prometheus,victoriametrics}
    echo "Copying grafana plugins and the VERSION file..."
    cp -r /usr/share/percona-dashboards/panels/* /srv/grafana/plugins
    cp /usr/share/percona-dashboards/VERSION /srv/grafana/PERCONA_DASHBOARDS_VERSION
    
    echo "Generating self-signed certificates for nginx..."
    bash /var/lib/cloud/scripts/per-boot/generate-ssl-certificate
    
    echo "Initializing Postgres..."
    /usr/pgsql-14/bin/initdb -D /srv/postgres14 --auth=trust --username=postgres
    
    echo "Enabling pg_stat_statements extension for PostgreSQL..."
    /usr/pgsql-14/bin/pg_ctl start -D /srv/postgres14 -o '-c logging_collector=off'
    /usr/bin/psql postgres postgres -c 'CREATE EXTENSION pg_stat_statements SCHEMA public'
    /usr/pgsql-14/bin/pg_ctl stop -D /srv/postgres14
    touch $INIT_FILE
fi
DIST_FILE=/srv/pmm-distribution
if [ ! -f $DIST_FILE ]; then
    echo "docker" > $DIST_FILE
fi

# pmm-managed-init validates environment variables.
pmm-managed-init

# Start supervisor in foreground
exec supervisord -n -c /etc/supervisord.conf
