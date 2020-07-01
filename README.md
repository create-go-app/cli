<h1 align="center">
  <img src="https://github.com/create-go-app/cli/blob/master/.github/images/cgapp_logo.svg?v2" width="224px"/><br/>
  Create Go App CLI
</h1>
<p align="center">Create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript)<br/>and deploy automation (Ansible, Docker) by running one CLI command!</p>

<p align="center"><a href="https://github.com/create-go-app/cli/releases" target="_blank"><img src="https://img.shields.io/badge/version-v1.2.0-blue?style=for-the-badge&logo=none" alt="cli version" /></a>&nbsp;<a href="https://pkg.go.dev/github.com/create-go-app/cli/pkg/cgapp?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<a href="https://gocover.io/github.com/create-go-app/cli/pkg/cgapp" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-98%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/cli" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡️ Quick start

Let's create a new project into `./app` folder with [Fiber](https://github.com/gofiber/fiber) as backend and [Nginx](https://nginx.org/) as web server:

```console
cgapp create -p ./app -b fiber -w nginx
```

Okay, it works! Now, you can deploy this project to a remote server or run on your local machine in isolated Docker containers (one for `Fiber` and one for `Nginx`):

```console
cd ./app
cgapp deploy -u john_doe --ask-become-pass
```

That's all you need to start! 🎉

## ⚙️ Installation

