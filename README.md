# REST API 
#### By: **Daffa' Alexander**, for Deall Backend Engineer Application Process (Technical Assessment)
---
## Key Points
 * The API hasn't deployed on a VPS as requested (only [locally using minikube](https://github.com/daffaalex22/seleksi-deall#how-to-deploy-the-project-locally), yaml files for [the API](https://github.com/daffaalex22/seleksi-deall/blob/main/api.yaml) and [the database](https://github.com/daffaalex22/seleksi-deall/blob/main/mysql.yaml) are attached). To compensate for this, the author adds some additional features including:
    - [Clean Code Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html);
    - Unit Testing as shown [here](https://github.com/daffaalex22/seleksi-deall/edit/main/README.md#unit-testing);
    - Additional endpoint for getting user by id as shown [here](https://github.com/daffaalex22/seleksi-deall/blob/main/app/routes/route.go#L35);
    - The usage of [JSON Web Token](https://jwt.io/) as shown [here](https://github.com/daffaalex22/seleksi-deall/blob/main/app/middlewares/auth.go).

 * The screenshot for proof of local deployment is attached [here](https://github.com/daffaalex22/seleksi-deall/tree/main/local-deployment-screenshots) inside the local-deployment-screenshots folder. View it according to the sorted order.
 * The Architecture Diagram can be accessed on [Figjam](https://www.figma.com/file/AtXxHmWi7XDiYXvN7MscQb/Architecture-Diagram?node-id=0%3A1). It is also attached inside the root folder as [local-architecture-diagram.png](https://github.com/daffaalex22/seleksi-deall/blob/main/local-architecture-diagram.png).
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
This is a REST API mini-project created for Deall Application Proccess for Backend Engineer position. The project is created in Go Programming Language and uses MySQL as the RDBMS. The project is using Uncle Bob's Clean Architecture of Systems. See more [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). Without further ado, lets jump right in into the contents.


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
Deploying the project locally is pretty straightforward. Follow the [How To Deploy The Project Locally](https://github.com/daffaalex22/seleksi-deall#how-to-deploy-the-project-locally) part and access with routes provided in the [API Documentation](https://documenter.getpostman.com/view/17548510/UzJFudHb). 

#### Running Locally
It can be also ran without Docker or Kubernetes. 
1. Clone the repo.
2. Provides the following environment variables. Make sure MySQL is installed and schema created. Type those values in to their respective variables.
```
DB_URL=localhost:3306
DB_USER=<<YOUR_MYSQL_USERNAME>>
DB_PASSWORD=<<YOUR_MYSQL_PASSWORD>>
DB_NAME=<<YOUR_CREATED_DB>>
```
Or use the following launch.json for running on debug mode on vscode:
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "cwd": "${workspaceFolder}",
            "env": {
                "DB_URL": "localhost:3306",
                "DB_USER": <<YOUR_MYSQL_USERNAME>>,
                "DB_PASSWORD": <<YOUR_MYSQL_PASSWORD>>,
                "DB_NAME": <<YOUR_CREATED_DB>>,
            },
            "args": []
        }
    ]
}
```
3. Run using debug mode on vscode or type:
```
go run main.go
```
4. Then access the following url.
> http://localhost:8080/

#### Unit Testing
Run the following command to test the API. 
> The following command is applicable for Windows Powershell. Other OS might edit some of the command for it to run.
```bash
go test ./... -coverprofile="cover.out"
go test ./business/ -coverprofile="cover.out"
go tool cover -html="cover.out"
```


### Public IP Address

<br>

## How To Deploy The Project Locally
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
kubectl get pods
```

5. Forward service to be LOCALLY accessible. With the following command, minikube will return a url for accessing the API. The API is can then be tested on postman.
```
minikube service api-service --url
```

6. Stop minikube
```
minikube stop
```

