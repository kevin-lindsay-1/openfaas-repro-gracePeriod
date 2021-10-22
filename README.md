## Reproduction repo for OpenFaaS grace period termination

### non-standard dependencies:
- [helmfile](https://github.com/roboll/helmfile)
- [tilt](tilt.dev)
- installation of openfaas using the `faas-netes/openfaas` chart, with the following settings:
```yaml
# timeouts in this chart should be >2m

# replaces faas-netes with openfaas-operator
operator:
  create: true
  createCRD: true
```

### to start:
1. `tilt up`

### to reproduce:
1. invoke function with default `SLEEP_DURATION` of `2m`
2. delete the pod working on it
3. you should receive a 500 of "Can't reach service for: sleep.openfaas-fn."

you can also verify this by looking at the pod's yaml and noticing the `terminationGracePeriodSeconds` set to the default of `30`.
