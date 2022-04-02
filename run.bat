docker-compose -f docker-compose.yaml down
docker rmi bareksa-news
docker build -t bareksa-news .
docker-compose -f docker-compose.yaml up
