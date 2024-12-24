

cd deploy/postgres-op

kubectl create -f configmap.yaml  # configuration
kubectl create -f operator-service-account-rbac.yaml  # identity and permissions
kubectl create -f postgres-operator.yaml  # deployment
kubectl create -f api-service.yaml  # operator API to be used by UI

# add repo for postgres-operator-ui
helm repo add postgres-operator-ui-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator-ui && \
helm install postgres-operator-ui postgres-operator-ui-charts/postgres-operator-ui

# kubectl port-forward service/postgres-pastely-production-pooler 15432:5432
# kubectl port-forward service/postgres-pastely-production-pooler-repl 25432:5432