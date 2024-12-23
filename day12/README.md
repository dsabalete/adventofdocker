# Day 12 - Docker Compose Deep Dive

Welcome to day 12 of Advent of Docker! Yesterday we learned the basics of Docker Compose. Today, we’ll explore some more advanced features that make Docker Compose even more powerful.

## Environment Files

Remember when we hardcoded database credentials in our compose file yesterday? Let’s fix that using environment files. Create a file called `.env`:

```plaintext
POSTGRES_USER=postgres
POSTGRES_PASSWORD=my-super-secret-password
POSTGRES_DB=myapp
PORT=8080
SECRET_MESSAGE=Hello Docker!

```

Now modify the compose file to use the environment file:

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    volumes:
      - mydata:/data
    networks:
      - myapp-network
    depends_on:
      - db
  db:
    image: postgres
    ports:
      - "5432:5432"
    env_file:
      - .env
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - myapp-network
  nginx:
    image: nginx
    ports:
      - "80:80"
    networks:
      - myapp-network
    depends_on:
      - api
volumes:
  mydata:
  postgres_data:
networks:
  myapp-network:
```

You can specify different environment files per service:

```yaml
services:
  db:
    image: postgres
    env_file:
      - .env.db # Database specific env file

  api:
    build: .
    env_file:
      - .env.api # API specific env file
```

Remember to add your `.env` files to .gitignore to keep your secrets safe! You can provide an example configuration by creating a `.env.example` file:

```plaintext
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
PORT=8080
SECRET_MESSAGE=

```

## Healthchecks

Docker Compose provides a healthcheck feature to check if a service is healthy, and to wait for a service to be healthy before starting another service if it depends on it.

```yaml
services:
  api:
    build: .
    depends_on:
      db:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
  db:
    image: postgres
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5
```

Here we’ve added:

- `healthcheck`: Defines how Docker should check if a service is healthy
- `depends_on` with condition: Waits for the database to be healthy before starting the API

This ensures our API only starts when the database is ready and your application is healthy.

## Compose Watch

Docker Compose Watch is a fantastic feature for development. It automatically syncs your code changes into containers and can trigger rebuilds:

```yaml
services:
  api:
    build: .
    ports:
      - "8081:8081"
    environment:
      - PORT=8081
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
    develop:
      watch:
        - path: ./
          target: /go
          action: rebuild
```

With this configuration:

1. Any changes to the `./` directory are synced instantly into the container and trigger a rebuild

Start your development environment with:

```bash
docker compose watch

```

## Profiles

Profiles help you manage different service combinations. For example, you might want development tools only in development:

```yml
services:
  api:
    build: .
  db:
    image: postgres
  adminer:
    image: adminer
    profiles:
      - dev
    ports:
      - "9000:8080"
```

Now you can:

- Start normal services: docker compose up
- Include dev tools: docker compose --profile dev up

## Best Practices

1. Use health checks for critical services
2. Never store secrets in your compose file
   - Use secret files
   - Or use environment variables
3. Use profiles to separate development and production services
4. Leverage watch for faster development
5. Document dependencies using `depends_on`
