### Backend master class [Golang + Postgres + Kubernetes + gRPC]

- Create and manage bank accounts.
- Record all balance changes to each of the accounts.
- Perform a money transfer between 2 accounts.

The programming language we will use to develop the service is Golang, but the course is not just about coding in Go. You will learn a lot of different topics regarding backend web development. They are presented in 6 sections:

1. design the database, generate codes to talk to the DB in a consistent and reliable way using transactions, understand the DB isolation levels, and how to use it correctly in production. Besides the database, you will also learn how to use docker for local development, how to use Git to manage your codes, and how to use GitHub Action to run unit tests automatically.

2. build a set of RESTful HTTP APIs using Gin - one of the most popular Golang frameworks for building web services. This includes everything from loading app configs, mocking DB for more robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.  

3. build your app with Docker and deploy it to a production Kubernetes cluster on AWS. The lectures are very detailed with a step-by-step guide, from how to build a minimal docker image, set up a free-tier AWS account, create a production database, store and retrieve production secrets, create a Kubernetes cluster with EKS, use GitHub Action to automatically build and deploy the image to the EKS cluster, buy a domain name and route the traffics to the service, secure the connection with HTTPS and auto-renew TLS certificate from Let's Encrypt.

4. managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP requests at the same time, embedding Swagger documentation as part of the backend service, partially updating a record using optional parameters, and writing structured logger HTTP middlewares and gRPC interceptors.


## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [DB Docs](https://dbdocs.io/docs)

    ```bash
    npm install -g dbdocs
    dbdocs login
    ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

    ```bash
    npm install -g @dbml/cli
    dbml2sql --version
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### Setup infrastructure

- Create the bank-network

    ``` bash
    make network
    ```

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```
