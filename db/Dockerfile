# Use the official PostgreSQL image
FROM postgres:14.5

# Set environment variables
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=nkyVdeMfJGTayFRDVCaacSjefRtURy
ENV POSTGRES_DB=books_db

# Copy the initialization script
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
