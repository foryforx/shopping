
# Shopping API using Gin, Golang, JWT, Dependency Injection, Unit testing

## Description
This is an shopping cart service implementation in Go (Golang)+ Gin projects.
To create a simple shopping cart service for a small product store and identify, design & implement relevant APIs. 

Requirements & User stories
1.	Add and remove items (with quantities) to the cart
2.	View the current cart, showing items, quantities and total price with promotions
Current promotions
The shop at any time has a variety of different promotions available. These need to be easily changed or added to at any time. Current promotions include:
1.	If you buy 2 or more trousers, you get 15% off belts and shoes.
2.	If you buy 2 shirts, each additional shirt only costs $45.
3.	If you purchase 3 or more shirts, all ties are half price.
Current inventory availability and prices

| ItemName  | Stock | Price |
| ------------- | ------------- | ------------- |
| Belts  | 10  | $20  |
| Shirts  | 5  | $60  |
| Suits  | 2  | $300  |
| Trousers  | 4  | $70  |
| Shoes  | 1  | $120  |
| Ties  | 8  | $20  |




This project has  4 Domain layer section :
 * Models Layer 
 * Repository Layer
 * Business logic/Usecase Layer  
 * API/Delivery Layer
 
 
 Architecture
 Ecah model has this below structure(eg: Product, cart, promotion, User)
 ![alt text](https://github.com/karuppaiah/shopping/blob/master/architecture.png)
 

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

At this time, you will have a new data.db created in root directory. 

Site runs at http://127.0.0.1:8080/ping

Postman V2 request: https://github.com/karuppaiah/shopping/blob/master/golang%20shopping.postman_collection

Always get the JWT token and use them in authorization header for response.
```

# Docker

```bash
Run from docker:

Docker image in dockerhub: https://hub.docker.com/r/karups289/kal_shopping/

docker pull karups289/kal_shopping:1.0

docker run -p 8080:8080 --name kalshopping karups289/kal_shopping:1.0

```

# Tasks handled
 * PRODUCT MANAGEMENT - ADD PRODUCT, DELETE PRODUCT . GET ALL PRODUCTS 
 * CART MANAGEMENT - ADD ITEM TO CURRENT USER CART, DELETE ITEM FROM CURRENT USER CART, GET ALL ITEMS OF CURRENT USER CART
    * USER BASED CART OPERATION RESTRICTION
    * VERIFY THE PRODUCT AND STOCK AVAILABILITY BEFORE ADDING OF ITEM TO CART
    * REFRESH CART DISCOUNT AND TOTAL BASED ON CURRENT PROMOTION
    * GET CART ITEMS RETURNS CART ITEMS WITH RESPECTIVE DISCOUNT PRICE AND FINAL TOTAL
 * PROMOTION - ADD PROMOTION, DELETE PROMOTION . GET ALL PROMOTIONS 
    * PROMOTION LOGIC ALLOWED BASED ON FIXED DISCOUNT AND PERCENTAGE DISCOUNT
    * PROMOTION ALLOWED ON MULTI-BUY ON SINGLE PRODUCT
    * VALIDATIONS OF QTY, SOURCE PRODUCT, DESTINATION PRODUCT ARE DONE
 * USER - ADD USER, DELETE USER . GET ALL USER
    * USER MANAGEMENT CAN BE DONE ONLY BY ADMIN USER
 * JWT AUTHENTICATION FOR API'S
    * GET TOKEN
    * REFRESH TOKEN
    * AUHTORIZE ON EVERY API CALL
 * MODEL->REPOSITORY->BUSINESS LOGIC->API'S STRUCTURE FOR EACH SYSTEM
 * UNITTESTING FOR BUSINESS LOGIC AND REPOSITORY LAYER
 * SQLITE DB FOR STORAGE
 * DOCKER IMAGE FOR APPLICATION WITH GOLANG 1.10.3 AND PUSHED TO DOCKERHUB
 * MOCKS WRITTER FOR REPOSITORY AND USECASE
 * CREATE SAMPLE DATA SCRIPT
 * EXECUTION SCRIPT FOR THIS APPLICATION
 * FEW CONFIGS ARE HARDCODED FOR NOW(CAN BE FETCHED FROM ENV. VARIABLES)
 * PERFORMANCE TESTING - 
 ![alt text](https://github.com/karuppaiah/shopping/blob/master/performance_testing.png)
 

# TODO :
- [X] Handle existing promotion re-evaluation based on promotion delete
- [X] User management and link to JWT
- [X] Docker image and publish in hub.docker.com
- [ ] GOlang IOC
- [ ] redis storage for JWT token
- [ ] Move to Mongo DB
- [ ] Write more unit testing for more coverage
- [ ] React frontend
- [ ] Utilites for message Q like kafka(producer and consumer)
- [ ] Tree structure for promotion to enable double source product rule
- [X] Stress testing scripts





