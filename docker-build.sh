docker build --build-arg SVC=api -t api .
docker build --build-arg SVC=rpc/stock -t stocksvc .
docker build --build-arg SVC=rpc/item -t itemsvc .

docker run -d -p 8889:8889 --name=api_1 api
docker run -d --name=stocksvc_1 stocksvc
docker run -d --name=itemsvc_1 itemsvc