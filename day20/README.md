# Day 20 - Docker Security Best Practices

Welcome to Day 20 of Advent of Docker! Today we’ll cover essential security practices for Docker. Security is crucial - a misconfigured container can compromise your entire host system. These are some of the most important rules in my opinion, but the list is definitely not complete. Security is a never-ending journey :)

## 1. Keep Host & Docker Updated

The host kernel is shared between containers, so vulnerabilities like Dirty COW or Leaky Vessels can affect all containers. Always:

- Update the host system regularly
- Keep Docker Engine updated
- Enable automatic security updates where possible

## 2. Protect the Docker Daemon Socket

The Docker socket (`/var/run/docker.sock`) is a critical security point. Access to it means root access to the host.

Don’t do this:

```yaml
volumes:
  - "/var/run/docker.sock:/var/run/docker.sock" # Dangerous!
```

Never:

- Expose the Docker daemon socket to containers
- Enable TCP socket without TLS
- Run Docker with -H tcp://0.0.0.0:2375

## 3. Use Official Images

Base images are your security foundation:

### _Good_

```dockerfile
FROM ubuntu:22.04

```

### _Better_

```dockerfile
FROM ubuntu:22.04-slim

```

### _Best for production_

```dockerfile
FROM scratch # For compiled languages
FROM distroless # For interpreted languages

```

## 4. Run as Non-Root

Create a dedicated user in your Dockerfile:

```dockerfile
RUN groupadd -r app && useradd -r -g app -d /home/app -s /sbin/nologin -c "Docker image user" app
USER app

```

## 5. Use Multi-Stage Builds

Reduce attack surface by keeping only what’s needed! Check out Day 13 for more details!

## 6. Limit Resources

Limit the impact of a DoS attack by setting resource limits, like memory and CPU. Check out Day 19 for more details!

## 7. Use Vulnerability Scanners

Vulnerability scanners help you catch security issues in your images. There are quite a few available, but Docker Desktop ships one called Docker Scout by default:

![Docker Scout](https://adventofdocker.com/_astro/docker-scout.DhYLLqKJ_aPOEr.webp)

One important thing to note is that vulnerability scanners will always have false-positives. They will find issues that are not actually a problem. Always verify the results and check if the issue is actually a problem for your use case!

## Conclusion

As said in the beginning, this list is not complete and there are many more things to consider. Security is a never-ending journey and you should always be learning.

Here are some resources that I found helpful:

- [Docker Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Docker_Security_Cheat_Sheet.html)
- [Docker Security](https://spacelift.io/blog/docker-security)
- [Docker Engine Security](https://docs.docker.com/engine/security/)
