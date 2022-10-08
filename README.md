## What is this?

Handle Kubernetes secret resources. You can
- list secrets
- reveal secret values
- etc.


## Install

```bash
go install github.com/dubs11kt/revs@latest
```


## How to use?

```bash
# Create secret for test command
$ kubectl create secret generic test-secret -n default --from-literal=foo=bar --from-literal=boo=far --from-literal=moo=gar

# Show secrets
$ revs list
default-token-kvx5z
test-secret

# Reveal secret value
$ revs reveal --namespace default --secret test-secret
boo : far
foo : bar
moo : gar

```
