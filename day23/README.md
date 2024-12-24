# Day 23 - Resources

## How containers work by Julia Evans

[Julia Evans’ zine about containers](https://wizardzines.com/zines/containers/) is an absolute gem! It’s a fun, illustrated guide that breaks down complex container concepts into digestible pieces. Julia has a unique talent for making technical topics approachable and enjoyable. The zine covers everything from namespaces and cgroups to container runtime internals, all with delightful hand-drawn illustrations. If you’re a visual learner or just want a fresh perspective on container technology, this is a must-read, although it is only available if you pay for it!

## Docker Documentation

The [official Docker documentation](https://docs.docker.com/) is your comprehensive guide to all things Docker. While it might seem obvious, it’s worth highlighting that this is one of the best-maintained technical documentations out there.

## Docker Source Code

For those interested in the deeper technical aspects, Docker’s source code is available in two main repositories:

[BuildKit](https://github.com/moby/buildkit/) is the engine behind Docker’s build system. It’s a fascinating look into how Docker images are constructed, with features like:

- Concurrent dependency resolution
- Efficient caching mechanisms
- Advanced build primitives

[Moby](https://github.com/moby/moby) is the open-source foundation of Docker Engine. Diving into this codebase will help you understand:

- Container runtime implementation
- Image management
- Networking architecture
- Storage drivers

I spend a lot of time reading the source code of Docker and BuildKit. Highly recommend it if you want to understand how Docker works under the hood!

## Docker Registry Specification

The [Distribution Specification](https://distribution.github.io/distribution/) is essential reading if you want to understand how Docker images are stored and distributed, but otherwise it’s not that interesting.

## OCI Specification

The [Open Container Initiative (OCI) Specification](https://specs.opencontainers.org/image-spec/) defines the industry standards for container formats and runtimes. This is crucial reading for:

- Understanding container image formats
- Learning about container runtime requirements
- Grasping the standardization that enables container portability
- Exploring the technical requirements for container platforms

These specifications form the foundation of modern container technology and are essential reading if you want to deeply understand how containers work at a standardized level.

## My Blog :)

Well, can’t leave out my own blog, right? I write a lot about Docker, mostly real world examples etc. Check out some of my best posts:

- [Say Goodbye to Docker Volumes](https://dev.to/code42cate/say-goodbye-to-docker-volumes-j9l)
- [5 Cheap Ways to Deploy Docker Containers](https://dev.to/code42cate/5-cheap-ways-to-deploy-docker-containers-3dke)
- [Don’t build Docker apps without this flag](https://dev.to/code42cate/dont-build-docker-apps-without-this-flag-48kg)
