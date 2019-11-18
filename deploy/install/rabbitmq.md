# REF: https://github.com/helm/charts/tree/master/stable/rabbitmq

# RabbitMQ
helm install --name codefly-rabbitmq --set rabbitmq.username=admin stable/rabbitmq

# GET Password & Erlang Cookie
kubectl get secret --namespace default ${secret_name} -o jsonpath="{.data.rabbitmq-password}" | base64 --decode
kubectl get secret --namespace default ${secret_name} -o jsonpath="{.data.rabbitmq-erlang-cookie}" | base64 --decode
