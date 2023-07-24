FROM postgres

# Set environment variables
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB mydatabase

# Copy the SQL scripts to the docker-entrypoint-initdb.d directory
COPY init-db.sh /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
