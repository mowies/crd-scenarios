# crd-scenarios

Scenarios of what happens to CRDs in a Helm chart during management of the chart lifecycle.

This repo is set up in different tag releases that simulate different scenarios.
A summary of the versions can be seen below:

| API Versions | Chart Version | Operator Version | Features                                                                           |
|--------------|---------------|------------------|------------------------------------------------------------------------------------|
|              |               |                  | Experiments with Books API                                                         |
| Box v1       | 0.5.0         | v5               | Add Box v1 API version                                                             |
| Box v1, v2   | 0.6.0         | v6               | Add Box v2 API version and set as storage version                                  |
| Box v1, v2   | 0.7.0         | v7               | Set Box v1 as deprecated                                                           |
| Box v1, v2   | 0.8.0         | v8               | Set Box v1 as unserved and deprecated, update controller fetched API version to v2 |
| Box v1, v2   | 0.9.0         | v9               | Add new optional field to Box v2                                                   |
| Box v1, v2   | 0.10.0        | v10              | Remove Box v1 API version                                                          |
v10 is still todo for next week
