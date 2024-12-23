### Container Communication

Create the network

```bash
docker network create myapp-network
```

Run the containers

```bash
docker run -d -p 3000:8080 --name api --network myapp-network hello-world-go
```

```bash
docker run -d -p 80:80 --name frontend --network myapp-network nginx
```

### Custom Hostnames

```bash
docker run -p 3000:8080 --name api --hostname custom-api --network myapp-network hello-world-go
```

### Verify hostname inside the container

```bash
docker exec -it api hostname
```
