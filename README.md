[![learning-cloud-native-go/myapp](https://img.shields.io/github/stars/learning-cloud-native-go/myapp?style=for-the-badge&logo=go&logoColor=333333&label=learning-cloud-native-go/myapp&labelColor=f9f9f9&color=00ADD8)](https://github.com/learning-cloud-native-go/myapp)
[![learning-cloud-native-go.github.io](https://img.shields.io/github/stars/learning-cloud-native-go/learning-cloud-native-go.github.io?style=for-the-badge&logo=go&logoColor=333333&label=learning-cloud-native-go.github.io&labelColor=f9f9f9&color=00ADD8)](https://github.com/learning-cloud-native-go/learning-cloud-native-go.github.io)
[![learning-rust.github.io](https://img.shields.io/github/stars/learning-rust/learning-rust.github.io?style=for-the-badge&logo=rust&label=learning-rust.github.io&logoColor=333333&labelColor=f9f9f9&color=F46623)](https://learning-rust.github.io)

[![github.com](https://img.shields.io/badge/dumindu-866ee7?style=for-the-badge&logo=GitHub&logoColor=333333&labelColor=f9f9f9)](https://github.com/dumindu)
[![buymeacoffee](https://img.shields.io/badge/Buy%20me%20a%20coffee-dumindu-FFDD00?style=for-the-badge&logo=buymeacoffee&logoColor=333333&labelColor=f9f9f9)](https://www.buymeacoffee.com/dumindu)

---

# Learning Cloud Native Go - Workspace (Draft)

🌱 Cloud Native Application Development is one way of speeding up the building of web applications using microservices, containers, and orchestration tools.

This Go workspace includes:
1. A Dockerized RESTful Book API with CRUD functionality.
2. A Database migrator CLI built with Goose.
3. A mono-repo folder structure 
   1. To support multiple API applications/ services, CLI applications/ tools
   2. To support shared config, model, repository packages between multiple applications.

> **🗂️ Database migrations**
> 
> dbmigrate --dir=../../../services/apis/bookapi/migrations up

> **🧭️ OpenAPI Specification**
>
> swag init -g ./services/apis/bookapi/cmd/app/main.go -o docs/openapi/bookapi -ot yaml --v3.1 --parseDependency

## 📁 Project structure (Expected)

```shell
├── README.md
│
├── apps # TODO: Web and native apps
│   └── web
│       ├── backend     # React: admin facing web app 
│       └── frontend    # React: customer facing web app
│
├── services # TODO: API and serverless apps
│   ├── apis
│   │   ├── userapi     # Go module: User API
│   │   └── bookapi     # Go module: Book API ✅Implemented
│   │
│   └── lambdas
│       ├── userdbmigrator      # Go module: user-migrate-db - Lambda
│       ├── bookdbmigrator      # Go module: book-migrate-db - Lambda
│       ├── bookzipextractor    # Go module: book-extract-zip - Lambda
│       └── bookcsvimporter     # Go module: book-import-csv - Lambda
│
├── tools # TODO: CLI apps
│   └── db
│       └── dbmigrate # Go module: Database migrator ✅Implemented
│
├── infrastructure # TODO: IaC 
│   ├── dev
│   │   └── localstack  # Infrastructure for dev environment for Localstack
│   │
│   └── terraform
│       ├── environments
│       │   ├── dev     # Terraform infrastructure for development environment
│       │   ├── stg     # Terraform infrastructure for staging environment
│       │   └── prod    # Terraform infrastructure for production environment
│       ├── global
│       │   ├── iam     # Global IAM roles/policies
│       │   └── s3      # Global S3 infrastructure like log-export
│       └── modules
│           ├── security    # IAM, SSO, etc per service
│           ├── networking  # VPC, subnets
│           ├── compute     # ECS, Fargate task definitions, Lambda
│           ├── serverless  # Lambda functions
│           ├── database    # RDS
│           ├── storage     # S3
│           ├── messaging   # SQS, EventBridge
│           └── monitoring  # CloudWatch dashboards, alarms
│
├── shared # Shared Go and TypeScript packages
│   ├── go
│   │   ├── configs       # Go module: shared between multiple applications ✔️ Partially Implemented
│   │   ├── errors        # Go module: shared between multiple applications ✔️ Partially Implemented
│   │   ├── models        # Go module: shared between multiple applications ✔️ Partially Implemented
│   │   ├── repositories  # Go module: shared between multiple applications ✔️ Partially Implemented
│   │   └── utils         # Go module: shared between multiple applications ✔️ Partially Implemented
│   │
│   └── ts # TODO
│
└── compose.yml
```
