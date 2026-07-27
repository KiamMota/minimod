# mm (minimod)

The easiest way to start a Go project.

## Installation

```bash
go install github.com/kiammota/minimod@latest
```

## Usage

Create a project:

```bash
mm myproject
```

Create a project with an MIT license:

```bash
mm myproject -l MIT
```

## Output

```text
myproject/
├── go.mod
├── main.go
├── .gitignore
├── README.md
└── LICENSE
```

## Why?

Instead of running:

```bash
mkdir myproject
cd myproject
go mod init myproject
touch main.go
```

just run:

```bash
mm myproject
```

## License

MIT
