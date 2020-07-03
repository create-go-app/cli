<h1 align="center">
  <img src="https://github.com/create-go-app/cli/blob/master/.github/images/cgapp_logo@2x.png" width="224px"/><br/>
  Create Go App CLI
</h1>
<p align="center">Create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript)<br/>and deploy automation (Ansible, Docker) by running one CLI command!</p>

<p align="center"><a href="https://github.com/create-go-app/cli/releases" target="_blank"><img src="https://img.shields.io/badge/version-v1.2.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/create-go-app/cli/pkg/cgapp?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="https://gocover.io/github.com/create-go-app/cli/pkg/cgapp" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-98%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/cli" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ [Quick start](https://create-go.app/quick-start/)

First of all, [download](https://golang.org/dl/) and install Go. Version `1.11` or higher is required.

Installation is done by using the [`go build`](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) command with `$GOPATH/bin`:

```console
go build -i -o $GOPATH/bin/cgapp github.com/create-go-app/cli
```

Let's create a new project into `./app` folder with [Fiber](https://github.com/gofiber/fiber) as backend and [Nginx](https://nginx.org/) as web server:

```console
cgapp create -p ./app -b fiber -w nginx
```

Okay, it works! Now, you can deploy this project to a remote server or run on your local machine in isolated Docker containers. Go to the root project folder and type command:

```console
cgapp deploy -u john_doe --ask-become-pass
```

That's all you need to start! 🎉

## 📖 [Official Documentation](https://create-go.app/)

Unfortunately, we are unable to include all helpful documentation to the `README` file. That's why, the best way to better explore all the features of the **Create Go App CLI** is to read the [Official Documentation](https://create-go.app/).

> 🔥 We've put together a subject index specifically for you, so you can find any answer you want in seconds!

- [Detailed guides](https://create-go.app/detailed-guides/)
  - [CLI Installation](https://create-go.app/detailed-guides/installation/)
    - [Alternative installations](https://create-go.app/detailed-guides/installation/#alternative-installations)
  - [Understanding CLI commands and options](https://create-go.app/detailed-guides/commands-and-options/)
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

### `create`

CLI command to create a new project with the selected configuration.

```bash
cgapp create -p <PATH> -b <BACKEND> -f <FRONTEND> -w <WEBSERVER> -d <DB>
```

| Option | Argument    | Required? | Description                       | Default value        |
| ------ | ----------- | --------- | --------------------------------- | -------------------- |
| `-p`   | `PATH`      | no        | path to create project            | current folder, `./` |
| `-b`   | `BACKEND`   | no        | backend for your project          | `net/http`           |
| `-f`   | `FRONTEND`  | no        | frontend for your project         |                      |
| `-w`   | `WEBSERVER` | no        | web/proxy server for your project |                      |
| `-d`   | `DB`        | no        | database for your project         |                      |

### `deploy`

CLI command for deploy Docker containers with your project to a remote server or run on your local machine.

```bash
cgapp deploy -p <PLAYBOOK> -u <USER> -s <HOST> -n <NETWORK> [ARGS...]
```

| Option | Argument   | Required? | Description                                    | Default value         |
| ------ | ---------- | --------- | ---------------------------------------------- | --------------------- |
| `-p`   | `PLAYBOOK` | no        | the Ansible playbook name                      | `deploy-playbook.yml` |
| `-u`   | `USER`     | yes       | an username of remote's server or local's user |                       |
| `-s`   | `HOST`     | no        | a host name from your inventory file           | `localhost`           |
| `-n`   | `NETWORK`  | no        | a network name for your Docker containers      | `cgapp_network`       |

If you need to deploy (or run) a project asking for a password for the `USER`, simply add an `--ask-become-pass` argument in `[ARGS...]` section.

## 📝 [Production-ready project templates](https://create-go.app/production-templates/)

**Golang:**

- [x] [`net/http`](https://create-go.app/production-templates/net-http-go/) — Backend template with Golang built-in [net/http](https://golang.org/pkg/net/http/) package.
- [x] [`fiber`](https://create-go.app/production-templates/fiber-go/) — Backend template with [Fiber](https://github.com/gofiber/fiber).
- [ ] [`echo`](https://create-go.app/production-templates/echo-go/) _WIP_ — Backend template with [Echo](https://github.com/labstack/echo).

**JavaScript:**

- [x] [`react-js`](https://create-go.app/production-templates/react-js/) — Frontend template with [React.js](https://github.com/facebook/react).
- [x] [`preact`](https://create-go.app/production-templates/preact-js/) — Frontend template with [Preact](https://github.com/preactjs/preact).

**TypeScript:**

- [x] [`react-ts`](https://create-go.app/production-templates/react-ts/) — Frontend template with [React.js](https://github.com/facebook/react) TypeScript.

## 🐳 [Configured Docker containers](https://create-go.app/docker-containers/)

**Web/Proxy server:**

- [x] [`nginx`](https://create-go.app/docker-containers/nginx/) — Docker container with Nginx.

**Database:**

- [ ] [`postgres`](https://create-go.app/docker-containers/postgres/) _WIP_ — Docker container with PostgreSQL.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `Create Go App CLI`:

- Add a [GitHub Star](https://github.com/create-go-app/cli) to the project.
- Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%20%26%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%21%20%F0%9F%9A%80%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Donate some money to the project's author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).

Together, we can make this project **better** every day! 😘

## ⚠️ License

`Create Go App CLI` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/create-go-app/cli/blob/master/LICENSE). Official logo was created by [Vic Shóstak](https://github.com/koddr) and distributed under [Creative Commons](https://creativecommons.org/licenses/by-sa/4.0/) license (CC BY-SA 4.0 International).
