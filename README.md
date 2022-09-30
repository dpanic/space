# SPACE PROJECT
[![Go Report Card](https://goreportcard.com/badge/github.com/dpanic/space)](https://goreportcard.com/report/github.com/dpanic/space)

Agnostic API built in Go Lang, made for purposes of test building blocks.

You should set .envrc and .aws_credentials


## Flow
* Create project, you receive ID based on project name
* Update project by ID with height_plateaus and building_limits data
* System will calculate split_building_limits and save it to disk
* Revision of data will be increased, in order to prevent overwriting each others data

## Architecture
Client -> ALB -> ECS -> Docker

* ALB: App Load Balancer (SSL/TLS termination)
* ECS: Elastic Cloud Service (3 instances, with auto scaling)
* Docker (Go Binary) + NFS (AWS EFS)


## Features
* Init of basic project
* REST API 
    * Create:     POST /projects
    * Read:       GET /projects/:id
    * Update:     PUT /projects/:id
    * Delete:     DELETE /projects/:id
    * Read all:   GET /projects/
    
* Persistent storage adapters (disk, extensible to s3 etc.)
* Create deploy to ECS
    * Build in Docker
    * Make file with AWS push to ECR
    * Deploy to ECS, 3 instances different regions
    * APP Load Balancer handling HTTPS termination with valid SSL/TLS Certificate
    * Attach EFS disk to 3 instances

* Implement logic - 2h %
* Concurrent access
* HTTP Basic Auth
* Create Postman Collection for DEVELOPMENT
* Create Postman Collection for PRODUCTION %
* API shows version and last built time, uptime of service
* Create Unit Tests - 1h % 


## Build
```make build```

## Deploy
```make deploy```

## Run
```make run```

