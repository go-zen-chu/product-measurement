# product-measurement

[![Actions Status](https://github.com/go-zen-chu/product-measurement/workflows/ci/badge.svg)](https://github.com/go-zen-chu/product-measurement/actions/workflows/ci.yml)
[![Actions Status](https://github.com/go-zen-chu/product-measurement/workflows/push-image/badge.svg)](https://github.com/go-zen-chu/product-measurement/actions/workflows/push-image.yml)

Measure, visualize all about your products.

## Design

![](docs/design.drawio.svg)

## Install

### Deploying to your local k8s cluster

1. Install [kind](https://kind.sigs.k8s.io/) or minikube that runs k8s cluster locally
1. Deploy grafana, mysql to your local k8s cluster

    ```bash
    kubectl apply -k ./k8s/base
    kubectl port-forward -nproduct-measurement svc/mysql 3306:3306
    ```

### Deploy to production k8s cluster

As above, deploy to your k8s cluster.
If you want to store your data in external DB, you can edit your k8s manifests.

## Usage

### Import data from datasources

First, you can test locally using [example-config.yaml](./example-config.yaml).

```bash
go run ./cmd/importer -config-path example-config.yaml -filter "sample spreadsheet"
```

If you want to try importing from API, copy and rename example-config.yaml and edit `jira` config.

```bash
cp ./example-config.yaml ./your_config.yaml
# import data from datasource and store to DB
./importer -config-path your_config.yaml 
# import data specifying datasource
./importer -config-path your_config.yaml -filter "sample project"
```

### Visualize data in Grafana

Imported data is stored in DB (MySQL is only supported yet).
You can visualize your data through SQL queries defined in your Grafana dashboards.

-> [example dashboards](./dashboards/)

```bash
kubectl port-forward -nproduct-measurement svc/grafana 3000:3000

# open your browser to see grafana
open "http://localhost:3000"
```

## How to develop this project?

### JIRA

1. Generate JIRA API Token from <a href="https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/">Manage API tokens for your Atlassian account</a>. FYI: <a href="https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/">Basic auth for REST APIs</a>
2. Use the token for authorizing <a href="https://developer.atlassian.com/cloud/jira/software/rest/intro/">Jira Software Cloud REST API</a>
