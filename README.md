# reveals

## What is this?

Reveal Kubernetes secret resources values

## Install

```bash
$ go install github.com/dubs11kt/reveals@latest
```

## How to use?

```bash
# Create secret to test command
$ kubectl create secret generic test-secret -n default --from-literal=foo=bar --from-literal=boo=far --from-literal=moo=gar

# Exec
$ reveals --namespace default --secret test-secret
boo : far
foo : bar
moo : gar
```

