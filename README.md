<h1 align="center">
  <img src="https://github.com/create-go-app/cli/blob/master/.github/images/cgapp_logo@2x.png" width="224px"/><br/>
  Create Go App CLI
</h1>
<p align="center">Create a new production-ready project with <b>backend</b> (Golang), <b>frontend</b> (JavaScript, TypeScript)<br/>and <b>deploy automation</b> (Ansible, Docker) by running one CLI command.<br/><br/>Focus on <b>writing</b> code and <b>thinking</b> of business-logic! The CLI will take care of the rest.</p>

<p align="center"><a href="https://github.com/create-go-app/cli/releases" target="_blank"><img src="https://img.shields.io/badge/version-v1.4.1-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/create-go-app/cli?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="https://gocover.io/github.com/create-go-app/cli/pkg/cgapp" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-94%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/cli" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ [Quick start](https://create-go.app/quick-start/)

First of all, [download](https://golang.org/dl/) and install **Go**. Version `1.11` or higher is required.

Installation is done by using the [`go build`](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) command with `$GOPATH/bin`:

```bash
go build -i -o $GOPATH/bin/cgapp github.com/create-go-app/cli
```

Let's create a new project via **interactive console UI** (or **CUI**) into current folder:

```bash
cgapp create
```

Okay, it works! Now, you can run this project on your **local machine** or deploy to a **remote server**. Project works in isolated Docker containers and automates via Ansible playbook:

```bash
cgapp deploy
```

### ~ Docker-way to quick start

If you don't want to install Create Go App CLI to your system, you feel free to using our official [Docker image](https://hub.docker.com/r/koddr/cgapp) and run CLI from isolated container:

```bash
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest
```

> ☝️ Also, with this Docker image, you do **not** have to worry about installing tools/CLI of frontend UI libraries/frameworks. Everything is **already included** to this Docker image: `create-react-app`, `preact-cli`, `vue-cli`, `ng-cli` and `degit` (for Svelte).

That's all you need to start! 🎉

## 📖 [Official Documentation](https://create-go.app/)

Unfortunately, we are unable to include all helpful documentation to the `README` file. That's why, the best way to better explore all the features of the **Create Go App CLI** is to read the [Official Documentation](https://create-go.app/).

> 🔥 We've put together a subject index specifically for you, so you can find any answer you want in seconds!

- [Detailed guides](https://create-go.app/detailed-guides/)
  - [CLI Installation](https://create-go.app/detailed-guides/installation/)
    - [Alternative installations](https://create-go.app/detailed-guides/installation/#alternative-installations)
  - [Understanding CLI commands and options](https://create-go.app/detailed-guides/commands-and-options/)
    - [`init`](https://create-go.app/detailed-guides/commands-and-options/#init)
    - [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)
    - [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy)
  - [Run project on your local machine](https://create-go.app/detailed-guides/run-on-local/)
  - [Deploy project to a production server](https://create-go.app/detailed-guides/deploy-to-server/)
  - [Make your own template](https://create-go.app/detailed-guides/make-custom-template/)
  - [Make your own container](https://create-go.app/detailed-guides/make-custom-container/)
- [FAQ](https://create-go.app/faq/)
  - [Automation of a deploy process](https://create-go.app/automation/)
  - [How do I ask the right question?](https://create-go.app/ask-question/#how-do-i-ask-the-right-question)
- [Official logo](https://create-go.app/logo/)

## ⚙️ [Commands & Options](https://create-go.app/detailed-guides/commands-and-options/)

### [`init`](https://create-go.app/detailed-guides/commands-and-options/#init)

CLI command for generate a default `.cgapp.yml` config file in current folder:

```bash
cgapp init
```

<details>
<summary>Generated config file</summary>

<br/>

```yaml
# Project config.
project:
  # Backend for your project.
  # (Required)
  # String: `net/http`, `fiber`, `echo`, `gin`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - backend: fiber

  # Frontend for your project.
  # (Optional, to skip set to `none`)
  # String:
  #   - `react`, `react:<template>`
  #   - `preact`, `preact:<template>`
  #   - `vue`
  #   - `svelte`
  #   - `angular`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - frontend: svelte

  # Web/Proxy server for your project.
  # (Optional, to skip set to `none`)
  # String: `nginx`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - webserver: nginx

  # Web/proxy server for your project.
  # (Optional, to skip set to `none`)
  # String: `postgres`
  # User template: supported, set to URL (without protocol),
  # like `github.com/user/template`
  - database: postgres

# Automation config.
roles:
  # Ansible roles for deploy your project.
  # (Optional, to skip set to empty or comment)
  # Objects list: `deploy`
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

### [`create`](https://create-go.app/detailed-guides/commands-and-options/#create)

CLI command to create a new project with the selected configuration.

There's two ways to create a new project:

- [x] With an interactive console UI (or CUI).
- [x] From configuration file (by default, in `$PWD/.cgapp.yml`).

#### Create with the interactive console UI

Run `create` command **without** any arguments:

```bash
cgapp create
```

#### Create from the config file

Run `create` command **with** `--use-config` (or `-c`) argument:

```bash
cgapp create --use-config
```

### [`deploy`](https://create-go.app/detailed-guides/commands-and-options/#deploy)

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

#### Deploy from the config file

Run `deploy` command **with** `--use-config` (or `-c`) argument:

```bash
cgapp deploy --use-config
```

## 📝 [Production-ready project templates](https://create-go.app/production-templates/)

**Backend:**

- [x] [`net/http`](https://create-go.app/production-templates/net-http-go/) — Backend template with Golang built-in [net/http](https://golang.org/pkg/net/http/) package.
- [x] [`fiber`](https://create-go.app/production-templates/fiber-go/) — Backend template with [Fiber](https://github.com/gofiber/fiber).
- [ ] [`echo`](https://create-go.app/production-templates/echo-go/) _WIP_ — Backend template with [Echo](https://github.com/labstack/echo).
- [ ] [`gin`](https://create-go.app/production-templates/gin-go/) _WIP_ — Backend template with [Gin](https://github.com/gin-gonic/gin).

**Frontend:**

- [x] [`react`](https://create-go.app/production-templates/react-js/) — Frontend app with [React.js](https://reactjs.org).
- [x] [`preact`](https://create-go.app/production-templates/preact-js/) — Frontend app with [Preact](https://preactjs.com).
- [ ] [`vue`](https://create-go.app/production-templates/vue-js/) _WIP_ — Frontend app with [Vue.js](https://vuejs.org).
- [x] [`svelte`](https://create-go.app/production-templates/svelte-js/) — Frontend app with [Svelte](https://svelte.dev).
- [ ] [`angular`](https://create-go.app/production-templates/angular-js/) _WIP_ — Frontend app with [Angular](https://angular.io).

## 🐳 [Configured Docker containers](https://create-go.app/docker-containers/)

**Web/Proxy server:**

- [x] [`nginx`](https://create-go.app/docker-containers/nginx/) — Docker container with [Nginx](https://nginx.org).

**Database:**

- [ ] [`postgres`](https://create-go.app/docker-containers/postgres/) _WIP_ — Docker container with [PostgreSQL](https://postgresql.org).

## 🤔 Why another CLI?

Yes, when we started this project, we asked ourselves this question too and... came to the conclusion, that about **8-10** routine steps in each project can be automated with a smart CLI.

The Create Go App project allow you to prepare and deploy your project **without** any unnecessary headaches.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%20%26%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%21%20%F0%9F%9A%80%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Donate some money to the project's author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).

Together, we can make this project **better** every day! 😘

## ⚠️ License

`Create Go App CLI` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/create-go-app/cli/blob/master/LICENSE). Official logo was created by [Vic Shóstak](https://github.com/koddr) and distributed under [Creative Commons](https://creativecommons.org/licenses/by-sa/4.0/) license (CC BY-SA 4.0 International).
