# SPACE PROJECT
[![Go Report Card](https://goreportcard.com/badge/github.com/dpanic/space)](https://goreportcard.com/report/github.com/dpanic/space)

Agnostic API built in Go Lang, made for purposes of test building blocks.

**Building limits**: The areas (polygons) on your site where you are allowed to build

**Height plateaus**: Areas (polygons) on your site with different elevation. In reality, your building site is a continuous irregular terrain, but before building, you level your
terrain into discrete plateaus with constant elevation.

**Building splits**: Corresponding areas (polygons) of building limits that match certain heigh plateaus


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

## Testing

### Option 1
1. If you're familiar with Go, you can perform Unit tests in splits_test.go:
```/usr/local/go/bin/go test -timeout 30s -run ^TestSplits$ space/backend/logic -count=1 -v```
2. It will output splits_%d.json file which you can load in https://geojson.io and see your results

### Option 2
1. You can run POSTMAN collection with UPDATE
2. It will output GeoJSON
3. You can copy paste to https://geojson.io

### Demo
Building limits matched all 3 heigh plateaus.
<br>
<img src="tests/m_3.png?raw=true" width="80%" />

Building limits matched 2 heigh plateaus, 1 unmatched.
<br>
<img src="tests/m_2_u_1.png?raw=true" width="80%" />

Building limits matched 1 heigh plateaus, 2 unmatched.
<br>
<img src="tests/m_1_u_2.png?raw=true" width="80%" />

2 Building limits matched 4 heigh plateaus, 0 unmatched.
<br>
<img src="tests/m_4.png?raw=true" width="80%" />

## Features
* Init of basic project
* REST API 
    * Create:             POST /projects
    * Read:               GET /projects/:id
    * Read GeoJSON:       GET /projects/:id?o=geojson
    * Update:             PUT /projects/:id
    * Update GeoJSON:     PUT /projects/:id?o=geojson
    * Delete:             DELETE /projects/:id
    * Read all:           GET /projects/
    
Parameter geojson outputs data compatible for geojson.io and similar viewers

* Persistent storage adapters (disk, extensible to s3 etc.)
* Create deploy to ECS
    * Build in Docker
    * Make file with AWS push to ECR
    * Deploy to ECS, 3 instances different regions
    * APP Load Balancer handling HTTPS termination with valid SSL/TLS Certificate
    * Attach EFS disk to 3 instances

* Concurrent access
* HTTP Basic Auth
* Create Postman Collection for DEVELOPMENT and PRODUCTION
* API shows version and last built time, uptime of service
* Render output for geojson on UPDATE and READ API routes

* Implement logic 
    * compare one by one 
    * create intersections and differences
    * add colors to building_splits

* Create Unit Tests


## Build
```make build```

## Deploy
```make deploy```

## Run
```make run```

