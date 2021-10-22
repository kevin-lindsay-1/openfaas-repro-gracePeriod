## Reproduction repo for OpenFaaS grace period termination

dependencies:
- [helmfile](https://github.com/roboll/helmfile)
- [tilt](tilt.dev)

to start:
1. `tilt up`

to reproduce:
1. invoke function with default `SLEEP_DURATION` of `2m`
2. delete the pod working on it
3. you should receive a 500 of "Can't reach service for: sleep.openfaas-fn."

you can also verify this by looking at the pod's yaml and noticing the `terminationGracePeriodSeconds` set to the default of `30`.
