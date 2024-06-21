# Monitoring course

This contains the Spin Go application and Kubernetes resource files for the Monitoring course at Polytech Montpellier. This is the repository of Alexandre Sollier.

The application is running at [wasm.a2.serpentard.dopolytech.fr](http://wasm.a2.serpentard.dopolytech.fr/person).

## Usage

### Backend

You need [Spin](https://developer.fermyon.com/spin/v2/install), [the Go toolchain](https://go.dev/learn/) and [the TinyGo compiler](https://tinygo.org/getting-started/) to run this application. It is available in the [`api/`](./api/) directory.

This application is dependent on a Postgres database running. You can run the Docker Compose stack for this:

```sh
docker compose up -d
```

You can then immediately start the Spin application in watch mode:

```sh
spin watch
```

The server will be running on port 3000.

### Deployment

The Kubernetes deployment files are available in the [`kube/`](./kube/) directory.

First, deploy the Postgres database to your cluster using the [Bitnami postgresql](https://artifacthub.io/packages/helm/bitnami/postgresql) chart:

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami --force-update
helm install postgres -f db/values.yml -n api-monitoring --create-namespace --version 15.5.7 bitnami/postgresql
```

You will also need to deploy the CRDs for the [Gateway API](https://gateway-api.sigs.k8s.io/):

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gatewayclasses.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_gateways.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_grpcroutes.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.0.0/config/crd/experimental/gateway.networking.k8s.io_tlsroutes.yaml
```

You'll also need to install [the SpinKube operator](https://www.spinkube.dev/docs/spin-operator/installation/installing-with-helm/) to your cluster. Apply the
`SpinAppExecutor` to the namespace where you'll deploy the application:

```sh
kubectl apply -f https://github.com/spinkube/spin-operator/releases/download/v0.2.0/spin-operator.shim-executor.yaml -n api-monitoring
```

Finally, you can deploy the application to the cluster:

```sh
kubectl apply -n api-monitoring -f api/
```

## Services

The Kubernetes distribution used is [RKE2](https://docs.rke2.io/) running on two [Rocky Linux 9](https://rockylinux.org/) nodes with SELinux enabled.

### Cilium

The CNI installed in this cluster is [Cilium](https://docs.cilium.io/en/stable/overview/intro/). 
The configuration files are available in the [`kube/cilium/`](./kube/cilium/) directory.

The following configuration has been applied:
- The Envoy DaemonSet is enabled
- The Hubble Relay and UI are enabled
- Support for the Gateway API is enabled
- Replacement for kube-proxy is enabled
- Prometheus metrics are enabled

Because of a conflict with the internal network at Polytech, the pool of IPv4 addresses allocated to the pods are set to `10.42.0.0/16`.

An IPv4 address pool has also been deployed to the cluster for the Cilium Load Balancer to attribute public IPv4 addressed to `LoadBalancer` services (notably, for gateways and ingress controllers).

### Local Path Provisioner

To provide a storage class for PVs, Rancher's [Local Path Provisioner](https://github.com/rancher/local-path-provisioner) was installed in the cluster.
There was no specific configuration applied to this install.

I encountered a problem with this provisioner and SELinux, where the provisioner couldn't create directories in the provisioning path (`/opt/local-path-provisioner`) because of a "Permission Denied" error.

To fix this, I needed to run the following commands on each node:

```sh
mkdir /opt/local-path-provisioner
chcon -Rt container_file_t /opt/local-path-provisioner
```

This creates the provisioning path, and set the correct type of security context (`container_file_t`) to it so it can be used by the provisioner.

### cert-manager

To generate and manage the TLS certificates, [cert-manager](https://cert-manager.io/) was installed in the cluster. Its configuration files are available in [`kube/cert-manager/`](./kube/cert-manager/).

First, I applied a `ClusterIssuer` custom resource as usual, that is used to issue certificates using Let's Encrypt and HTTP01 challenges.
For the HTTP01 challenge, instead of the usual Ingress solver, the Gateway API solver was used with the Gateway deployed for this application, as described [here](https://cert-manager.io/docs/configuration/acme/http01/#configuring-the-http-01-gateway-api-solver).

Finally, I modified the Gateway resource as described [here](https://cert-manager.io/docs/usage/gateway/) to add the `cert-manager.io/cluster-issuer` annotation, 
the HTTP listener for the HTTP01 challenges, and the HTTPS listener that will use the TLS certificate generated by cert-manager.
