# reveal-secrets

## What is this?

Reveal Kubernetes secret resources values

## How to use?

```
# Create secret to test command
$ kubectl create secret generic test-secret -n default --from-literal=foo=bar --from-literal=boo=far --from-literal=moo=gar

# Exec
$ ./reveal-secret-values --namespace default --secret test-secret
boo : far
foo : bar
moo : gar

```

