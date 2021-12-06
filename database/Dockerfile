FROM postgres:latest
# create_fixtures.sql creates database and users table, runs after postgres initialization
COPY create_fixtures.sql /docker-entrypoint-initdb.d
