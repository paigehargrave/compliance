# Docker EE InSpec Profiles

[InSpec](https://inspec.io), by Chef Software, is an open source command-line tool that can be used to audit many types of infrastructures against pre-defined security controls and benchmarks.

Profiles are used by the InSpec tool to scan an active instance of Docker EE, Universal Control Plane and Docker Trusted Registry and ensure that all of the components have been configured to meet applicable security requirements and baselines. We've included the following InSpec profiles for Docker EE:

- FedRAMP Moderate
- FedRAMP High

## Using the InSpec tool

You can download the InSpec tool at https://www.inspec.io/downloads/. If you prefer, you can also use the official Docker image ([chef/inspec](https://store.docker.com/community/images/chef/inspec)) to execute an audit. Refer to the [InSpec documentation](https://www.inspec.io/docs/) for full CLI usage instructions.

Before you begin, you need to create a `profile-attribute.yml` file which contains your UCP and DTR login information. You can use the `profile-attribute.example.yml` file as an example.

You can then use the InSpec commands below to audit your cluster at a chosen baseline:

1. Set correct directory

    ```sh
    cd validation/inspec
    ```

2. Audit cluster at FedRAMP Moderate baseline

    ```sh
    inspec exec FedRAMP/Moderate --attrs profile-attribute.yml
    ```

We also maintain a Docker image that already includes our InSpec profiles and the InSpec CLI. If you prefer, you can use it as follows:

```sh
docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD":/share docker/compliance-inspec exec FedRAMP/Moderate --attrs profile-attribute.yml
```
