# Dekeris Operator

## Creating

```
operator-sdk init --domain dekeris.github.io --license 'none' --owner 'Alessio G. Baroni' --project-name dekeris --repo github.com/dekeris/operator
```

Personalizza il Makefile

```
operator-sdk create api --controller --group games --kind Dekeris --version v1alpha1 --resource
```

```
make bundle bundle-build bundle-push catalog-build catalog-push
```
