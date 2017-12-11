# Small Golang Web server in Docker for AWS EC2 testing

### How to run:

``docker run -d -p 8080:8080 --rm --name ec2-testing vensder/ec2-testing``

Open [http://localhost:8080/Hi%20there!](http://localhost:8080/Hi%20there!) or any other random path after slash in your web browser.

View logs:

``sh
docker logs ec2-testing
``

Stop container:

``sh
docker stop ec2-testing
``

### How to build your own image if you don't have Go (but have a docker):

``sh
git clone https://github.com/vensder/ec2-testing.git
``

``sh
cd ec2-testing
``

``sh
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8-alpine go build -v ec2-testing.go
``

Make binary file smaller:

``sh
strip ec2-testing
``

``sh
mkdir -p bin/ && mv ec2-testing bin/
``

Build docker image:

``sh
docker build -t ec2-testing .
``

So you got smallest docker image with built-in web-server, it has size less than 8 MB!
You can use it even in embedded systems.

