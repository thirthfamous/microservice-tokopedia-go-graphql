# Microservices Tokopedia Go GraphQL Clone 
Microservices Tokopedia Go GraphQL Clone made with Go, MySQL, RabbitMQ, NGINX

---

### Features (Endpoints)
* Register User
* Login User
* Create Product
* Find All Product
* Create Order
* Find All Order
* Find By Order ID
* Customer Pay


### Prerequisite
* Install Docker at [Docker official site](https://www.docker.com/products/docker-desktop/)

---

### Installation
1. Run Docker
2. Run the commands 
```sh
// clone the project
git clone https://github.com/thirthfamous/microservice-tokopedia-go-graphql.git

// go to the project directory
cd microservice-tokopedia-go-graphql

// build images
docker-compose build --no-cache 

// run the app
// hit the api at localhost:8080
docker-compose up -d

// Wait around 1 minute, because RabbitMQ service is slow to start


// run the test
// go to specific service
cd customer/test
cd product/test
cd order/test
cd payment/test

// run it
go test test/

```

3. Import Tokopedia Go GraphQL.postman_collection.json to your Postman
4. Enjoy the app

---
### Service Pattern in Action
![Alur Choreography SAGA pattern drawio (1)](https://user-images.githubusercontent.com/30696403/168213059-c99b186e-227b-4011-a57e-690d1abd8c14.png)


---
### Architecture
![microservices tokopedia clone drawio (1)](https://user-images.githubusercontent.com/30696403/168212961-8045a1ce-3446-4860-a2b0-1b26b208d1e8.png)


### Error 
Q : I got 502 bad gateway, how to fix it

A : Run command "docker-compose restart nginx-proxy" without double quotes
