# Day 10 - Docker Networking Basics

On Day 5 you learned that Docker containers have some network isolation. Today, we’re going to dive deeper into Docker networking and learn how containers can communicate with each other!

## Publishing Ports Refresher

Remember our HTTP server from Day 5? By default, containers are isolated and we needed to publish ports to access them from our host:

```bash
$ docker run -p 8080:8080 hello-world-go

```

The -p flag maps port 8080 from the container to port 8080 on your host. You can also map to a different host port:

```bash
$ docker run -p 3000:8080 hello-world-go

```

### _Publishing All Ports_

If your container exposes multiple ports (using the EXPOSE instruction), you can publish all of them at once with -P:

```bash
$ docker run -P hello-world-go

```

Docker will automatically assign random host ports. Use docker ps to see the mappings:

```bash
$ docker ps
CONTAINER ID   IMAGE           PORTS                     NAMES
abc123def456   hello-world-go  0.0.0.0:55000->8080/tcp  cool_name

```

## Container Communication

What if containers need to talk to each other? Let’s create a simple network:

```bash
$ docker network create myapp-network

```

Now run two containers on this network:

```bash
$ docker run -d --name api --network myapp-network hello-world-go
$ docker run -d --name frontend --network myapp-network nginx

```

Containers on the same network can reach each other using their container names as hostnames! From inside the frontend container:

```bash
$ docker exec -it frontend curl http://api:8080
Hello World!

```

## Network Types

Docker supports several network types:

- bridge: The default network driver. Good for containers on a single host (what we used above)
- host: Removes network isolation, container uses host’s network directly
- none: Disables networking completely
- overlay: For connecting containers across multiple Docker hosts
- macvlan: Assigns a MAC address to containers, making them appear as physical devices

You can read more about network drivers in the official documentation.

## Custom Hostnames

While Docker uses container names as hostnames by default, you can set custom ones:

```bash
$ docker run --name api --hostname custom-api --network myapp hello-world-go

```

Now other containers on the network can reach this container using either api or custom-api as the hostname.

You can verify the hostname from inside the container:

```bash
$ docker exec -it api hostname
custom-api

```
