# Shipreporting Platform [WIP]
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2141ec55b93b4431b9399ed118e7d20c)](https://www.codacy.com/manual/deeper-x/shipreporting-platform?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deeper-x/shipreporting-platform&amp;utm_campaign=Badge_Grade) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/deeper-x/shipreporting-platform) [![Go Report Card](https://goreportcard.com/badge/github.com/deeper-x/shipreporting-platform)](https://goreportcard.com/report/github.com/deeper-x/shipreporting-platform) ![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/deeper-x/shipreporting-platform?include_prereleases)
[![CircleCI](https://circleci.com/gh/deeper-x/shipreporting-platform.svg?style=svg)](https://circleci.com/gh/deeper-x/shipreporting-platform)
![Actions Badge](https://github.com/deeper-x/shipreporting-platform/workflows/Go/badge.svg)

## USER TOKEN GENERATION (Application server)

```bash
cd <shipflow_home>
pipenv shell
python manage drf_create_token <USERNAME>
```

## ACTIVE SERVICES

### LIVE DATA

- Moored
- Anchored
- Arrivals today
- Departures today
- Arrivals previsions
- Shipped goods
- Traffic list
- Shifting previsions

### Authentication
System read auth data from remote ws system, then build a session token 

![Auth](https://raw.githubusercontent.com/deeper-x/shipreporting-platform/master/project/auth_system.png)


### Systemd configuration

```bash
$ sudo cat > /etc/systemd/system/shipreporting-platform.service <<HEREDOC 
[Unit]
Description=Shipreporting service middleware
Documentation=https://github.com/deeper-x/shipreporting-platform
After=network.target

[Service]
Type=simple
User=shipflow
Environment=GOPATH=/home/shipflow/go
WorkingDirectory=/home/shipflow/go/bin/
ExecStart=/home/shipflow/go/bin/shipreporting-platform
Restart=on-failure

[Install]
WantedBy=multi-user.target

HEREDOC
  ```

### Install & Run instance

```bash
go get -d ./...
go build -o ${GOBIN}
sudo systemctl restart shipreporting-platform
```
