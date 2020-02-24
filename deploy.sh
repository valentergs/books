docker build -t valentergs/books:latest -t valentergs/books:$SHA -f .Dockerfile .
docker push valentergs/books:latest
docker push valentergs/books:$SHA
kubectl apply -f k8s
kubectl set image deployments/books books=valentergs/books:$SHA