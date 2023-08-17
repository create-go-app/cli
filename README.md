[![Create Go App][repo_logo_img]][repo_url]

# Create Go App CLI – Create a new production-ready project by running one command

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![Wiki][repo_wiki_img]][repo_wiki_url]
[![License][repo_license_img]][repo_license_url]

Create Go App (or `cgapp` for a short) is a **complete** and **self-contained** 
solution for developers of **any qualification** to create a production-ready 
project with **backend** (Go), **frontend** (JavaScript, TypeScript) and 
**deploy automation** (Ansible, Docker) by running **only one** CLI command.

Focus on **writing your code** and **thinking of the business-logic**! The CLI 
will take care of the rest.

## ⚡️ Quick start

First, all GNU/Linux and macOS users available way to easily install `cgapp` via
[Homebrew][brew_url].

Tap a new formula:

```console
brew tap create-go-app/tap
```

Install the latest `cgapp` version:

```console
brew install create-go-app/tap/cli
```

> 💡 Hint: see the repository's [Release page][repo_releases_url], if you 
> want to download MS Windows `exe` files, ready-made `deb`, `rpm`, `apk` 
> or `Arch Linux` packages.

An alternative way is a native installation using Go (version `1.21.0` or 
higher is required):

```console
go install github.com/create-go-app/cli/v5/cmd/cgapp@latest
```

Next, generate an example configuration file in the current dir:

```console
cgapp -init
```

When the generation process is finished, open the `.cgapp.yml` file and fill 
out your credentials and other info.

> ❗️ Please be careful: please **do not** specify secret data in plain text
> in the configuration file! There are environment variables (`env`) for this,
> which you can use in the `.cgapp.yml` file for absolutely any settings.
> 
> 💡 Hint: you can always refer to our [Wiki page][repo_wiki_url] to set the 
> configuration correctly.

Let's create a new project:

```console
cgapp -create
```

Now, make something awesome in the `backend` and `frontend` dirs, and you're 
ready to automatically deploy to the remote host:

```console
cgapp -deploy
```

That's all you need to know to start! 🎉

### 🐳 Docker-way to quick start

