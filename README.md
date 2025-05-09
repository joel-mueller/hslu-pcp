# HSLU PCP Projekt Go

Joel Müller und Leo Hrvat

## TODO vor abgabe

- [ ] Code snippets nochmals prüfen die im bericht drin sind
- [ ] Rechtschreibung prüfen bericht
- [ ] TODos im bericht fertig machen
- [ ] Uber code drüber schauen und evt noch mehr tests schreiben

## Prerequisites

Before running the project, ensure that Go is installed on your system.

To check if Go is installed and verify the version, run:
```shell
go version
```

If go is not installed, follow the [Installation Guide](https://go.dev/doc/install) and try again.

## Run the application

To run the application, you need to be in the root of this project. Then type in the terminal

```shell
go run .
```

To build the binary of the project, run following command in the root of the Project. Afterwards, you should get a binray file called `hslu-pcp`

```shell
./package.sh
```

Now, you should have a binray file of the application in your directory. To run this, run following command:

```shell
./hslu-pcp
```

## Convert the Bericht to PDF Document

Note: To convert the Bericht to a pdf, you need to have `pandoc` and `tectonic` installed.

```shell
./submission.sh
```