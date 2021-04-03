# Nginx with NodeJS Challenge

Challenge made for the fullcycle developer course,
the objective of this challenge is to make a NodeJS application, that receive's a http GET request and insert a random
name into the mysql database and render every new register.

That setting three docker containers (Node, Nginx, MySql), and making them comunicate through a network, also utilizing
nginx as a reverse-proxy to make the NodeJS app container receive the requests from the host, . 

## execution:
`docker-compose up -d`