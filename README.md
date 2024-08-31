# Taxsi

```
version: '3'

services:
  reverse-proxy:
    image: traefik:v3.1
    # Enables the web UI and tells Traefik to listen to docker
    command: --api.insecure=true --providers.docker
    ports:
      # The HTTP port
      - "8000:8000"
      # The Web UI (enabled by --api.insecure=true)
      - "8080:8080"
    volumes:
      - ./example/traefik.yaml:/etc/traefik/traefik.yaml
      - ./example/dynamic.yaml:/etc/traefik/dynamic.yaml
      - .:/plugins-local/src/github.com/nzin/taxsi/
```