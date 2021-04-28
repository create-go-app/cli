<h1 align="center">
  <img src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/>
  Create Go App CLI
</h1>
<p align="center">Create a new production-ready project with <b>backend</b> (Golang), <b>frontend</b> (JavaScript, TypeScript)<br/>and <b>deploy automation</b> (Ansible, Docker) by running one CLI command.<br/><br/>Focus on <b>writing</b> code and <b>thinking</b> of business-logic! The CLI will take care of the rest.</p>

<p align="center"><a href="https://github.com/create-go-app/cli/releases" target="_blank"><img src="https://img.shields.io/badge/version-v2.0.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/create-go-app/cli?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="https://gocover.io/github.com/create-go-app/cli/pkg/cgapp" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-94%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/cli" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ Quick start

First of all, [download](https://golang.org/dl/) and install **Go**. Version `1.16` or higher is required.

Installation is done by using the [`go install`](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies) command and rename installed binary in `$GOPATH/bin`:

```bash
go install github.com/create-go-app/cli/cmd/cgapp@latest
```

Also, macOS and GNU/Linux users available way to install via [Homebrew](https://brew.sh/):

```bash
# Tap a new formula:
brew tap create-go-app/cli

# Installation:
brew install create-go-app/cli/cgapp
```

Let's create a new project via **interactive console UI** (or **CUI** for short) in current folder:

```bash
cgapp create
```

Next, open the generated Ansible inventory file (called `hosts.ini`) and fill in the variables according to your server configuration. And you're ready to **automatically deploy** this project:

```bash
cgapp deploy
```

That's all you need to know to start! 🎉

### 🐳 Docker-way to quick start

If you don't want to install Create Go App CLI to your system, you feel free to using our official [Docker image](https://hub.docker.com/r/koddr/cgapp) and run CLI from isolated container:

```bash
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest [COMMAND]
```

> 🔔 Please note: the [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy) command is currently **unavailable** in this image.

## 📺 Video screencast

A short video screencast to introduce main features of the Create Go App CLI.

<a align="center" href="https://youtu.be/5-DNZFU9TOQ" target="_blank">
  <img src="https://create-go.app/assets/images/youtube-preview.jpg" alt="youtube preview"/><br/>
  🔗 https://youtu.be/5-DNZFU9TOQ
</a>

## 📖 Official Documentation

The best way to better explore all the features of the **Create Go App CLI** is to read the [Official Documentation](https://create-go.app/) and take part in [Discussions](https://github.com/create-go-app/cli/discussions). We've put together a subject index specifically for you, so you can find any answer you want in seconds!

- [Detailed guides](https://create-go.app/detailed-guides/)
  - [CLI Installation](https://create-go.app/detailed-guides/installation/)
    - [Alternative installations](https://create-go.app/detailed-guides/installation/#alternative-installations)
  - [Understanding CLI commands and options](https://create-go.app/detailed-guides/commands-and-options/)
    - [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)
    - [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy)
  - [Working with the official Docker image](https://create-go.app/detailed-guides/official-docker-image/)
  - [Run project on your local machine](https://create-go.app/detailed-guides/run-on-local/)
  - [Deploy project to a production server](https://create-go.app/detailed-guides/deploy-to-server/)
- [Automation of the deploy process](https://create-go.app/automation/)
- [How do I ask the right question?](https://create-go.app/ask-question/#how-do-i-ask-the-right-question)
- [Official logo](https://create-go.app/logo/)

## ⚙️ Commands & Options

### `create`

CLI command for create a new project with the interactive console UI.

```bash
cgapp create
```

- 📺 Preview: https://recordit.co/LTxFQloedn
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#create

### `deploy`

CLI command for deploy Docker containers with your project via Ansible to the remote server.

> ☝️ Make sure that you have [Python 3.9+](https://www.python.org/downloads/) and [Ansible 2.10+](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-specific-operating-systems) installed on your computer.

```bash
cgapp deploy [OPTION]
```

| Option | Values                                                                                                            | Default | Required? |
| ------ | ----------------------------------------------------------------------------------------------------------------- | ------- | --------- |
| `-K`   | `boolean` — prompt you to provide the remote user sudo password (_a standard Ansible `--ask-become-pass` option_) | `false` | No        |

- 📺 Preview: https://recordit.co/ewjG9dgMPX
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#deploy

## 📝 Production-ready project templates

### Backend

- Backend template with Golang built-in [net/http](https://golang.org/pkg/net/http/) package:
  - [`net/http`](https://github.com/create-go-app/net_http-go-template) — simple REST API with CRUD and JWT auth.
- Backend template with [Fiber](https://github.com/gofiber/fiber):
  - [`fiber`](https://github.com/create-go-app/fiber-go-template) — complex REST API with CRUD, JWT auth and Renew token.

### Frontend

- Pure JavaScript:
  - `vanilla` — generated template with pure JavaScript app.
  - `vanilla-ts` — generated template with pure TypeScript app.
- [React](https://reactjs.org/):
  - `react` — generated template with React app.
  - `react-ts` — generated template with TypeScript for React app.
- [Preact](https://preactjs.com/):
  - `preact` — generated template with Preact app.
  - `preact-ts` — generated template with TypeScript for Preact app.
- [Vue.js](https://vuejs.org/):
  - `vue` — generated template with Vue.js app.
  - `vue-ts` — generated template with TypeScript for Vue.js app.
- [Svelte](https://svelte.dev/):
  - `svelte` — generated template with Svelte app.
  - `svelte-ts` — generated template with TypeScript for Svelte app.
- [Lit](https://lit.dev/) web components:
  - `lit-element` — generated template with Lit app.
  - `lit-element-ts` — generated template with TypeScript for Lit app.

> ☝️ Frontend part will be generate using awesome tool [Vite.js](https://vitejs.dev/) under the hood. So, you'll always get the latest version of `React`, `Preact`, `Vue`, `Svelte`, `Lit` or pure JavaScript/TypeScript templates for your project!

## 🚚 Pre-configured Ansible roles

### Web/Proxy server

- Roles for run Docker container with [Traefik Proxy](https://traefik.io/traefik/):
  - `traefik` — configured Traefik container with simple ACME challenge via CA server.
  - `traefik-acme-dns` — configured Traefik container with complex ACME challenge via DNS provider.
- Roles for run Docker container with [Nginx](https://nginx.org):
  - `nginx` — pure Nginx container with "the best practice" configuration.

> 👍 Since Create Go App CLI `v2.0.0`, we're strongly recommend to use **Traefik Proxy** as default proxy server for your projects. The main reason: this proxy provides _automatic_ SSL certificates from Let's Encrypt out of the box. Also, Traefik was built on the Docker ecosystem and has a _really good looking_ and _useful_ Web UI.

## 📚 FAQ

### Why another CLI?

When we started this project, we asked ourselves this question too and... came to the conclusion, that approximately 8 out of 10 routine operations at the start of a new project and/or the deployment of an existing one **can be automated**. And it would be better to have all the necessary functions inside one CLI. That's why we transferred all our experience to the Create Go App CLI, which we use ourselves!

So, yes, this CLI gives you the ability to prepare everything you need to **start a new project** (as `create-react-app` for the React.js ecosystem does) and **deploy an existing project** to a remote server in configured and fully isolated Docker containers.

### How to reduce binary size of the CLI?

```console
CGO_ENABLED=0 go install -ldflags="-s -w" github.com/create-go-app/cli/cmd/cgapp@latest
```

> ☝️ By the way, if you install the CLI by Homebrew, you already have this optimization.

### How to install older version?

```console
go install github.com/create-go-app/cli/cmd/cgapp@v2.0.0
```

Found all available CLI versions on our [pkg.go.dev](https://pkg.go.dev/github.com/create-go-app/cli?tab=versions) page.

> ☝️ Don't forget to rename binary after installation, according to the version you have installed! This must be done to avoid confusion with the latest version.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%20%26%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%21%20%F0%9F%9A%80%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **$100** and we get $25).

Together, we can make this project **better** every day! 😘

## ⚠️ License

`Create Go App CLI` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/create-go-app/cli/blob/master/LICENSE). Official logo was created by [Vic Shóstak](https://shostak.dev/) and distributed under [Creative Commons](https://creativecommons.org/licenses/by-sa/4.0/) license (CC BY-SA 4.0 International).
