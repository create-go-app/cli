<h1 align="center">[WIP] 🕶 Create Go App CLI</h1>
<h3 align="center">Set up a new Go (Golang) full stack app by running one CLI command!</h3>

<img width="100%" src="images/screenshot.jpg" alt="Create Go App screenshot"/>

<br/>

## The Why?

<img align="right" width="196px" src="images/logo_gopher.png" alt="Create Go App logo"/>

WIP

<br/>

## Requirements

- Go `1.11+`
- Go Modules

## Install

First, get `cgapp` CLI package (wait for getting all dependencies, please):

```console
foo@bar:~$ go get github.com/create-go-app/cli/cmd/cgapp
```

Second, install to the `$GOPATH/bin` directory as usual:

```console
foo@bar:~$ go install github.com/create-go-app/cli/cmd/cgapp
```

That's all you need! 🎉

## Usage

```console
foo@bar:~$ cgapp [COMMAND] [OPTION] [value]
```

For example, create new app into `./app` folder with [Echo](https://github.com/labstack/echo) on backend and [Preact](https://github.com/preactjs/preact) as frontend:

```console
foo@bar:~$ cgapp init -b echo -f preact -p ./app
```

## Commands & Options

You can see all available CLI commands (short and full names) by running command with `--help` option.

### `init`

```console
foo@bar:~$ cgapp init --help

NAME:
   cgapp init - init new app

USAGE:
   cgapp init [command options] [arguments...]

OPTIONS:
   --path value, -p value      path to create app, ex. ~/projects/my-app (default: ".")
   --backend value, -b value   backend for your app, ex. net/http, Echo, Gin, Iris (default: "net/http")
   --frontend value, -f value  frontend for your app, ex. (P)React, Vue, Svelte (default: "none")
   --help, -h                  show help (default: false)
```

> **Tip:** if you just run `cgapp init`, it's create into current folder built-in backend with `net/http` package and without any frontend.

### `docker`

```console
foo@bar:~$ cgapp docker --help

NAME:
   cgapp docker - create configured Docker containers

USAGE:
   cgapp docker command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command
   Configured Docker containers:
     nginx  container with Nginx (alpine:latest) and Certbot

OPTIONS:
   --path value, -p value  path to create containers, ex. ~/projects/my-app (default: ".")
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```

## Available production-ready templates

**Golang:**

- (_default_) [create-go-app/net_http-go-template](https://github.com/create-go-app/net_http-go-template) — Backend template with built-in `net/http` ([pkg/net/http](https://golang.org/pkg/net/http/)).
- [WIP][create-go-app/echo-go-template](https://github.com/create-go-app/echo-go-template) — Backend template with Echo ([labstack/echo](https://github.com/labstack/echo)).

**JavaScript:**

- [WIP][create-go-app/react-js-template](https://github.com/create-go-app/react-js-template) — Frontend template with React.js ([facebook/react](https://github.com/facebook/react)).
- [create-go-app/preact-js-template](https://github.com/create-go-app/preact-js-template) — Frontend template with Preact.js ([preactjs/preact](https://github.com/preactjs/preact)).

## Configured production-ready Docker containers

**Web server:**

- [create-go-app/nginx-certbot-docker](https://github.com/create-go-app/nginx-certbot-docker) — Docker container with Nginx and Certbot.

## User templates & containers

Create Go App CLI provide creation your own template, instead of those prepared by authors. It's easy! Just specify `-b` or/and `-f` with address to your repository and run:

```console
foo@bar:~$ cgapp init -b github.com/user1/my-back-template -f github.com/user2/my-front-template
```

> **Please note:** the _https_ protocol will add automatically.

## How to update CLI to latest version?

Similar to install, but add `-u` option:

```console
foo@bar:~$ go get -u github.com/create-go-app/cli/cmd/cgapp
```

And now, install again:

```console
foo@bar:~$ go install github.com/create-go-app/cli/cmd/cgapp
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
| <img align="center" width="100px" src="images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — IT specialists around the world, who are ready to share their experience to solve your business objectives. | [https://1wa.co](https://1wa.co) |
|                                                                                                        | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div>                           |                                  |

## License

MIT
