# Что же это такое?

В данном репозитории находится провайдер для Crossplane, который позволяет управлять ресурсами Boundary. Создан на основе [Upjet](https://github.com/crossplane/upjet-provider-template).
Опубликован в [Marketplace](https://marketplace.upbound.io/providers/releaseband/provider-boundary).
Вся информация по обновлению, добавлению новых ресурсов и т.д. в репозитории темплейта.
Если коротко, что обычно делается в репе:
 - в папке config правим референсы на нужные ресурсы
 - в makefile обновляем версию провайдера терраформ
 - вмерживаем темплейт для обновления мейкфайлов, докерфайлов и т.д.
 - запускаем генерацию кода
 - делаем ветку с версией провайдера и пушим в нее изменения
 - собирается образ и пушится в репозиторий


# Provider boundary

`provider-boundary` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
boundary API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/releaseband/provider-boundary):

```
up ctp provider install releaseband/provider-boundary:v0.1.0
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-boundary
spec:
  package: releaseband/provider-boundary:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/releaseband/crossplane-provider-boundary).

## Developing

Run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/releaseband/crossplane-provider-boundary/issues).
