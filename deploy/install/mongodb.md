# REF: https://github.com/helm/charts/tree/master/stable/mongodb

# MongoDB
helm install --name codefly-mongodb stable/mongodb

# Get Password
kubectl get secret --namespace default ${secret_name} -o jsonpath="{.data.mongodb-root-password}" | base64 --decode
