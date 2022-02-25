# choppy

## What is this?

Handle Kubernetes secret resources


## Install

```bash
$ go install github.com/dubs11kt/choppy@latest
```


## How to use?

```bash
# Create secret to test command
$ kubectl create secret generic test-secret -n default --from-literal=foo=bar --from-literal=boo=far --from-literal=moo=gar

# Exec
$ choppy --namespace default --secret test-secret
boo : far
foo : bar
moo : gar
```
