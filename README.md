# coco 

<p align="center"><a href="#readme"><img src="https://user-images.githubusercontent.com/3170758/198417756-46a68599-9c53-4ab5-948e-033b74ff1dc3.png" alt="logo"/></a></p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/synerops/coco"><img src="https://pkg.go.dev/badge/github.com/synerops/coco.svg" alt="docs" /></a>
  <a href="https://goreportcard.com/report/github.com/synerops/coco"><img src="https://goreportcard.com/badge/github.com/synerops/coco" alt="GoReportCard" /></a>
  <a href="https://github.com/synerops/coco/blob/main/LICENSE"><img alt="license" src="https://img.shields.io/github/license/synerops/coco"></a>
</p>

`coco` is a _colorized-intended_ library to provide an easy way to display logs on cli-based applications.

Its purpose is to be flexible and easy to configure. In simple words: **get what you set**.

Accompanied by the fantastic library [color](https://github.com/fatih/color), `coco` is aimed at identifying each 
log level by a color you can define (or use our default proposal)

## `[INFO]: Install`

```shell
go get github.com/synerops/coco
```

## `[INFO]: Examples`

### `Getting Started`

Without any extra customization, you can configure the library as follows.

```go
// default options pre-defined in the package
opts := coco.Default()

// new instance of the library
ui := coco.New(opts)

// use with one of the levels already defined, 
// in this case: SUCCESS
ui.Success("this is a successful message")

// other served methods are:
ui.Error(errors.New("this is a sample error message"))

ui.Warning("this is a sample warning message")

ui.Info("this is a sample informative message")
```

### `Default Options`

This is the configuration defined behind `coco.Default()`

#### Levels

- **error**, displayed in the terminal using `os.Stderr`
- **success**, displayed in the terminal using `os.Stdout`
- **warning**, displayed in the terminal using `os.Stdout`
- **info**, hidden by default using `io.Discard`, this setting can be changed

#### Colors

- **error**: `color.FgRed`
- **success**: `color.FgGreen`
- **warning**: `color.FgYellow`
- **info**: `color.FgHiCyan`

![colors](https://user-images.githubusercontent.com/3170758/197928467-a369d767-7c25-410f-8db7-0816ea1ec75f.png)

### Customization

#### Change the printable format

```go
// tl;dr
opts.SetFormat(coco.Error, "[%s] -> %s\n")
```

By default, the format to display each level is: `[%s]: %s\n` 
(that's why the level legend is involved in squared brackets). You can change that per level using `opts.SetFormat`.

```go
// default options pre-defined in the package
opts = coco.Default()

// change the format only on the `error` level
opts.SetFormat(coco.Error, "[%s] -> %s\n")
```

![format](https://user-images.githubusercontent.com/3170758/197928510-edf9353f-19c2-4a2b-88fd-25bd41acfb12.png)

#### Change the output writer (a.k.a `io.Writer`)

```go
// tl;dr
opts.SetWriter(coco.Info, os.Stdout)
```

You can change the output of the message and use your own Writer
(based on the [`io.Writer` interface](https://pkg.go.dev/io#Writer)).

A good example of this method is to use accompanied with a flag to determine whether the application is under
debug mode.

```go
// check whether the DEBUG env var is defined and change the output of the INFO level messages
if debugMode := os.Getenv("DEBUG"); debugMode != "" {
    opts.SetWriter(coco.Info, os.Stdout)
}
```

#### Change the color (a.k.a `color.Attribute`)

```go
// tl;dr
opts.SetColor(coco.Warning, color.FgHiYellow)
```

You can change the color of the level legends to satisfy your particular needs, just make sure to provide an attribute 
based on the [`color.Attribute` constant](https://pkg.go.dev/github.com/fatih/color#Attribute).

```go
opts.SetColor(coco.Warning, color.BgHiRed)
```

![color](https://user-images.githubusercontent.com/3170758/197928538-a3ba5994-41e9-414a-80da-497cd868aba2.png)

### New levels

Coco aims to fit your needs, so you can create custom Levels and make them part of the app.

```go
// default options pre-defined in the package
opts := coco.Default()

// define a new level called NOTICE
const levelName = "notice"

// use coco's Output struct to define the structure of the new level
noticeOutput := coco.Output{
    Format: coco.GlobalFormat,
    Writer: os.Stdout,
    Color:  color.FgHiMagenta,
}

if err := opts.NewLevel(levelName, noticeOutput); err != nil {
    fmt.Printf("error while defining a new level: %s\n", err)
}

// new instance of the library
ui := coco.New(opts)

// use with one of the level you just created 
ui.Log(levelName, "this is a sample notice message")
```

![notice](https://user-images.githubusercontent.com/3170758/197928562-31bb5177-fb3f-4cab-becf-efe1983769d9.png)

## `[INFO]: License`

The MIT License (MIT) - see [`LICENSE`](https://github.com/synerops/coco/blob/main/LICENSE) for more details

## `[INFO]: Contribute`

Feel free to contribute to any aspect of the project. Any initiative will be highly appreciated.

## `[INFO]: TO-DO`

- [x] Improve documentation
- [x] Export the package to [Go Packages](https://pkg.go.dev/)
- [ ] Create a method to change the `GlobalFormat` and apply it across the levels

## `[INFO]: Sponsor`

`coco` is sponsored by my tech startup [SynerOps](https://synerops.com) where I'm the founder and solo-developer. 
Feel free to reach me out.

![synerops](https://avatars.githubusercontent.com/u/70786831?s=50&v=4)
