rm data.db
# get token for auth api's
# curl -X POST -F username=admin -F password=admin http://localhost:8080/login
# use the token value from above response inside these below mentioned authorization
TOKEN=$(curl -X POST -F username=admin -F password=admin http://localhost:8080/login| awk -v  x=1 '{print $x}'| awk -F ":" '{print $7}'| awk -F '"' '{print $2}')


# # GET ALL PRODUCTS
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Belts\", \"stock\": 10, \"price\": 20}" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Shirts\", \"stock\": 5, \"price\": 60 }" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Suits\", \"stock\": 2, \"price\": 300 }" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Trousers\", \"stock\": 4, \"price\": 70 }" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Shoes\", \"stock\": 1, \"price\": 120 }" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Ties\", \"stock\": 8, \"price\": 20 }" http://localhost:8080/auth/products
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"name\": \"Kids Blue Shorts\", \"stock\": 10, \"price\": 9.90 }" http://localhost:8080/auth/products
# # DELETE A PRODUCT
curl -i -X DELETE  -H "Authorization:Bearer $TOKEN" http://localhost:8080/auth/products?id=7

# # GET CART ITEMS
# curl -i -X GET -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" http://127.0.0.1:8080/auth/cart

# # ADD CART ITEM
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"prodid\": 1, \"items\": 5  }" http://localhost:8080/auth/cart
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{  \"prodid\": 2, \"items\": 3 }" http://localhost:8080/auth/cart
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{  \"prodid\": 3,  \"items\": 2 }" http://localhost:8080/auth/cart
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{  \"prodid\": 4,  \"items\": 2 }" http://localhost:8080/auth/cart

# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" -d "{  \"prodid\": 1, \"items\": 5}" http://localhost:8080/auth/cart
# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" -d "{  \"prodid\": 2,  \"items\": 3 }" http://localhost:8080/auth/cart
# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" -d "{ \"prodid\": 3,  \"items\": 2 }" http://localhost:8080/auth/cart

# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" -d "{ \"prodid\": 1, \"items\": 5 }" http://localhost:8080/auth/cart
# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"prodid\": 2, \"items\": 3 }" http://localhost:8080/auth/cart
# curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"prodid\": 3, \"items\": 2 }" http://localhost:8080/auth/cart

# # DELETE cart item
# curl -i -X GET -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" http://127.0.0.1:8080/auth/cart
curl -i -X DELETE  -H "Authorization:Bearer $TOKEN" http://localhost:8080/auth/cart?prodid=4
# curl -i -X GET -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" http://127.0.0.1:8080/auth/cart

# # GET PROMOTION ITEM
curl -i -X GET -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" http://127.0.0.1:8080/auth/promotion
# # ADD PROMOTION ITEM
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"sprodid\": 4, \"sminqty\": 2 , \"dprodid\": 2, \"dminqty\": 0, \"disctype\": \"P\",\"discount\": 15, \"priority\": 1  }" http://localhost:8080/auth/promotion
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"sprodid\": 4, \"sminqty\": 2 , \"dprodid\": 5, \"dminqty\": 0, \"disctype\": \"P\",\"discount\": 15, \"priority\": 2  }" http://localhost:8080/auth/promotion
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"sprodid\": 2, \"sminqty\": 2 , \"dprodid\": 2, \"dminqty\": 2, \"disctype\": \"F\",\"discount\": 15, \"priority\": 3  }" http://localhost:8080/auth/promotion
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"sprodid\": 2, \"sminqty\": 3 , \"dprodid\": 6, \"dminqty\": 0, \"disctype\": \"P\",\"discount\": 50, \"priority\": 4  }" http://localhost:8080/auth/promotion
curl -i -X POST -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN"  -d "{ \"sprodid\": 2, \"sminqty\": 3 , \"dprodid\": 6, \"dminqty\": 0, \"disctype\": \"P\",\"discount\": 50, \"priority\": 5  }" http://localhost:8080/auth/promotion


# # DELETE PROMOTION ITEM
curl -i -X DELETE  -H "Authorization:Bearer $TOKEN" http://localhost:8080/auth/promotion?id=5

curl -i -X GET -H "Content-Type: application/json" -H "Authorization:Bearer $TOKEN" http://127.0.0.1:8080/auth/promotion


