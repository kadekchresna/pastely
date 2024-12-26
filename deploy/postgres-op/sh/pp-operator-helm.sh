
helm repo add postgres-operator-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator && \
helm install postgres-operator postgres-operator-charts/postgres-operator

helm repo add postgres-operator-ui-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator-ui && \
helm install postgres-operator-ui postgres-operator-ui-charts/postgres-operator-ui -n default