If you don't want to install Create Go App CLI to your system, you feel free 
to using our official [Docker image][docker_url] and run CLI from isolated 
container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/cgapp:latest [COMMAND]
```

> ❗️ Warning: the `-deploy` command is currently **unavailable** in this 
> image.

## 📖 Project Wiki

The best way to better explore all the features of the **Create Go App CLI** 
is to read the project [Wiki][repo_wiki_url] and take part in 
[Discussions][repo_discussions_url] and/or [Issues][repo_issues_url]. 

Also, the most frequently asked questions (FAQ) is [here][repo_wiki_faq_url].

## ⚙️ Commands & Options

### `-init`

CLI command for generate an example config file (called `.cgapp.yml`) in the 
current dir.

```console
cgapp -init
```

### `-create`

CLI command for create a new project in the current dir.

```console
cgapp -create [OPTION]
```

| Option | Description                                            | Type     | Default        | Required? |
|--------|--------------------------------------------------------|----------|----------------|-----------|
| `-p`   | Enables to define a path (or URL) for the config file. | `string` | `".cgapp.yml"` | No        |

![cgapp_create][cgapp_create_gif]

- 📺 Full demo video: https://recordit.co/OQAwkZBrjN
- 📖 Docs: https://github.com/create-go-app/cli/wiki/Command-create

### `-deploy`

CLI command for deploy Docker containers with your project via Ansible to 
the remote host.

```console
cgapp -deploy [OPTION]
```

Make sure that you have [Python 3.8+][python_url] and
[Ansible 2.9+][ansible_url] installed on your computer.

| Option | Description                                            | Type     | Default        | Required? |
|--------|--------------------------------------------------------|----------|----------------|-----------|
| `-p`   | Enables to define a path (or URL) for the config file. | `string` | `".cgapp.yml"` | No        |

![cgapp_deploy][cgapp_deploy_gif]

- 📺 Full demo video: https://recordit.co/ishTf0Au1x
- 📖 Docs: https://github.com/create-go-app/cli/wiki/Command-deploy

## 📝 Production-ready project templates

### Backend

- Backend template with Go built-in [net/http][net_http_url] package:
  - [`net/http`][cgapp_net-http-template_url] — simple REST API with CRUD 
    and JWT auth.
- Backend template with [Fiber][fiber_url]:
  - [`fiber`][cgapp_fiber-template_url] — complex REST API with CRUD, JWT auth 
    with renew token, DB and cache.
- Backend template with [go-chi][chi_url]:
  - [`chi`][cgapp_chi-template_url] — a basic application with health check.

### Frontend

- Pure JavaScript frontend template:
  - `vanilla` — generated template with pure JavaScript app.
  - `vanilla-ts` — generated template with pure TypeScript app.
- Frontend template with [React][react_url]:
  - `react` — generated template with a common React app.
  - `react-ts` — generated template with a TypeScript version of the React app.
  - `react-swc` — generated template with a common React app with SWC.
  - `react-swc-ts` — generated template with a TypeScript version of the React app with SWC.
- Frontend template with [Preact][preact_url]:
  - `preact` — generated template with a common Preact app.
  - `preact-ts` — generated template with a TypeScript version of the Preact app.
- Frontend template with [Vue.js][vuejs_url]:
  - `vue` — generated template with a common Vue.js app.
  - `vue-ts` — generated template with a TypeScript version of the Vue.js app.
- Frontend template with [Svelte][svelte_url]:
  - `svelte` — generated template with a common Svelte app.
  - `svelte-ts` — generated template with a TypeScript version of the Svelte app.
- Frontend template with [Solid][solid_url]:
  - `solid` — generated template with a common Solid app.
  - `solid-ts` — generated template with a TypeScript version of the Solid app.
- Frontend template with [Lit][lit_url] web components:
  - `lit` — generated template with a common Lit app.
  - `lit-ts` — generated template a TypeScript version of the Lit app.
- Frontend template with [Qwik][qwik_url] web components:
  - `qwik` — generated template with a common Qwik app.
  - `qwik-ts` — generated template a TypeScript version of the Qwik app.
- Frontend template with [Next.js][nextjs_url]:
  - `next` — running interactive wizard for generate the Next.js app.
- Frontend template with [Nuxt][nuxt_url]:
  - `nuxt` — running interactive wizard for generate the Nuxt v3 app.

> ☝️ Frontend part will be generated using awesome tool [Vite.js][vitejs_url] 
> under the hood. So, you'll always get the latest version of `React`, 
> `Preact`, `Vue`, `Svelte`, `Solid`, `Lit`, `Qwik`, or pure 
> JavaScript/TypeScript templates for your project!
> 
> The `Next.js` and `Nuxt` frontend parts will be generated using the 
> `create-next-app` and `nuxi` utilities.
>
> Please make sure that you have `npm` version `7` or higher installed to 
> create the frontend part of the project correctly. If you run the 
> `cgapp -create` command using our [Docker image][docker_url], `npm` of the 
> correct version is **already** included.

## 🚚 Pre-configured Ansible roles

### Web/Proxy server

- Roles for run Docker container with [Traefik Proxy][traefik_url]:
  - `traefik` — configured Traefik container with a simple ACME challenge 
    via CA server.
  - `traefik-acme-dns` — configured Traefik container with a complex ACME 
    challenge via DNS provider.
- Roles for run Docker container with [Nginx][nginx_url]:
  - `nginx` — pure Nginx container with "the best practice" configuration.

> ✌️ Since Create Go App CLI `v2.0.0`, we're recommended to use **Traefik 
> Proxy** as default proxy server for your projects. The main reason: this 
> proxy provides _automatic_ SSL certificates from Let's Encrypt out of the 
> box. Also, Traefik was built on the Docker ecosystem and has a _really 
> good-looking_ and _useful_ Web UI.

### Database

- Roles for run Docker container with [PostgreSQL][postgresql_url]:
  - `postgres` — configured PostgreSQL container with apply migrations for 
    backend.

### Cache (key-value storage)

- Roles for run Docker container with [Redis][redis_url]:
  - `redis` — configured Redis container for backend.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of 
`Create Go App CLI`:

- Add a [GitHub Star][repo_url] to the project.
- Tweet about project [on your Twitter][twitter_url] for all of your friends 
  and colleagues.
- Write interesting articles about project on [Dev.to][dev_to_url], or 
  personal blog.
- Join DigitalOcean at our [referral link][author_do_ref_url] (your profit is 
  **$100**, and we get $25).
- Leave a review on our [ProductHunt][cgapp_product-hunt_url] page.

[![Product Hunt][cgapp_product-hunt_img]][cgapp_product-hunt_url]

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.

Together, we can make this project **better** every day! 😘

## ⚠️ License

[`Create Go App CLI`][repo_url] is free and open-source software licensed under 
the [Apache 2.0 License][repo_license_url]. Official [logo][repo_logo_url] was 
created by [Vic Shóstak][author] and distributed under 
[Creative Commons][repo_cc_url] license (CC BY-SA 4.0 International).

<!-- Go -->

[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/create-go-app/cli
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-88%25-success?style=for-the-badge&logo=none
[go_dev_url]: https://pkg.go.dev/github.com/create-go-app/cli/v5

<!-- Repository -->

[repo_url]: https://github.com/create-go-app/cli
[repo_logo_url]: https://github.com/create-go-app/cli/wiki/Logo
[repo_logo_img]: https://github.com/create-go-app/cli/assets/11155743/95024afc-5e3b-4d6f-8c9c-5daaa51d080d
[repo_license_url]: https://github.com/create-go-app/cli/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_v2_url]: https://github.com/create-go-app/cli/tree/v2
[repo_v3_url]: https://github.com/create-go-app/cli/tree/v3
[repo_issues_url]: https://github.com/create-go-app/cli/issues
[repo_pull_request_url]: https://github.com/create-go-app/cli/pulls
[repo_discussions_url]: https://github.com/create-go-app/cli/discussions
[repo_releases_url]: https://github.com/create-go-app/cli/releases
[repo_wiki_url]: https://github.com/create-go-app/cli/wiki
[repo_wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[repo_wiki_faq_url]: https://github.com/create-go-app/cli/wiki/FAQ

<!-- Project -->

[cgapp_deploy_gif]: https://user-images.githubusercontent.com/11155743/116796941-3c421e00-aae9-11eb-9575-d72550814d7a.gif
[cgapp_create_gif]: https://user-images.githubusercontent.com/11155743/116796937-38160080-aae9-11eb-8e21-fb1be2750aa4.gif
[cgapp_product-hunt_url]: https://www.producthunt.com/posts/create-go-app?utm_source=badge-review&utm_medium=badge&utm_souce=badge-create-go-app#discussion-body
[cgapp_product-hunt_img]: https://api.producthunt.com/widgets/embed-image/v1/review.svg?post_id=316086&theme=light
[cgapp_chi-template_url]: https://github.com/create-go-app/chi-go-template
[cgapp_fiber-template_url]: https://github.com/create-go-app/fiber-go-template
[cgapp_net-http-template_url]: https://github.com/create-go-app/net_http-go-template

<!-- Author -->

[author]: https://github.com/koddr
[author_do_ref_url]: https://m.do.co/c/b41859fa9b6e

<!-- Readme links -->

[twitter_url]: https://twitter.com/intent/tweet?text=Wow%21%20%F0%9F%8E%89%20Create%20a%20new%20production-ready%20project%20with%20backend%20%28Golang%29%2C%20frontend%20%28JavaScript%2C%20TypeScript%29%0Aand%20deploy%20automation%20%28Ansible%2C%20Docker%29%20by%20running%20one%20CLI%20command%20%F0%9F%91%89%20https%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli
[dev_to_url]: https://dev.to/
[redis_url]: https://redis.io/
[postgresql_url]: https://postgresql.org/
[nginx_url]: https://nginx.org/
[traefik_url]: https://traefik.io/traefik/
[vitejs_url]: https://vitejs.dev/
[vuejs_url]: https://vuejs.org/
[react_url]: https://reactjs.org/
[preact_url]: https://preactjs.com/
[nextjs_url]: https://nextjs.org/
[nuxt_url]: https://nuxt.com/
[svelte_url]: https://svelte.dev/
[solid_url]: https://github.com/solidjs/solid
[lit_url]: https://lit.dev/
[qwik_url]: https://github.com/BuilderIO/qwik
[chi_url]: https://github.com/go-chi/chi
[fiber_url]: https://github.com/gofiber/fiber
[net_http_url]: https://golang.org/pkg/net/http/
[docker_url]: https://hub.docker.com/r/koddr/cgapp
[python_url]: https://www.python.org/downloads/
[ansible_url]: https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-specific-operating-systems
[brew_url]: https://brew.sh/
