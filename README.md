<h1 align="center">🕶 Create Go App CLI</h1>
<h3 align="center">Set up a new Go (Golang) full stack app by running one CLI command!</h3>

<p align="center"><img src="https://img.shields.io/badge/cli_version-0.8.1-blue?style=for-the-badge&logo=none" alt="cli version" />&nbsp;<img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<img src="https://img.shields.io/badge/securego.io-success-success?style=for-the-badge&logo=none" alt="gosec" />&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="lisense" /></p>

[![asciicast](https://asciinema.org/a/0wjeLuiQPjE0fLVLFBrnJAYve.svg)](https://asciinema.org/a/0wjeLuiQPjE0fLVLFBrnJAYve)

## Requirements

- Go `1.11+`
- Go Modules

## Install

First, get `cgapp` CLI package (_wait for getting all dependencies, please_):

```console
go get github.com/create-go-app/cli/cmd/cgapp
```

Second, install to the `$GOPATH/bin` directory, as usual:

```console
go install github.com/create-go-app/cli/cmd/cgapp
```

That's all you need! 🎉

## Usage

```console
cgapp [command] [command options] [arguments...]
```

For example, create new app into `./app` folder with [Fiber](https://github.com/gofiber/fiber) on backend and [Preact](https://github.com/preactjs/preact) as frontend:

```console
cgapp create -p ./app -b fiber -f preact
```

## Commands & Options

You can see all available CLI commands (short and full names) by running command with `--help` option.

### `create`

```console
cgapp create --help

NAME:
   cgapp create - create new Go app

USAGE:
   cgapp create [command options] [arguments...]

OPTIONS:
   --path value, -p value       path to create app, ex. ~/projects/my-app (default: ".")
   --backend value, -b value    backend for your app, ex. Fiber, Echo (default: "net/http")
   --frontend value, -f value   frontend for your app, ex. Preact, React.js, React.ts (default: "none")
   --webserver value, -w value  web/proxy server for your app (default: "nginx")
   --database value, -d value   database for your app, ex. Postgres (default: "none")
   --help, -h                   show all available CLI commands
```

> **Tip:** if you just run `cgapp create`, it's create into current folder built-in backend with `net/http` package and configured Docker container with `Nginx` and `Certbot`. Without any frontend.

## Available production-ready templates

**Golang:**

- `net/http` [create-go-app/net_http-go-template](https://github.com/create-go-app/net_http-go-template) — Backend template with built-in net/http ([pkg/net/http](https://golang.org/pkg/net/http/)).
- `fiber` [create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template) — Backend template with Fiber ([gofiber/fiber](https://github.com/gofiber/fiber)).
- `echo` [WIP] [create-go-app/echo-go-template](https://github.com/create-go-app/echo-go-template) — Backend template with Echo ([labstack/echo](https://github.com/labstack/echo)).

**JavaScript:**

- `react-js` [create-go-app/react-js-template](https://github.com/create-go-app/react-js-template) — Frontend template with React.js ([facebook/react](https://github.com/facebook/react)).
- `react-ts` [create-go-app/react-ts-template](https://github.com/create-go-app/react-ts-template) — Frontend template with React.js TypeScript ([facebook/react](https://github.com/facebook/react)).
- `preact` [create-go-app/preact-js-template](https://github.com/create-go-app/preact-js-template) — Frontend template with Preact ([preactjs/preact](https://github.com/preactjs/preact)).

## Configured production-ready Docker containers

**Web/Proxy server:**

- `nginx` [create-go-app/nginx-certbot-docker](https://github.com/create-go-app/nginx-certbot-docker) — Docker container with Nginx and Certbot.

**Database:**

- `postgres` [WIP] [create-go-app/postgres-docker](https://github.com/create-go-app/postgres-docker) — Docker container with PostgreSQL.

## User templates & containers

Create Go App CLI provide creation your own template, instead of those prepared by authors. It's easy! Just specify `--backend` (`-b`), `--frontend` (`-f`), `--webserver` (`-w`) and `--database` (`-d`) with addresses to your repositories and run:

```console
cgapp create \
             -b github.com/user/my-back-template \
             -f gitlab.com/user/my-front-template \
             -w github.com/user/my-webserver-container-template \
             -d bitbucket.org/user/my-database-container-template
```

> **Please note:** the _https_ protocol will add automatically.

## How to update CLI to latest version?

Similar to install, but add `-u` option:

```console
go get -u github.com/create-go-app/cli/cmd/cgapp
```

And now, install again:

```console
go install github.com/create-go-app/cli/cmd/cgapp
```

## Developers

- Idea and active development by [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Project assistance

If you want to say «thank you» or/and support active development `create-go-app/cli`:

1. Add a GitHub Star to project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Set%20up%20a%20new%20Go%20%28Golang%29%20full%20stack%20app%20by%20running%20one%20CLI%20command%21%26url%3Dhttps%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **$100** and we get $25).
5. Become a sponsor.

Thanks for your support! 😘 Together, we make this project better every day.

### Sponsors

| Logo                                                                                                   | Sponsor description                                                                                                                 | URL                              |
| ------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| <img align="center" width="100px" src=".github/images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — Team who making UX efficiency review, friendly UI design, smart backend microservices, high-quality web apps and many more. | [https://1wa.co](https://1wa.co) |
|                                                                                                        | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div>                           |                                  |

## License

MIT
