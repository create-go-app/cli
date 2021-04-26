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

# If you want to reduce binary size, you can run this command with options:
CGO_ENABLED=0 go install -ldflags="-s -w" github.com/create-go-app/cli/cmd/cgapp@latest
```

Also, macOS and GNU/Linux users available way to install via [Homebrew](https://brew.sh/):

```bash
# Tap a new formula:
brew tap create-go-app/cli

# Installation:
brew install create-go-app/cli/cgapp
```

Let's create a new project via **interactive console UI** (or **CUI** for short) into current folder:

```bash
cgapp create
```

Okay, it works! Now, you can create a new pre-configured Ansible inventory file in current folder:

```bash
cgapp generate -p traefik
```

Open the generated file (called `hosts.ini`) and fill in the variables according to your server configuration. Now you are ready to automatically deploy the project on the remote server:

```bash
cgapp deploy
```

That's all you need to know to start! 🎉

### ~ Docker-way to quick start

If you don't want to install Create Go App CLI to your system, you feel free to using our official [Docker image](https://hub.docker.com/r/koddr/cgapp) and run CLI from isolated container:

```bash
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest
```

With this Docker image, you do **not** have to worry about installing tools/CLI of frontend UI libraries/frameworks. Everything is **already included**: `create-react-app`, `preact-cli`, `vue-cli`, `ng-cli` and `degit` (for Svelte and Sapper).

Available commands for [official Docker image](https://create-go.app/detailed-guides/official-docker-image/):

- [x] [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)
- [x] [`generate`](https://create-go.app/detailed-guides/commands-and-options/#generate)

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
    - [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)
    - [`generate`](https://create-go.app/detailed-guides/commands-and-options/#generate)
    - [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy)
  - [Working with the official Docker image](https://create-go.app/detailed-guides/official-docker-image/)
  - [Run project on your local machine](https://create-go.app/detailed-guides/run-on-local/)
  - [Deploy project to a production server](https://create-go.app/detailed-guides/deploy-to-server/)
  - [Make your own project template](https://create-go.app/detailed-guides/make-custom-template/)
  - [Make your own Ansible role](https://create-go.app/detailed-guides/make-custom-ansible-roles/)
- [FAQ](https://create-go.app/faq/)
  - [Automation of a deploy process](https://create-go.app/automation/)
  - [How do I ask the right question?](https://create-go.app/ask-question/#how-do-i-ask-the-right-question)
- [Official logo](https://create-go.app/logo/)

## ⚙️ Commands & Options

### `create`

CLI command to create a new project with the selected configuration.

For create a new project with the interactive console UI, please run `create` command (_without_ any arguments):

```bash
cgapp create
```

- 📺 Preview: https://recordit.co/LTxFQloedn
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#create

### `generate`

CLI command for generate a new pre-configured Ansible inventory file, called `hosts.ini`, for specified proxy server in current folder:

```bash
cgapp generate [OPTION]
```

| Option | Values | Default | Required? |
| --- | --- | --- | --- |
| `-p`,&nbsp;`--proxy` | <br/><ul><li>`traefik` — a basic ACME challenge via Let's Encrypt server;</li><li>`traefik:dns` — more complex ACME challenge via choosen DNS provider, supports challenge to getting SSL certificates for you subdomains;</li></ul> | `traefik` | Yes |

- 📺 Preview: https://recordit.co/yvlnIu8Lyp
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#generate

<details>
<summary>Example inventory file</summary>

<br/>

```ini
[cgapp_project]
127.0.0.1 # CHANGE THIS TO YOUR REMOTE SERVER IP!

[cgapp_project:vars]
# Set Ansible default variables to start playbook:
ansible_user=root       # remote sudo user name
ansible_become=true     # ask become password for remote sudo user
ansible_connection=ssh  # connection to remote server via SSH

# Set Python 3 default path:
ansible_python_interpreter_path=/usr/bin/python3

# Set Docker network name:
docker_network=cgapp_network

# Set your project domain:
project_domain=example.com

# Set directory on your remote server for store project files:
server_dir=/var/www/cgapp

