## Credit service
A simple golang RESTful service that authenticates user, with react web app has front-end

## Prerequisite
* Docker
* Docker-compose
* Golang* 
* Node*
* Yarn*

 `* - To run in local`

## Commands

To start service
```
docker-compose up -d
```

* If you have latest version of docker you can use
```
docker compose up -d
```

To stop service
```
docker-compose down -v
```

If you face network credit_network declared as external, but could not be found. Please create the network manually using `docker network create credit_network` and try again.

### Features
* Create User - Create user with required details like first name, last name, ssn, dob etc -> /app/register
* Questions - Fetches collections of stored questions and answers from db -> /app/questions
* Submission - Submit collection of questions and user chosen answer for the question along with user auth -> /app/logout
* Order Credit - Order credit for a specific product code along with user auth -> /app/order

> Note: To test these API's please find the postman collection named "credit-report.postman_collection.json" attached in this repo.

### Libaries used
|Name  | Use  |
|:-----| :--- |
|[gin](https://github.com/gin-gonic/gin)| Web frame work|

