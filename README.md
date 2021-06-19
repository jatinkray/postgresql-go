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

#### Additional info

You can check the server where the application runs on using http:/yourIP/server endpoint 

#### Download and install

curl -L https://github.com/s-marchenko/postgresql-go/releases/download/v1.0.0/website_linux_amd64 --output postgresql
chmod 775 postgresql