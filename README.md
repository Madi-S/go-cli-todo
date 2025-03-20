# GO CLI Todo application

## Usage

Help:

```bash
go run . --help
```

Print existing todos:

```bash
go run . --print | go run . -print
```

Add todo:

```bash
go run . --add "Watch TV"
```

Remove todo by index:

```bash
go run . --delete 0
```

Toggle todo by index:

```bash
go run . --toggle 1
```

Edit existing todo:

```bash
go run . -edit "1:Play football"
```
