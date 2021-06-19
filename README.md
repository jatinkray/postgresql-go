## GoWeb-PostgreSQL

Based on project [GoWeb-PostgreSQL](https://adaickalavan.github.io/portfolio/golang-web-application-with-postgresql/)

This application is written in Golang + HTML + JS

We use it with Terraform code for some Workshops, Meetups, Labs, etc.

#### How to use it

Here is a list of variables which you need to be add to env variables:

export DBPORT= "Your DB Port"

export DBHOST= "Your DB Host"

export DBUSER= "Your DB user"

export DBPASS= "Your DB Password"

export DBNAME= "Your DB Name"

#### Application Run local using docker

```shell
docker compose build 
docker compose up
```

The application will be running at http://localhost:8080

#### Deploy application to kubernetes using helm
Before executing to kubernetes cluster . kindly review [charts/postgresgoapp/values.yaml](https://github.com/jatinkray/postgresql-go/blob/main/charts/postgresgoapp/values.yaml) for application config
```shell
kubectl create namespace myapp
helm upgrade --install postgresgoapp charts/postgresgoapp --namespace myapp
```

Note: if you are deploying the chart on minikube / docker-desktop. kindly update your /etc/hosts file for kubernetes ingress.
e.g. 
```text
127.0.0.1 localhost
127.0.0.1 goapp.local
```