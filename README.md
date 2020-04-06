# CODENAMES GENERATOR
Generates values so you can play a virtual game of Codenames (currently only supports generation of the words)

## Install
Clone the repo
```
git clone git@github.com:wamphlett/codenames-generator.git

```
Make the executable from the repo root
```
make build
```

## Usage
First you need a file with some words, this needs to be done in the format [found here](https://github.com/wamphlett/codenames-generator/blob/master/wordlist.example.json)

```
./codename-generator generate words
```

For more commands
```
./codename-generator --help
```


## Todo
- Make the generator able to generate a grid with the colors