# Set user (owner) name & group name (to create files on server):
system_user=root
system_group=docker
```

</details>

### `deploy`

CLI command for deploy Docker containers with your project via Ansible to a remote server.

> ☝️ You should only run this command from the **root folder** of your project, which created with `create` and `generate` commands! It's a necessary condition for everything to work perfectly.

Make sure that you have [Python 3.8+](https://www.python.org/downloads/) and [Ansible 2.9+](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-specific-operating-systems) installed on your computer. Run `deploy` command:

```bash
cgapp deploy [OPTION]
```

| Option | Values | Default | Required? |
| --- | --- | --- | --- |
| `-K`,&nbsp;`--ask-become-pass` | <br/><ul><li>`boolean` — prompt you to provide the remote user sudo password (standard Ansible `--ask-become-pass` option);</li></ul> | `false` | No |

- 📺 Preview: https://recordit.co/ewjG9dgMPX
- 📖 Docs: https://create-go.app/detailed-guides/commands-and-options/#deploy

## 🤔 Why another CLI?

When we started this project, we asked ourselves this question too and... came to the conclusion, that approximately 8 out of 10 routine operations at the start of a new project and/or the deployment of an existing one **can be automated**. And it would be better to have all the necessary functions inside one CLI. That's why we transferred all our experience to the Create Go App CLI, which we use ourselves!

So, yes, this CLI gives you the ability to prepare everything you need to **start a new project** (as `create-react-app` for the React.js ecosystem does) and **deploy an existing project** to a remote server in configured and fully isolated Docker containers.

## 📝 Production-ready project templates

**Backend:**

- [x] [`net/http`](https://create-go.app/production-templates/net-http-go/) — Backend template with Golang built-in [net/http](https://golang.org/pkg/net/http/) package.
- [x] [`fiber`](https://create-go.app/production-templates/fiber-go/) — Backend template with [Fiber](https://github.com/gofiber/fiber).

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

> ☝️ Frontend part of your project will be generated **using official CLI** from the authors of each frontend UI library/framework (_under the hood_). So, you'll always get the latest version of `React`, `Preact`, `Vue.js`, `Angular`, `Svelte` or `Sapper` for your project from their authors!

## 🚚 Configured Ansible roles

**Web/Proxy server:**

- [x] `traefik` — role for run Docker container with [Traefik Proxy](https://traefik.io/traefik/).
- [x] `nginx` — role for run Docker container with [Nginx](https://nginx.org).

> ☝️ Since Create Go App CLI `v2.0.0`, we're strongly recommend to use **Traefik Proxy** as default proxy server for your projects.

## 👤 My own custom templates?

Create Go App CLI provide works with **your own** custom templates, instead of those prepared by authors. Just specify backend (`-b`) and frontend (`-f`) with addresses to their repositories in `create` command:

```console
cgapp create \
  -b github.com/user1/my-template-1 \
  -f bitbucket.org/user2/my-template-2
```

> ☝️ The `https://` protocol will be added automatically!

## 👵 How to install older version?

You can do it by using a version suffix in `go install` command:

```bash
go install github.com/create-go-app/cli/cmd/cgapp@2.0.0
```

> ☝️ Don't forget to rename binary after installation, according to the version you have installed! This must be done to avoid confusion with the latest version.
>
> For example: `mv $GOPATH/bin/cgapp $GOPATH/bin/cgapp_v2_0_0` and run it by `cgapp_v2_0_0 [COMMAND]`.

Found all available CLI versions on our [pkg.go.dev](https://pkg.go.dev/github.com/create-go-app/cli?tab=versions) page.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%20%26%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%21%20%F0%9F%9A%80%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **$100** and we get $25).

Together, we can make this project **better** every day! 😘

## ⚠️ License

`Create Go App CLI` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/create-go-app/cli/blob/master/LICENSE). Official logo was created by [Vic Shóstak](https://shostak.dev/) and distributed under [Creative Commons](https://creativecommons.org/licenses/by-sa/4.0/) license (CC BY-SA 4.0 International).
