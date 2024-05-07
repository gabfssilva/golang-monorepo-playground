VERSION 0.8

all-artifact:
    BUILD ./libs/result+artifact

all-unit-test:
    BUILD ./libs/hello+unit-test
    BUILD ./services/one+unit-test
    BUILD ./services/two+unit-test

all-docker:
    BUILD ./services/one+docker
    BUILD ./services/two+docker

all-release:
    BUILD ./services/one+release
    BUILD ./services/two+release

dev-up:
    FROM +all-docker
    LOCALLY
    RUN docker-compose up

dev-down:
    LOCALLY
    RUN docker-compose down
