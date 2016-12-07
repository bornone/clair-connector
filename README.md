# It is the connector to enable integration clair with private or public docker register


## Build clair image which contains NVD data

```
docker build --rm -t quay.io/coreos/clair:cfc .

docker save -o clair.tar quay.io/coreos/clair:cfc

docker build --rm -t postgres:cfc .

docker save -o postgres.tar postgres:cfc

```
## Install clair by K8s

```
kubectl create secret generic clairsecret --from-file=./config.yaml
kubectl create -f clair-kubernetes.yaml
```

## Delete clair by K8s

```
kubectl delete rc clair clair-postgres
kubectl delete svc clairsvc postgres 
kubectl delete secret clairsecret 
```

