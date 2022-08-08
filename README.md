## HEXAGONAL USER MICROSSERVICE

A REST API backend application using hexagonal architecture to handle login logic and user CRUD and JWT Token authentication

### How to start?

* Create .env file containing the follow vars.

  * SERVICE_PORT
  * SERVICE_HOST
  * POSTGRES_HOST
  * POSTGRES_PORT
  * POSTGRES_USER
  * POSTGRES_PASSWORD
  * POSTGRES_DB_NAME

* Clone the service locally and build image running **`make build`**.
* Run the service running **`make up`**.
* Make sure service is up and running
* Now you can send request to the service using swagger on <http://127.0.0.1:8001/api/v1/swagger> or a postman  collection UserMicrosservices.postman_collection.json

* Run test running **`make test`**.
