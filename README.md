# Small Golang Web server in Docker for AWS EC2 testing

If you run it on AWS EC2 Environment, Web server show you page with some meta information about current EC2 instance:

* ami-id
* hostname
* instance-id
* instance-type
* local-hostname
* local-ipv4
* public-ipv4
* public-keys
* security-groups

And environment variables inside the docker container too.


### How to run:

``sh
docker run -d -p 8080:8080 --rm --name ec2-testing vensder/ec2-testing
``

Open [http://localhost:8080/Hi%20there!](http://localhost:8080/Hi%20there!) or any other random path after slash in your web browser.

View logs:

``sh
docker logs ec2-testing
``

Stop container:

``sh
docker stop ec2-testing
``
