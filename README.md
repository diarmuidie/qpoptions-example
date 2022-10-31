# Knative QPOptions (Queue Proxy) Extension Example

This repo contains example code for a [Knative Serving](https://knative.dev/docs/serving/) queue proxy container which runs a custom [QPOption extension](https://knative.dev/docs/serving/queue-extensions/) to add an `x-foo:bar` HTTP header to all responses.

# Install
To build the queue proxy container run:

```bash
./hack/build.sh
```

Then update the `queue-sidecar-image` field in the config-deployment configmap on your Knative cluster to point at the new queue proxy image.

Now any time Knative starts up a new queue proxy it will run with this image.
