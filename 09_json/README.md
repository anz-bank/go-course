# Lab 9 - JSON puppy

- Create directory `09_json/USERNAME` containing a copy of upstream master `08_project/USERNAME`
- Add JSON tags to puppy data type
- Test marshalling and unmarshalling using [require.JSONEq](https://godoc.org/github.com/stretchr/testify/require#JSONEq)
- Add command line flag `-d FILE` with long form `--data FILE` using [kingpin.v2](https://godoc.org/gopkg.in/alecthomas/kingpin.v2)
- FILE should contain an array of puppies in JSON format. Parse this file and store its contents.
