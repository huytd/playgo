# Playgo - a real playful Golang playground

[![Build Status](https://travis-ci.org/huytd/go-play.svg?branch=master)](https://travis-ci.org/huytd/go-play)

![](screenshot.png)

## Why use this?

Because:

- Can import **any** package
- ~~The official Playground's UI sucks~~ :trollface:
- This one run locally with no need of container/docker
- Smaller code-base, easier to customize
- Use it the way you want (cli mode and web mode)
- Forget the code format, just type the code anyway you want
- ... (add more awesome stuff here) ...

## How to install?

```
go get -u github.com/huytd/go-play
```

That's all! (Given that you have a working `GOPATH` configured)

Or run in Docker container:
```
docker build -t go-play .
# may be: docker push ...
docker run -d -p 3000:3000 go-play
```

## What is this?

This is the simple version of Go Playground. It run locally with no container needed.

There are *2* modes available:

### Command-line mode

![](climode.png)

Also called as `cli` mode. In this mode, the playground will read the code from `os.Stdin` and execute it. This enable the ability to integrate `playgo` with other editors such as **vim**, **sublime**, **atom**,...

Usage:

```
echo 'print("Hello, do some math, 1 + 1 = ", 1 + 1)' | playgo
```

or

```
cat something.txt | playgo
```

### Web mode

![](webmode.png)

If you don't like using `cli` mode, you can use the web IDE by run the following command:

```
playgo -mode=web
```

The web IDE will be started at [http://localhost:3000](http://localhost:3000) by default.

You can change the port by:

```
PLAYGO_PORT=8080 playgo -mode=web
```

## License
This project is licensed under the terms of the **MIT** license.

## For Developers

It would be nice if you want to contribute to this project. I really need your help, there are a lot of things to do.

Feel free to create a pull request or make an issue to report bugs/request new features. You can see the list of things to do below.

:bow:

## To Do:

[x] Support `gofmt` - Recommend to use it only in full code mode
[ ] Fully Support `import` - Currently available in full code mode
[ ] Support `func` - Currently available in full code mode
[ ] Support `channel` (real problem is: streaming output)
[ ] Execute code in containers or isolated environment for more security?
[ ] Create a command line code editor / or a VIM plugin?
[ ] Autocomplete for web IDE
