# Day 17 - Introduction to Github Actions

Hey there! Welcome to Day 17 of Advent of Docker. Yesterday we learned about DockerHub and how to share our images. Today we’re going to take it one step further and automate everything with GitHub Actions!

If you’ve ever thought “I wish I didn’t have to manually build and push my Docker images every time”, then you’re in for a treat. GitHub Actions is going to be your new best friend for automation :D

## What is GitHub Actions?

Every time you push code, create a PR, or even on a schedule, Github Actions can automatically:

- Build your Docker images
- Run your tests
- Push images to registries
- Deploy your applications
- And pretty much anything else you can think of!

The best part? It’s completely free for public repositories, and you get a really generous free tier for private ones too. No more excuses for not automating your workflows!

## Key Concepts

Before we dive in, let’s understand some key terms:

- Workflow: A configurable automated process made up of one or more jobs
- Job: A set of steps that execute on the same runner
- Step: An individual task that can run commands or actions
- Action: A reusable unit of code that can be shared across workflows
- Runner: A server that runs your workflows

## Setting Up Our First Workflow

Let’s create a simple workflow that builds a Docker image. First, we need to create a .github/workflows directory in our repository:

```bash
mkdir -p .github/workflows

```

Then, create a new file called docker-build.yml in that folder:

```Dockerfile
name: Docker Build

on:
  push:
    branches:
      - main
  workflow_dispatch:

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Login to GitHub Container Registry
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Build and push Docker image
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            ghcr.io/your-name/your-repo:latest
            ghcr.io/your-name/your-repo:${{ github.sha }}

```

Make sure to replace your-name and your-repo with your actual GitHub username and repository name.

You will also have to enable read and write permissions for the GitHub Container Registry in your repository settings:

![Permissions](https://adventofdocker.com/_astro/permissions.BykLA6xh_2nJ9pD.webp)

Now with that done, whenever you push code, Github actions will do:

1. Checkout your code (git clone)
2. Login to the docker registry (docker login)
3. Build your image (docker build)
4. Push your image to the docker registry (docker push)

You can see all of your workflow runs in the “Actions” tab:

![workflows](https://adventofdocker.com/_astro/workflows.B0mnZbnV_Z26L1ap.webp)

If you click on a workflow run you can see the jobs:

![jobs](https://adventofdocker.com/_astro/workflow.mgROGye7_Z26szf.webp)

Clicking on a job will show you the steps and detailed logs.

![steps](https://adventofdocker.com/_astro/job.FDNCLy2Q_EeiW6.webp)

If you want to see the output of your workflow, check out the packages page of your repository (you can find that on the right sidebar of your repository).

Here you get the instructions on how to pull your image and what images exist!

![packages](https://adventofdocker.com/_astro/packages.BK_UdsMo_Z1EqFTr.webp)

Now at this point I could write a 500 page documentation on Github Actions, but I don’t think this will help you. Instead, take this as a starting point and try to automate some of your own workflows. If you have any struggles, check out the Github Actions Documentation.

## Best Practices

When working with GitHub Actions, keep these tips in mind:

1. Use the latest versions for actions (e.g., @v4 instead of @v3)
2. Leverage caching to speed up builds
3. Keep secrets secure using repository secrets!
4. Use matrix builds for testing across multiple versions/platforms
5. Monitor usage to stay within free tier limits
