# REST API 
### By **Daffa' Alexander**: for Deall Backend Engineer Application Process (Technical Assessment)

<br>

## Description
This is a REST API mini-project created for Deall Application Proccess for Backend Engineer position. The project is created in Go Programming Language and uses MySQL as the RDBMS. Without further ado, lets jump right in into the contents.

The project is using Uncle Bob's Clean Architecture of Systems. See more [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

<br>

## Tech Stacks
As previously mentioned, the project is created using **Go**  Programming Language and **MySQL** RDBMS. The Technologies/Libraries use in this project are followings:
 * [Echo](echo.labstack.com) Web Framework
 * [GORM](https://gorm.io/) for the ORM
 
 Testing:
 * [vektra/mockery](https://github.com/vektra/mockery) for creating some mocks
 * [stretchr/testify](https://github.com/stretchr/testify) for unit testing
 
 Utilities:
 * [spf13/viper](https://github.com/spf13/viper) for environment configurations
 * [codegangsta/gin](https://github.com/codegangsta/gin) for live reloading

Deployment:
 * Docker
 * [Minikube]() for Local Kubernetes Cluster

<br>

## Accessing The API
The API Documentation can be accessed [here](https://documenter.getpostman.com/view/17548510/UzJFudHb).

### Local
Running the project locally is pretty straightforward. Follow the How To Run The Project part and access the following address with the route provided in the [API Documentation](https://documenter.getpostman.com/view/17548510/UzJFudHb)
> http://localhost:8080/

### Public IP Address
---

<br>

## How To Run The Project
The project is (currently) can only be ran locally.

#### Prerequisites:
* The local machine has already had docker and minikube installed (Check here for docker installation and her for minikube installation)

Assumming that the project is already on the local machine

```bash
# starting minikube using docker
$ minikube start --driver=docker

# applying yaml for mysql application deployment
# secrets, configmap, deployment and service are 
# inside the same file
$ kubectl apply -f mysql.yaml

# applying yaml for api application deployment
$ kubectl apply -f api.yaml

# check if the application is running
$ kubectl get all
$ kubectl get pods

# check if the application is running
$ kubectl get all
$ kubectl get pods

# forward service to be LOCALLY accessible
# With the following command, minikube 
# will return a url for accessing the API.
# The API is can then be tested on postman.
$ minikube service api-service --url

# stop minikube
$ minikube stop
```

