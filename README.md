# Weather Buddy

A simple CLI written in Go that fetches the current weather from weatherapi.com

## Installation

```bash
go install .
```

## Usage

### Get weather in Heilbronn, Germany

```bash
wbuddy
```

### Get weather of any location

```bash
wbuddy [LOCATIONNAME_HERE]

# Examples
wbuddy Paris
wbuddy Tokyo
wbuddy Istanbul

```

## Contributing

### Starting

```bash
# run local
make run

# build windows
make buildwin
```

### Todos

- [ ] Fix location strings with spaces not working
- [ ] Extend README.md
