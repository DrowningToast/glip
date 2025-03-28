docker build -t inventory_backend .

docker-compose -f docker-compose.all.yml up -d

echo "Pushing database schema to USA1 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it usa1_backend bun db:push
echo "\n"

echo "Pushing database schema to USA2 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it usa2_backend bun db:push
echo "\n"

echo "Pushing database schema to USA3 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it usa3_backend bun db:push
echo "\n"

echo "Pushing database schema to EU1 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it eu1_backend bun db:push
echo "\n"

echo "Pushing database schema to EU2 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it eu2_backend bun db:push
echo "\n"

echo "Pushing database schema to EU3 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it eu3_backend bun db:push
echo "\n"

echo "Pushing database schema to APAC1 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it apac1_backend bun db:push
echo "\n"

echo "Pushing database schema to APAC2 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it apac2_backend bun db:push
echo "\n"

echo "Pushing database schema to APAC3 warehouses ..."
docker-compose -f docker-compose.all.yml exec -it apac3_backend bun db:push
echo "\n"
echo "Health check USA1 ..."
curl -X GET http://localhost:8000/health
echo "\n"

echo "Health check USA2 ..."
curl -X GET http://localhost:8001/health
echo "\n"

echo "Health check USA3 ..."
curl -X GET http://localhost:8002/health
echo "\n"

echo "Health check EU1 ..."
curl -X GET http://localhost:8003/health
echo "\n"

echo "Health check EU2 ..."
curl -X GET http://localhost:8004/health
echo "\n"

echo "Health check EU3 ..."
curl -X GET http://localhost:8005/health
echo "\n"

echo "Health check APAC1 ..."
curl -X GET http://localhost:8006/health
echo "\n"

echo "Health check APAC2 ..."
curl -X GET http://localhost:8007/health
echo "\n"

echo "Health check APAC3 ..."
curl -X GET http://localhost:8008/health
echo "\n"
