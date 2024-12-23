# Day 16 - Introduction to DockerHub

Hey there! Welcome to Day 16 of Advent of Docker. Today we’re going to learn about container registries and specifically focus on DockerHub - the default and most popular container registry.

## What is a Container Registry?

A container registry is like a GitHub for Docker images. It’s a place where you can:

- Store your Docker images
- Share them with others
- Version control your containers
- Collaborate with your team

Just like you don’t keep all your code on your local machine, you shouldn’t keep all your Docker images locally either! While there are many container registries available, DockerHub is the most widely used and is the default registry for Docker.

## Using DockerHub

DockerHub is where all those images we’ve been pulling come from! Remember all those docker pull commands we’ve been using? They were all pulling from DockerHub.

Let’s learn how to push our own image to DockerHub:

1. First, create a DockerHub account if you haven’t already

2. Login to DockerHub from your terminal:

```bash
$ docker login

```

3. Tag your image with your DockerHub username:

```bash
$ docker tag hello-world-go YOUR_USERNAME/hello-world-go

```

At this point you should make sure that your Docker image is not including any sensitive information. After pushing it, everybody will be able to see it (public by default!). Once you made sure that its fine, you can push the image:

4. Push the image:

```bash
$ docker push YOUR_USERNAME/hello-world-go

```

Here’s what the output looks like:

```bash
$ docker push jonassliplane/hello-world-go
Using default tag: latest
The push refers to repository [docker.io/jonassliplane/hello-world-go]
fd68882aa775: Pushed
21e76093d1a0: Pushed
5f70bf18a086: Mounted from library/docker
0cc9e9f5f204: Mounted from library/golang
a724eb87f2bc: Mounted from library/golang
34622762e218: Mounted from library/golang
1b650b30121e: Mounted from library/golang
450df234f6ef: Mounted from library/golang
31c72bb4adf1: Mounted from library/golang
latest: digest: sha256:bc7 580dd3eafd8cc1 size: 2213

```

Now anyone can pull your image with:

```bash
$ docker pull YOUR_USERNAME/hello-world-go

```

If you push again, you will overwrite the previous image. You probably don’t want to do that, so you should use a specific tag instead of latest.

```bash
$ docker tag hello-world-go YOUR_USERNAME/hello-world-go:v1.0.0
$ docker push YOUR_USERNAME/hello-world-go:v1.0.0

```

Usually you would do this in an automated build process, we will cover that in the next few days!

## The DockerHub Dashboard

DockerHub also provides a way to explore images from other users.

![Explore DockerHub](https://adventofdocker.com/_astro/dockerhub-explore.DOUVx2r3_ZKk0GE.webp)

The explore page is a great way to find images that other people have published! Check them out, maybe you can find something useful!

DockerHub provides a web interface where you can view your own repositories and manage your images.DockerHub

![Dashboard](https://adventofdocker.com/_astro/dockerhub.CG-d7QxG_Z275b4t.webp)

## Best Practices

When working with DockerHub, follow these best practices:

1. Use Specific Tags Instead of relying on latest:

```bash
$ docker tag myapp username/myapp:v1.0.0

```

2. Document Your Images

   - Add a good README
   - Include usage examples
   - List environment variables

3. Regular Updates

   - Keep base images updated
   - Rebuild regularly for security
   - Maintain version tags
