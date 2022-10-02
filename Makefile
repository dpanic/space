REPOSITORY_URI=442267940013.dkr.ecr.eu-central-1.amazonaws.com
REPOSITORY=space
CLUSTER=arn:aws:ecs:eu-central-1:442267940013:cluster/space
SERVICE=space
REGION=eu-central-1
AWS_SHARED_CREDENTIALS_FILE=.aws_credentials

all: build deploy


run:
	go run main.go

build:
	$(eval COMMIT_HASH=$(shell git rev-parse HEAD | cut -c1-7))
	$(eval TIME_STAMP=$(shell date +%Y.%m%d.%H%M))

	docker buildx build \
		-t $(REPOSITORY_URI)/${REPOSITORY}:latest \
		-t $(REPOSITORY_URI)/${REPOSITORY}:$(COMMIT_HASH) \
		-t $(REPOSITORY_URI)/${REPOSITORY}:$(TIME_STAMP) \
		--build-arg LDFLAGS="-w -s -X \"main.Version=${COMMIT_HASH}\"" \
		. -f ./Dockerfile

deploy: build
	AWS_SHARED_CREDENTIALS_FILE=.aws_credentials aws ecr get-login-password --region "${REGION}" | docker login --username AWS --password-stdin $(REPOSITORY_URI)

	docker push $(REPOSITORY_URI)/${REPOSITORY}:latest
	docker push $(REPOSITORY_URI)/${REPOSITORY}:$(COMMIT_HASH)
	docker push $(REPOSITORY_URI)/${REPOSITORY}:$(TIME_STAMP)

	$(eval TASKS=$(shell AWS_SHARED_CREDENTIALS_FILE=.aws_credentials aws ecs list-tasks --cluster "${CLUSTER}" --service "${SERVICE}" --output text --region "${REGION}" --query "taskArns"))

	for task in $(TASKS); do \
		AWS_SHARED_CREDENTIALS_FILE=.aws_credentials aws ecs stop-task --cluster "${CLUSTER}" --task "$$task" --region "${REGION}" | jq; \
	done

