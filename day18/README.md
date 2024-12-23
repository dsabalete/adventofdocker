# Day 18 - Deploying Docker to the Cloud

Hey there! Welcome to Day 18 of Advent of Docker. Yesterday we learned about Github Actions and automated Docker builds. Today we’ll explore how to deploy our Docker containers to the cloud and share it with the world!

## Why Deploy to the Cloud?

By deploying to the cloud I mean putting your application on a server somewhere and making it accessible to the public. There are usually many steps involved in this process. Where do you get a server? Do you need a domain? How do you manage a server?

You can of course do everything yourself, but this would require an additional 30 days of blog posts. So today we will focus on deploying to something called a “Platform as a Service” (PaaS). PaaS providers take care of all the server management for you, usually you simply provide a Docker image or Github Repository and they take care of the rest. For example, a typical PaaS provider would automatically deploy your Docker image on every git push, provide you with a free subdomain and take care of the SSL certificates.

There are many providers out there, and Docker enables you to deploy your application to basically any cloud provider. But because I am the co-founder of a PaaS company called sliplane.io, we are going to focus on that! :)

![Cloud meme](https://adventofdocker.com/_astro/this_is_the_cloud.BHVTuafk_1Jg7ti.webp)

## Setup

I assume you have your Github repository from yesterday ready. If not, create a new one and upload the `Dockerfile` and `main.go` from day 17. As a reminder, the `Dockerfile` should look like this:

```Dockerfile
FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o main .

EXPOSE 8080
CMD ["go", "run", "main.go"]

```

And our main.go should look like this:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!") // this will be printed as a response when you request /
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil) // this will start the HTTP server (blocking operation) and listen on port 8080
}

```

If you have both of these files in your Github repository, we can get started!

## Let’s deploy!

This whole thing should take you less than 5 minutes, stop the timer now! 🏃‍♂️⏳

First, go to sliplane.io and signup. You don’t need a credit card to get started and the first 48 hours are free. Nothing to lose!

Once you have signed up, you will be redirected to the dashboard. Here you can click on your Default Project and deploy your first service. We automatically created a Server for you in the background!

![Project](https://adventofdocker.com/_astro/project.DNhcaP-p_ZTBr4g.webp)

A project is a logical group of multiple services. For example, you could have a project called “My Awesome Shop” and then have one service for your Frontend, one for the Backend and one for the Database.

Now, click on the “Deploy Service” button. A wizard will guide you through the process of deploying your service, first select the server that we already created for you. In the next step, you can select if you want to deploy from a Github repository or from a Docker container registry, like DockerHub or GitHub Container Registry. In this case, we want to deploy from a Github repository, so select “Repository”.

Now you will see this:

![Deploy Service](https://adventofdocker.com/_astro/deploy.DvVHLd6b_ZMMxWu.webp)

You will first have to select the repository. Because this is the first time you use Sliplane, you will have to give us permission to access your repository. This will enable us to read your Dockerfile and deploy your application.

Then, you can pick the branch and Dockerfile you want to deploy. Usually we are pretty smart and will pick the right branch and Dockerfile for you, but you can also change it here!

After selecting the repository, you can pick if you want to make it public or not. Making it public will allow you to access it from the internet, which is what we want. You can simply keep the default settings.

Then you can set some environment variables (basically what you learned on Day 9). You can also connect some volumes for persistant storage, but you also don’t need that right now. You can keep all the default settings, and then click on “Deploy” at the bottom.

Now Sliplane will clone your repository, build your Docker image and deploy it. This will take maybe 1-2 minutes. You can monitor the deployment progress by opening the “Logs” tab.

![logs](https://adventofdocker.com/_astro/logs.gCHVGczN_Z2qELok.webp)

When the deployment is finished your service settings should look like this:

![service](https://adventofdocker.com/_astro/service.DKaSzetR_Z1qrL9D.webp)

You will see the Domain of your service, which you can use to access it from the internet. In my example, you can access it at https://test-repo.sliplane.app

Now on every git push we will build your image, deploy it and make sure that its working. No manual steps involved in deploying your application!

Sliplane has a lot more features. You can one-click deploy databases, add persistant storage, add custom domains, and much more. You have 48 hours to try it out for free, take it for a spin! If you have any questions you can simply use the chat on the bottom right or consult the [documentation](https://docs.sliplane.io/).

## Conclusion

Today you learned how to deploy Docker containers to the cloud using sliplane.io. While there are many ways to deploy Docker containers, using a PaaS is often the fastest way to get started. It handles the infrastructure, SSL certificates, and automatic deployments for you.

The key points:

- PaaS providers manage servers and deployment for you
- You only need to provide a Docker image or repository
- Automatic deployments happen on every git push
- Free subdomains and SSL certificates are usually included
