# faas-function

## Build

Assuming a docker account:

```
docker login
````

Build the local docker image (for raspberry):

```
faas-cli publish -f faas-test-api.yml --platforms linux/arm/v7
```

Deploy this on a running OpenFaaS instance - assuming there is already a `OPENFAAS_URL` environment variable:

```
faas-cli deploy -f faas-test-api.yml
```

Invoke the function:

```
faas-cli invoke faas-test-api


```
