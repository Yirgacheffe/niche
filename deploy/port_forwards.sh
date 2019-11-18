# RabbitMQ
kubectl port-forward -n default svc/codefly-rabbitmq 15672:15672

# Greetings Service D
kubectl port-forward -n niche-dev svc/niche-greetings-d 8090:8080

# MongoDB
kubectl port-forward --namespace default svc/codefly-mongodb 27017:27017 &
mongo --host 127.0.0.1 --authenticationDatabase admin -p $MONGODB_ROOT_PASSWORD
