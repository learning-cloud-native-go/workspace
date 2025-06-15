# Learning Cloud Native Go - Workspace

🌱 Cloud Native Application Development is one way of speeding up the building of web applications using microservices, containers, and orchestration tools.

> **🗂️ Database migrations**
> 
> dbmigrate --dir=../../../services/apis/bookapi/migrations up

> **🧭️ OpenAPI Specification**
>
> swag init -g ./services/apis/bookapi/cmd/app/main.go -o docs/openapi/bookapi -ot yaml --v3.1 --parseDependency
