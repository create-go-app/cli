<h1 align="center">
  <img src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/>
  Create Go App CLI
</h1>
<p align="center">Create a new production-ready project with <b>backend</b> (Golang), <b>frontend</b> (JavaScript, TypeScript)<br/>and <b>deploy automation</b> (Ansible, Docker) by running one CLI command.<br/><br/>Focus on <b>writing</b> code and <b>thinking</b> of business-logic! The CLI will take care of the rest.</p>

<p align="center"><a href="https://github.com/create-go-app/cli/releases" target="_blank"><img src="https://img.shields.io/badge/version-v1.7.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/create-go-app/cli?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="https://gocover.io/github.com/create-go-app/cli/pkg/cgapp" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-94%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/cli" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ Quick start

First of all, [download](https://golang.org/dl/) and install **Go**. Version `1.16` or higher is required.

Next, download the **latest** version of the Create Go App CLI to your system:

```bash
go get -u github.com/create-go-app/cli
```

Installation is done by using the [`go install`](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies) command and rename installed binary in `$GOPATH/bin`:

```bash
go install -ldflags="-s -w" github.com/create-go-app/cli && mv $GOPATH/bin/cli $GOPATH/bin/cgapp
```

Let's create a new project via **interactive console UI** (or **CUI** for short) into current folder:

```bash
cgapp create
```

Okay, it works! Now, you can run this project on your **local machine** or deploy to a **remote server**. Project works in isolated Docker containers and automates via Ansible playbook:

```bash
cgapp deploy
```

That's all you need to start! 🎉

### ~ Docker-way to quick start

If you don't want to install Create Go App CLI to your system, you feel free to using our official [Docker image](https://hub.docker.com/r/koddr/cgapp) and run CLI from isolated container:

```bash
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest
```

With this Docker image, you do **not** have to worry about installing tools/CLI of frontend UI libraries/frameworks. Everything is **already included**: `create-react-app`, `preact-cli`, `vue-cli`, `ng-cli` and `degit` (for Svelte and Sapper).

Available commands for [official Docker image](https://create-go.app/detailed-guides/official-docker-image/):

- [x] [`init`](https://create-go.app/detailed-guides/commands-and-options/#init)
- [x] [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)

> 🔔 Please note: a [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy) command is currently unavailable in this image.

## 📺 Video screencast

A short video screencast to introduce main features of the Create Go App CLI.

<a align="center" href="https://youtu.be/5-DNZFU9TOQ" target="_blank">
    <img src="https://create-go.app/assets/images/youtube-preview.jpg" alt="youtube preview"/><br/>
    🔗 https://youtu.be/5-DNZFU9TOQ
</a>

## 📖 Official Documentation

Unfortunately, we are unable to include all helpful documentation to the `README` file. That's why, the best way to better explore all the features of the **Create Go App CLI** is to read the [Official Documentation](https://create-go.app/) and explore [Discussions](https://github.com/create-go-app/cli/discussions).

> 🔥 We've put together a subject index specifically for you, so you can find any answer you want in seconds!

- [Detailed guides](https://create-go.app/detailed-guides/)
  - [CLI Installation](https://create-go.app/detailed-guides/installation/)
    - [Alternative installations](https://create-go.app/detailed-guides/installation/#alternative-installations)
  - [Understanding CLI commands and options](https://create-go.app/detailed-guides/commands-and-options/)
    - [`init`](https://create-go.app/detailed-guides/commands-and-options/#init)
    - [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)
    - [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy)
  - [Working with the official Docker image](https://create-go.app/detailed-guides/official-docker-image/)
  - [Run project on your local machine](https://create-go.app/detailed-guides/run-on-local/)
  - [Deploy project to a production server](https://create-go.app/detailed-guides/deploy-to-server/)
  - [Make your own template](https://create-go.app/detailed-guides/make-custom-template/)
  - [Make your own container](https://create-go.app/detailed-guides/make-custom-container/)
- [FAQ](https://create-go.app/faq/)
  - [Automation of a deploy process](https://create-go.app/automation/)
  - [How do I ask the right question?](https://create-go.app/ask-question/#how-do-i-ask-the-right-question)
- [Official logo](https://create-go.app/logo/)

## ⚙️ Commands & Options

### `init`

CLI command for generate a default `.cgapp.yml` config file in current folder:

```bash
cgapp init
```

- 📺 Preview: https://recordit.co/yvlnIu8Lyp
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#init

<details>
<summary>Generated config file</summary>

<br/>

```yaml
# Project config.
project:
  # Backend for your project.
  # (Required)
  # String:
  #   - `net/http`
  #   - `fiber`
  #   - `echo`
  #   - `gin`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - backend: fiber

  # Frontend for your project.
  # (Optional, to skip set to `none`)
  # String:
  #   - `react`
  #     - `react:<template>`
  #   - `preact`
  #     - `preact:<template>`
  #   - `vue`
  #     - `vue:<user/repo>` (for preset from GitHub)
  #     - `vue:<gitlab|bitbucket>:<user/repo>` (for presets from others)
  #   - `angular`
  #   - `svelte`
  #   - `sapper`
  #     - `sapper:<webpack>`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - frontend: svelte

  # Web/Proxy server for your project.
  # (Optional, to skip set to `none`)
  # String: `nginx`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - webserver: nginx

# Automation config.
roles:
  # Ansible roles for deploy your project.
  # (Optional, to skip set to empty or comment)
  # Objects list.
  - deploy:
    # Username of remote's server or local's user.
    # (Required)
    username: root

    # If you need to deploy (or run) a project asking for a password
    # for the user, set `become` to `true`. This is equivalent of
    # `--ask-become-pass`, a standard Ansible argument
    # to ask for a privilege escalation password.
    # (Optional)
    become: true

    # Host name from your inventory file (usually, at /etc/ansible/hosts).
    # (Required)
    host: localhost

    # Name of Docker network
    # (Required)
    network: cgapp_network

    # Filename of Ansible playbook in the root of the Create Go App project.
    # If you want to rename it, do it, but not to change destination of file!
    # (Required)
    playbook: deploy-playbook.yml
```

</details>

### `create`

CLI command to create a new project with the selected configuration.

There's two ways to create a new project:

- [x] With an interactive console UI (or CUI).
- [x] From configuration file (by default, in `$PWD/.cgapp.yml`).

#### Create with the interactive console UI

Run `create` command **without** any arguments:

```bash
cgapp create
```

- 📺 Preview: https://recordit.co/LTxFQloedn
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#create

#### Create from the config file

Run `create` command **with** `--use-config` (or `-c`) argument:

```bash
cgapp create --use-config
```

### `deploy`

CLI command for deploy Docker containers with your project to a remote server.

> ☝️ You should only run this command from the **root folder** of your project, which created with the `cgapp create` command! It's a necessary condition for everything to work perfectly.

There's, also, two ways to deploy your project:

- [x] With an interactive console UI (or CUI).
- [x] From configuration file (by default, in `$PWD/.cgapp.yml`).

#### Deploy with the interactive console UI

Run `deploy` command **without** any arguments:

```bash
cgapp deploy
```

- 📺 Preview: https://recordit.co/ewjG9dgMPX
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#deploy

#### Deploy from the config file

Run `deploy` command **with** `--use-config` (or `-c`) argument:

```bash
cgapp deploy --use-config
```

## 🤔 Why another CLI?

When we started this project, we asked ourselves this question too and... came to the conclusion, that approximately 8 out of 10 routine operations at the start of a new project and/or the deployment of an existing one **can be automated**. And it would be better to have all the necessary functions inside one CLI. That's why we transferred all our experience to the Create Go App CLI, which we use ourselves!

So, yes, this CLI gives you the ability to prepare everything you need to **start a new project** (as `create-react-app` for the React.js ecosystem does) and **deploy an existing project** to a remote server in configured and fully isolated Docker containers.

## 📝 Production-ready project templates

**Backend:**

- [x] [`net/http`](https://create-go.app/production-templates/net-http-go/) — Backend template with Golang built-in [net/http](https://golang.org/pkg/net/http/) package.
- [x] [`fiber`](https://create-go.app/production-templates/fiber-go/) — Backend template with [Fiber](https://github.com/gofiber/fiber).
- [ ] [`echo`](https://create-go.app/production-templates/echo-go/) (_WIP_) — Backend template with [Echo](https://github.com/labstack/echo).
- [ ] [`gin`](https://create-go.app/production-templates/gin-go/) (_WIP_) — Backend template with [Gin](https://github.com/gin-gonic/gin).

**Frontend:**

- [x] `react` — [React](https://reactjs.org/) frontend app.
  - `react:<template>` — CRA generated template for React app.
- [x] `preact` — [Preact](https://preactjs.com/) frontend app.
  - `preact:<template>` — Preact CLI generated template for Preact app.
- [x] `vue` — [Vue.js](https://vuejs.org/) frontend app.
  - `vue:<user/repo>` — Preset for generating Vue.js app from GitHub.
  - `vue:<gitlab|bitbucket>:<user/repo>` — Preset for generating Vue.js app from GitLab/BitBucket/etc.
- [x] `angular` — [Angular](https://angular.io/) frontend app.
- [x] `svelte` — [Svelte](https://svelte.dev/) frontend app.
- [x] `sapper` — [Sapper](https://sapper.svelte.dev/) frontend app for static websites.
  - `sapper:<webpack>` — Preset for generating Sapper with Webpack bundler.

> ☝️ Please note, that since version `v1.3.0`, frontend templates (_in the classical sense_) are **not** supported by the Create Go App CLI. Those templates, that we created ([`react-js`](https://github.com/create-go-app/react-js-template), [`react-ts`](https://github.com/create-go-app/react-ts-template) and [`preact-js`](https://github.com/create-go-app/preact-js-template)), are still available, but only for use as **user's custom templates**.
>
> Now, the frontend part of your project will be generated **using official CLI** from the authors of each frontend UI library/framework (_under the hood_). So, you'll always get the latest version of `React`, `Preact`, `Vue.js`, `Angular`, `Svelte` or `Sapper` for your project from their authors!

## 🐳 Configured Docker containers

**Web/Proxy server:**

- [x] [`nginx`](https://create-go.app/docker-containers/nginx/) — Docker container with [Nginx](https://nginx.org).

## 👤 Custom templates & containers?

Create Go App CLI provide works with **your own** custom templates, instead of those prepared by authors. Just specify backend, frontend and webserver with addresses to repositories in configuration file (`.cgapp.yml`):

```yaml
project:
  - backend: github.com/user1/my-template-1
  - frontend: gitlab.com/user2/my-template-2
  - webserver: bitbucket.org/user3/my-template-3
# ...
```

> ☝️ The `https://` protocol will be added automatically!

## 👵 How to install older version?

You can do it by using a version suffix in `go install` command:

```bash
# With an indication of the exact versions:
go install github.com/create-go-app/cli@1.6.0
```

> ☝️ Don't forget to rename binary after installation!
> 
> For example, according to the version you've installed: `mv $GOPATH/bin/cli $GOPATH/bin/cgapp_v1_6_0`

You can found all available Create Go App CLI versions on our [pkg.go.dev page](https://pkg.go.dev/github.com/create-go-app/cli?tab=versions).

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%20%26%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%21%20%F0%9F%9A%80%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Donate some money to the project's author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).

Together, we can make this project **better** every day! 😘

## ⚠️ License

`Create Go App CLI` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/create-go-app/cli/blob/master/LICENSE). Official logo was created by [Vic Shóstak](https://github.com/koddr) and distributed under [Creative Commons](https://creativecommons.org/licenses/by-sa/4.0/) license (CC BY-SA 4.0 International).
