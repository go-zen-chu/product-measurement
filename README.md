# product-measurement

[![Actions Status](https://github.com/go-zen-chu/product-measurement/workflows/ci/badge.svg)](https://github.com/go-zen-chu/product-measurement/actions/workflows/ci.yml)
[![Actions Status](https://github.com/go-zen-chu/product-measurement/workflows/push-image/badge.svg)](https://github.com/go-zen-chu/product-measurement/actions/workflows/push-image.yml)

measure, visualize about your product

![](docs/design.drawio.svg)

## Use locally

### Using kind

1. Install [kind](https://kind.sigs.k8s.io/)
2. Run k8s cluster locally
3. Deploy grafana, mysql to kind cluster

    ```bash
    kubectl apply -k ./k8s/base
    ```

## Develop

### JIRA

1. Generate JIRA API Token from <a href="https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/">Manage API tokens for your Atlassian account</a>. FYI: <a href="https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/">Basic auth for REST APIs</a>
2. Use the token for authorizing <a href="https://developer.atlassian.com/cloud/jira/software/rest/intro/">Jira Software Cloud REST API</a>
