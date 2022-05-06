
<img src="https://i.ibb.co/FmgGyD8/Facebook-cover-1.png" width="100%" />

# Binosearch

An auxiliary scanner written in Go that checks if your API is vulnerable to OWASP Top 10 API vulnerabilities.




## Authors

- [Akia Japhet Mamaoag - @mamaoag](https://www.github.com/mamaoag)


## Badges

[![MIT License](https://img.shields.io/github/license/mamaoag/binosearch?style=for-the-badge)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSE)

[![Issues](https://img.shields.io/github/issues/mamaoag/binosearch?style=for-the-badge)](https://github/mamaoag/binosearch/issues)

## Installation

Install Binosearch by cloning the project.

```bash
  git clone https://github.com/mamaoag/binosearch.git
  cd binosearch
```


## Usage


You can run Binosearch by running the following command:
```bash
  go run main.go
``` 

Or if you will provide your own wordlist,
```bash
  go run main.go path/of/wordlist
``` 

Follow the interactive commandline and it will run the scanner.


## Package


If you prefer to use the packages in this repository you would only need to use the following:

```bash
go get https://github.com/mamaoag/binosearch/services/scanner # Scanner for Wordlist 
go get https://github.com/mamaoag/binosearch/services/owasp # Contains OWASP Checklist 
```


## Contributing

Contributions are always welcome!

1. Create first an issue describing what can be improved / bug.

2. If the issue is tagged as good first issue or help wanted, You may then create a pull request linking to the issue.
