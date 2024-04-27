# snip-doc

snip-doc is a web application developed in Go, following along with Alex
Edwards' book, [_Let's Go_](https://lets-go.alexedwards.net/). The goal is to
learn how web services work in Go by building something. This repository is more
of a stepping stone to building production ready, scalable, cloud micro
services.

This book and the second part,
[_Let's Go Further_](https://lets-go-further.alexedwards.net/), are known for
their practical and objective way of building secure, scalable web applications
efficiently. They were also selected for the following reasons:

- Intentional usage of Go standard library
- Very few external dependencies
- Updated in April of 2024, using version 1.22
- Cost effective, about $150 for both books and source code
- Outstanding
  [reviews](https://www.goodreads.com/book/show/58044798-let-s-go-further?from_search=true&from_srp=true&qid=A4Z17K7RKM&rank=1) -
  `4.8` stars (even
  [Reddit](https://www.reddit.com/r/golang/comments/vp3ejz/is_lets_go_and_lets_go_further_worth_it/)
  agrees)

You can learn more about Alex and the books
[here](https://www.alexedwards.net/), he has been teaching Go professionally for
10+ years.

At it's core, snip-doc will be based on the book's application dubbed
_Snippetbox_. It offers a straight forward text snippet upload tool, similar to
Pastebin or Github Gists. I am changing the name to help signify the 2 different
code bases when developing the project.

## Learning Path

Pleas check out the [Conclusion](#conclusion) section to see resources that led
to this path.

### Let's Go

This section goes over the plan to efficietly understand _Let's Go_, while
minimizing time spend on front end and database operations.

The app will have a fully functioning front end and user authentication
database. But, I will lightly skim over the HTML templates and CSS parts. I will
lean a bit more into the database setup and operations but not much, only enough
to move along with the book project.

Again, the short term goal is to learn Go's web service capabilities with this
book as quickly and efficiently as possible. The long term goal is to dive deep
into the next installment of the series, _Let's Go Further_, and build
production ready, scalable, cloud micro services.

With this in mind, in both book projects, I will be taking more time and writing
the Go by hand to understand it better.

### Let's Go Further

When it comes to the second book, I will attempt to take a slightly different
approach and follow _TDD_ (Test Driven Development) guidelines from Chris James'
[Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests). Explanations
as to why this is a good idea are included in the link.

The following section gets you Go-ing with installing, running, and
understanding the project.

## Getting Started

To get started on your local machine, you only need [Go](https://go.dev)
installed, and a CLI prompt.

Download the project from Github, and run the following command from the root
directory to build and run the application.

```bash
go run ./cmd/web
```

It will beginning listening on port `4000` but you can also use any port with
the following command.

```bash
go run ./cmd/web/ -addr=":4200"
```

### Hot Reload

For hot reload of your application, try [air](https://github.com/cosmtrek/air).
This is not required but helps speed up development, it is similar to
[nodemon](https://github.com/remy/nodemon), install with this command.

```bash
go install github.com/cosmtrek/air@latest
```

Initialize _air_, you only need to do this once per project.

```bash
air init
```

Then, just run _air_ for hot reload during development. Use the `sigterm` for
your application to quit, in our case, `ctrl` + `c` simultaneously.

```bash
air
```

### Project Structure

Directory tree example, through chapter 5.

```bash
.(root)
├── README.md
├── cmd
│   └── web
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       ├── routes.go
│       ├── templates.go
│       └── tmp
│           ├── build-errors.log
│           └── main
├── go.mod
├── go.sum
├── internal
│   └── models
│       ├── errors.go
│       └── snippets.go
├── tmp
└── ui
    ├── html
    │   ├── base.tmpl
    │   ├── pages
    │   │   ├── home.tmpl
    │   │   └── view.tmpl
    │   └── partials
    │       └── nav.tmpl
    └── static
        ├── css
        │   └── main.css
        ├── img
        │   ├── favicon.ico
        │   └── logo.png
        └── js
            └── main.js

15 directories, 20 files
```

### Routes

Route pattern examples, through chapter 5.

| Route Pattern          | Handler           | Action                                    |
| ---------------------- | ----------------- | ----------------------------------------- |
| GET /{$}               | home              | Display home page                         |
| GET /snippet/view/{id} | snippetView       | Display a specific spippet                |
| GET /snippet/create    | snippetCreate     | Display a form for creating a new snippet |
| POST /snippet/create   | snippetCreatePost | Save a new snippet                        |
| GET /static/           | http.FileServer   | Serve a specific static file              |

## Conclusion

This project will be built, then mostly left here as a learning resource.

### Due Diligence

In March and April of 2024, many hours of research went into technology forums
and other blogs to select the most efficient and cost effctive way to build
proper web services. Not only with Go, but other languages and technologies as
well.

In the end, Go seemed most suitable for the use case of efficiently building
production ready, scalable, cloud micro services.

#### Research Resources

Along with an assortment of forums, and content creators I follow. I took a lot
of value from these resources to find the right tools.

- [Offical Go Docs](http://go.dev/doc)
- [Rob Pike (Go creator) - Concurrency is not parallelism](https://go.dev/blog/waza-talk) -
  Lecture
- [C.A.R. Hoare - Communicating Sequential Processes](https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf) -
  Academic paper recommended by Rob Pike
- [Mat Ryer (Engineering Director @ Grafana Labs) - How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) -
  Blog
- [Chris James - Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests) -
  Learning site recommended by Mat Ryer
- [Go by Example](https://gobyexample.com/) - Learning site recommended by
  Chris James

#### Famous Go Projects

These projects are heavily written in Go, lots of cloud based services.

- Docker
- Kubernetes
- Terraform
- Grafana
- Prometheus
- Istio
- Hugo
- Traefik
- Ollama
- Caddy
- frp
- etcd
- fzf
- esbuild
