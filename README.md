
# Shopping API using Gin, Golang, JWT, Dependency Injection, Unit testing

## Description
This is an shopping cart service implementation in Go (Golang)+ Gin projects.

This project has  4 Domain layer :
 * Models Layer 
 * Repository Layer
 * Business logic/Usecase Layer  
 * API/Delivery Layer

### How To Run This Project

```bash
#move to directory
cd $GOPATH/src/github.com/karuppaiah

# Clone into YOUR $GOPATH/src
git clone https://github.com/karuppaiah/shopping.git

#move to project
cd shopping

# Run the script
sh execute.sh

# Data populate and setup DB
sh createData.sh

At this time, you will have a new data.db created in root directory. Change the DB if needed.

Site runs at http://127.0.0.1:8080/ping

Postman V2 request: https://github.com/karuppaiah/shopping/blob/master/golang%20shopping.postman_collection

Always get the JWT token and use them in authorization header for response.



# TODO :
1. User management
2. Docker image and publish in hub.docker.com - In progress
3. GOlang IOC
4. redis storage for JWT token
5. Move to Mongo DB
6. Write more unit testing for more coverage
7. React frontend
8. Utilites for message Q like kafka(producer and consumer)
9. Tree structure for promotion to enable double source product rule
10. Stress testing scripts



```


