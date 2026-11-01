# kube-pfp

**Kubernetes PortForwardProxy** — крошечный reverse-proxy для доступа к `kubectl port-forward` снаружи.

## Зачем

`kubectl port-forward` по умолчанию слушает только `127.0.0.1`, поэтому до сервиса нельзя достучаться из другого контейнера, пода или машины. `kube-pfp` поднимает HTTP-прокси на `0.0.0.0`, который перенаправляет запросы на локальный порт port-forward'а.

## Запуск

```bash
kube-pfp <target-port> [listen-port]
```

| Аргумент | По умолчанию | Описание |
| --- | --- | --- |
| `target-port` | — (обязателен) | локальный порт, куда смотрит `kubectl port-forward` |
| `listen-port` | `8080` | порт, на котором слушает прокси |

Пример:

```bash
# в одном терминале
kubectl port-forward svc/my-service 12345:80

# в другом
kube-pfp 12345          # прокси на 0.0.0.0:8080 -> 127.0.0.1:12345
kube-pfp 12345 9000     # прокси на 0.0.0.0:9000 -> 127.0.0.1:12345
```

После этого сервис доступен по `http://<host>:8080`.

## Сборка

```bash
go build -o kube-pfp .
```

Релизы собираются через GoReleaser (linux/darwin, amd64/arm64) по тегу `v*.*.*`.

## Как это работает

Вся программа — `main.go`: парсит аргументы, создаёт `httputil.NewSingleHostReverseProxy` на `http://127.0.0.1:<target-port>` и запускает `http.ListenAndServe` на `0.0.0.0:<listen-port>`. Поддерживается только HTTP; WebSocket-трафик проксируется как есть.
