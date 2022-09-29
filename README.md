# SPACE PROJECT
Agnostic API built in Go Lang, made for purposes of test building blocks.

You should set .envrc and .aws_credentials


## Features:
* Init of basic project
* REST API 
    Create:     POST /projects
    Delete:     DELETE /projects/:id
    Read:       GET /projects/:id
    Read all:   GET /projects/
    Update:     PUT /projects/:id
    
* Persistent storage adapters (disk, extensible to s3 etc.)
* Create deploy to ECS
    * Build in Docker
    * Make file with AWS push to EKR
    * Deploy to ECS, 3 instances different regions
    * APP Load Balancer handling HTTPS termination with valid SSL/TLS Certificate
    * Attach EFS disk to 3 instances

* Add HTTP Basic Auth % 
* Create Unit Tests % 
* 


## Build
```make build```

## Deploy
```make deploy```