First of all, [download](https://golang.org/dl/) and install Go. Version `1.11` or higher is required.

Installation is done by using the [`go build`](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) command with `$GOPATH/bin`:

```console
go build -i -o $GOPATH/bin/cgapp github.com/create-go-app/cli
```

Check, that the CLI is installed correctly by the `--version` (or `-v`) command:

```console
cgapp --version
```

### ~ Alternative installations

We're using a [GoReleaser](https://github.com/goreleaser/goreleaser) project for shipping _standalone_ **Create Go App CLI** version to all major desktop platforms: _Apple macOS_, _GNU/Linux_, _MS Windows_. By default, for _amd64_ (x86_64) architecture.

If you need this version, please go to the repository [release page](https://github.com/create-go-app/cli/releases) and download zipped archive.

## 📚 Commands & Options

```console
cgapp [command] [command options] [arguments...]
```

> ☝️ Tip: you can see all available commands by running command with `--help` (or `-h`) option.

### `create`

CLI command to create a new project with the selected configuration.

#### Usage

```console
cgapp create [command options] [arguments...]
```

#### Options

```console
--path value, -p value       path to create app, ex. ~/projects/my-app (default: ".")
--backend value, -b value    backend for your app, ex. Fiber, Echo (default: "net/http")
--frontend value, -f value   frontend for your app, ex. Preact, React.js, React.ts (default: "none")
--webserver value, -w value  web/proxy server for your app, ex. Nginx (default: "none")
--database value, -d value   database for your app, ex. Postgres (default: "none")
--help, -h                   show help
```

> 🔔 Please note: by default, `cgapp create` command without any options will create into current folder default backend (`net/http`) without frontend or configured Docker containers!

### `deploy`

CLI command for deploy Docker containers with your project to a remote server or run on your local machine.

#### Usage

```console
cgapp deploy [command options] [arguments...]
```

#### Options

```console
--playbook value, -p value  name of Ansible playbook, ex. my-play.yml (default: "deploy-playbook.yml")
--username value, -u value  username of remote's server user or your local machine, ex. root (default: "none")
--host value, -s value      host name of remote's server or local machine (from Ansible inventory), ex. do_server_1 (default: "localhost")
--network value, -n value   network for Docker containers, ex. my_net (default: "cgapp_network")
--ask-become-pass           asking you to enter become user's password at start (default: false)
--help, -h                  show help
```

> 🔔 Please note: you should only run the `cgapp deploy` command from the root folder of your project, which created with the `cgapp create` command! _It's a necessary condition for everything to work perfectly._

## 📝 Available production-ready app templates

**Golang:**

- `net/http` [create-go-app/net_http-go-template](https://github.com/create-go-app/net_http-go-template) — Backend template with built-in net/http ([pkg/net/http](https://golang.org/pkg/net/http/)).
- `fiber` [create-go-app/fiber-go-template](https://github.com/create-go-app/fiber-go-template) — Backend template with Fiber ([gofiber/fiber](https://github.com/gofiber/fiber)).
- `echo` _WIP_ [create-go-app/echo-go-template](https://github.com/create-go-app/echo-go-template) — Backend template with Echo ([labstack/echo](https://github.com/labstack/echo)).

**JavaScript:**

- `react-js` [create-go-app/react-js-template](https://github.com/create-go-app/react-js-template) — Frontend template with React.js ([facebook/react](https://github.com/facebook/react)).
- `react-ts` [create-go-app/react-ts-template](https://github.com/create-go-app/react-ts-template) — Frontend template with React.js TypeScript ([facebook/react](https://github.com/facebook/react)).
- `preact` [create-go-app/preact-js-template](https://github.com/create-go-app/preact-js-template) — Frontend template with Preact ([preactjs/preact](https://github.com/preactjs/preact)).

## 🐳 Available production-ready Docker containers

**Web/Proxy server:**

- `nginx` [create-go-app/nginx-docker](https://github.com/create-go-app/nginx-docker) — Docker container with Nginx.

**Database:**

- `postgres` _WIP_ [create-go-app/postgres-docker](https://github.com/create-go-app/postgres-docker) — Docker container with PostgreSQL.

## 👤 User templates & containers

Create Go App CLI provide creation your own template, instead of those prepared by authors. It's easy! 😉

Just specify backend (`-b`), frontend (`-f`), webserver (`-w`) and database (`-d`) with addresses to your repositories and run:

```bash
cgapp create \
             -b github.com/user/my-back-template \
             -f gitlab.com/user/my-front-template \
             -w github.com/user/my-webserver-container-template \
             -d bitbucket.org/user/my-database-container-template
```

> 🔔 Please note: the `https://` protocol will add automatically!

## 🤔 FAQ

**— How to update CLI to latest version?**

You can just re-build the CLI. The latest version will be downloaded and installed automatically:

```console
go build -i -o $GOPATH/bin/cgapp github.com/create-go-app/cli
```

If you're using _standalone_ version, please go to the [release page](https://github.com/create-go-app/cli/releases) and download archive with a new version.

**— What do you use to automate the server deployment process?**

Each project is created with the required set of configs to start the build and deployment process on the production server or local machine. We use a helpful tool, called **[Ansible](https://docs.ansible.com)** for automate this.

In the root folder of the project you will find [`deploy-playbook.yml`](https://github.com/create-go-app/cli/blob/master/configs/deploy-playbook.yml) file. Is the Ansible playbook describing the build app & deployment to server process.

> 👀 Hey! Don't worry, if you are not familiar with this technology, read [this article](https://docs.ansible.com/ansible/latest/user_guide/playbooks.html) from the docs. _Besides, you can ask us about it in one of the issues._

**— What does this playbook do?**

- [x] Builds production-ready backend (Go) & frontend (JavaScript, TypeScript) apps, that you have chosen
- [x] Provides the best practices for web server and database configuration
- [x] Configures Docker network with containers for backend, static files from frontend, web server and database
- [x] Runs these Docker containers on your remote server or local machine

> 👌 We recommend to using our default playbook, but you are free to change them any way you want!

**— What should I do for deploy my project?**

1. Check, that you have Ansible `v2.9.x` (_or later_) is installed.
2. Add the right host to your inventory file (_into `/etc/ansible/hosts`_) on your local machine.
3. Be sure, you're working with a SSH key correctly:
   - Generate a new SSH key by command `ssh-keygen` on your local machine;
   - Add a **public** key part (_with `*.pub` extension_) into a bottom of `~/.ssh/authorized_keys` file (_with "one line" format_) on your remote server;
   - Save a **private** key part on your local machine;
4. Run the built-in `cgapp deploy` command (_from the root folder of your project_):

```console
cgapp deploy -p <PLAYBOOK_NAME> -u <USER> -s <HOST> -n <NETWORK_NAME> [EXTRA_VARS...]
```

- `<PLAYBOOK_NAME>` (_optional_, default: `deploy-playbook.yml`) is the Ansible playbook name
- `<USER>` (**required**) is an username of remote's server user (for example, `root`)
- `<HOST>` (_optional_, default: `localhost`) is a host name from your inventory file
- `<NETWORK_NAME>` (_optional_, default: `cgapp_network`) is a network name for your Docker containers

> 👌 If you need to deploy (_or run_) project with entering password for user, you can add `--ask-become-pass` variable in `[EXTRA_VARS...]` section. You will be prompted to enter `<USER>`'s password at the beginning. _This is an alias of the standard Ansible argument to ask for a privilege escalation password (see [docs](https://docs.ansible.com/ansible/latest/user_guide/become.html#become-command-line-options))._

**— Are there any video examples of working with the Create Go App CLI?**

<p align="center">
   <a href="https://youtu.be/e9443CCqxio">
      <img src="https://github.com/create-go-app/cli/blob/master/.github/images/youtube-preview.png" alt="cgapp youtube example"/><br/>
      🔗 youtu.be/e9443CCqxio
   </a>
</p>

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development `create-go-app/cli`:

1. Add a :octocat: GitHub Star to the project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Set%20up%20a%20new%20Go%20%28Golang%29%20full%20stack%20app%20by%20running%20one%20CLI%20command%21%26url%3Dhttps%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://shrts.website/do/server) (your profit is **\$100** and we get \$25).
5. Buy awesome [domain name with **5%** discount](https://shrts.website/reg/domain) at REG.COM.

Thanks for your support! 😘 Together, we make this project better every day.

### ~ Sponsors

| Logo                                                                                                           | Description                                                                                                                                         | URL                              |
| -------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| <img align="center" width="100px" src="https://github.com/create-go-app/cli/blob/master/.github/images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — Team who making UX efficiency review, friendly UI design, smart backend microservices, high-quality web apps and many more. | [https://1wa.co](https://1wa.co) |
|                                                                                                                | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div>                                           |                                  |

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
