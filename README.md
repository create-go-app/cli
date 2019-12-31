<h1 align="center">🕶 Create Go App CLI</h1>
<h3 align="center">Set up a new Go (Golang) full stack app by running one CLI command!</h3>

<img width="100%" src="images/screenshot.jpg" alt="Create Go App screenshot"/>

<br/>

## The Why?

<img align="right" width="196px" src="images/logo_gopher.png" alt="Create Go App logo"/>

Coming soon...

<br/>

## Install

Coming soon...

## Usage

```console
foo@bar:~$ cgapp [command] [value]
```

> **For example:** if you want to create new app called `my-app` into `./app` folder with [Echo](https://github.com/labstack/echo) on backend and [Preact](https://github.com/preactjs/preact) as frontend, just run:
>
>```console
>foo@bar:~$ cgapp -n my-app -b echo -f preact -p ./app
>```

You can see all available CLI commands (short and full names) by running `cgapp` with `--help`:

```console
foo@bar:~$ cgapp --help

GLOBAL OPTIONS:
   --name value, -n value      name of your go module, ex. github.com/user/my-app
   --path value, -p value      path to create app, ex. ~/projects/my-app (default: ".")
   --backend value, -b value   backend for your app, ex. Echo, Gin, Iris, net/http
   --frontend value, -f value  frontend for your app, ex. (P)React, Vue, Svelte (default: "none")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

> **Note:** `--path` and `--frontend` commands are _optional_, because they have default values.

## Available production-ready templates

**Golang:**

- [create-go-app/echo-go-template](https://github.com/create-go-app/echo-go-template) — Backend template with Echo ([labstack/echo](https://github.com/labstack/echo)).

**JavaScript:**

- [create-go-app/preact-js-template](https://github.com/create-go-app/preact-js-template) — Frontend template with Preact ([preactjs/preact](https://github.com/preactjs/preact)).

### User templates

Create Go App CLI provide creation your own template, instead of those prepared by authors. It's easy! Just specify `--backend` (or/and `--frontend`) with HTTPS address to your repository and run:

```console
foo@bar:~$ cgapp -n another-app -b https://github.com/user/my-backend-template
```

## Requirements

- Go `1.11+`
- Go Modules

## Developers

- Idea and active development by [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Project assistance

If you want to say «thank you» or/and support active development `create-go-app/cli`:

1. Add a GitHub Star to project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Set%20up%20a%20new%20Go%20%28Golang%29%20full%20stack%20app%20by%20running%20one%20CLI%20command%21%26url%3Dhttps%3A%2F%2Fgithub.com%2Fcreate-go-app%2Fcli).
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **\$100** and we get \$25).

Thanks for your support! 😘 Together, we make this project better every day.

### Sponsors

| Logo | Sponsor description | URL |
| ------------ | ------------- | ------------- |
| <img align="center" width="100px" src="images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — IT specialists around the world, who are ready to share their experience to solve your business objectives. | [https://1wa.co](https://1wa.co) |
|  | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div> |  |

## License

MIT
