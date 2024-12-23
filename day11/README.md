# Day 11 - Docker Compose Introduction

Hey there! Congrats to making it to day 11. Today we will learn about Docker Compose. Over the last few days, we’ve learned about environment variables, networking, and volumes. Let’s look at a typical command we might use to run a container with all these features:

```bash
docker run -d \
  --name api \
  --network myapp-network \
  -v mydata:/data \
  -e SECRET_MESSAGE="Hello Docker!" \
  -e PORT=8080 \
  -p 8080:8080 \
  hello-world-go

```

That’s quite a lot to remember and type! Now imagine running multiple containers, each with their own configuration. This is where Docker Compose comes in.

## What is Docker Compose?

Docker Compose is a tool that helps you define and run multi-container applications. Instead of long docker commands, you write a YAML file that describes your entire application stack.

Let’s convert our previous setup to a Docker Compose file. Create a new file called `docker-compose.yml`:

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - SECRET_MESSAGE=Hello Docker!
      - PORT=8080
    volumes:
      - mydata:/data
    networks:
      - myapp-network
volumes:
  mydata:
networks:
  myapp-network:
```

Now instead of that long docker command, you can simply run:

```bash
docker compose up

```

Docker Compose will:

1. Create the network
2. Create the volume
3. Build the image (if needed)
4. Start the container with all specified configuration

To stop everything:

```bash
docker compose down

```

Let’s add some more complexity to our application. We will add a database and a reverse proxy.

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - SECRET_MESSAGE=Hello Docker!
      - PORT=8080
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
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=myapp
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

What we did here is:

- Added a database service
- Added a reverse proxy service
- Added a dependency to the api service (so it starts after the database)
- Added a dependency to the nginx service (so it starts after the api)

For simplicity sake we hardcoded our secrets here. In a production environment you would of course not want to do that! How you can properly manage secrets is a topic for another day (probably tomorrow).

## Basic Docker Compose Commands

Here are the most important commands you’ll use:

- `docker compose up` - Start services
- `docker compose up -d` - Start in detached mode
- `docker compose down` - Stop and remove containers
- `docker compose ps` - List running services
- `docker compose logs` - View service logs
- `docker compose build` - Build or rebuild services
