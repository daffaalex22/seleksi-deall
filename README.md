# REST API 
### By **Daffa' Alexander**: for Deall Backend Engineer Application Process (Technical Assessment)

## Key Points
 * The api hasn't deployed on a VPS as requested (only locally using minikube, yaml files are attached). To compensate for this, the author adds some additional features including:
    - Clean Code Architecture
    - Unit Testing
    - Additional endpoint for getting user by id
 * The screenshot for prove of local deployment is attached [here](https://google.com) inside the screenshots folder
 * The API Documentation is accessible [here](https://documenter.getpostman.com/view/17548510/UzJFudHb)
 * Admin credential is the following
 > email: admin@gmail.com
 > 
 > password: deall123

 * User credential is the following
 > email: user@gmail.com
 > 
 > password: deall123

## Description
This is a REST API mini-project created for Deall Application Proccess for Backend Engineer position. The project is created in Go Programming Language and uses MySQL as the RDBMS. Without further ado, lets jump right in into the contents.

The project is using Uncle Bob's Clean Architecture of Systems. See more [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
<br>
## Tech Stacks
As previously mentioned, the project is created using **Go**  Programming Language and **MySQL** RDBMS. The Technologies/Libraries use in this project are followings:
 * [Echo](echo.labstack.com) Web Framework
 * [MySQL](https://www.mysql.com/) for the RDBMS
 * [GORM](https://gorm.io/) for the ORM
 
Testing & Utilities:
 * [vektra/mockery](https://github.com/vektra/mockery) for creating some mocks
 * [stretchr/testify](https://github.com/stretchr/testify) for unit testing
 * [spf13/viper](https://github.com/spf13/viper) for environment configurations

Deployment:
 * Docker
 * [Minikube](https://minikube.sigs.k8s.io/docs/) for Local Kubernetes Cluster

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
* Docker and minikube installed (Check here for docker installation and her for minikube installation)
* Kubectl installed (minikube usually has kubectl dependency so additional installation is not necessary)

1. Starting minikube using docker
```
minikube start --driver=docker
```

2. Applying yaml for mysql application deployment. Secrets, configmap, deployment and service are inside the same file.
```
kubectl apply -f https://raw.githubusercontent.com/daffaalex22/seleksi-deall/main/mysql.yaml
```

3. Applying yaml for API application deployment
```
kubectl apply -f https://raw.githubusercontent.com/daffaalex22/seleksi-deall/main/api.yaml
```

4. Check if the application is running
```
kubectl get all
```
Or
```
$ kubectl get pods
```

5. Forward service to be LOCALLY accessible. With the following command, minikube will return a url for accessing the API. The API is can then be tested on postman.
```
minikube service api-service --url
```

6. Stop minikube
```
minikube stop
```

