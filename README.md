# SPACE PROJECT
Agnostic API built in Go Lang, made for purposes of test building blocks.

You should set .envrc and .aws_credentials

aws ecs stop-task --cluster "${CLUSTER}" --task $(aws ecs list-tasks --cluster "${CLUSTER}" --service "${SERVICE}" --output text --query "taskArns[0]")


## Features:
* Init of basic project
* REST API 
    + Create:     POST /projects
    Update:     PUT /projects/:id
    Read:       GET /projects/:id
    + Delete:     DELETE /projects/:id
    Read all:   GET /projects/
    
* Persistent storage adapters (disk, extensible to s3 etc.)
* Create deploy to ECS
    * Build in Docker
    * Make file with AWS push to ECR
    * Deploy to ECS, 3 instances different regions
    * APP Load Balancer handling HTTPS termination with valid SSL/TLS Certificate
    * Attach EFS disk to 3 instances

* Implement logic - 2h %
* Concurrent access - 2h %
* HTTP Basic Auth
* Create Unit Tests - 1h % 
* Create Postman Collection for DEVELOPMENT
* Create Postman Collection for PRODUCTIOn


## Build
```make build```

## Deploy
```make deploy```

