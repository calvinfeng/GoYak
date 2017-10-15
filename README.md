# GoYak

This is a demo MVC server built with Go. It's basically a real time messaging application. The purpose is to show that 
Rails MVC framework can be built in Go using `Gorm` which is a Go ORM package for various SQL database. 

## Setup

### Dependency Management with `dep`
First of all, get `dep` for dependency management

```
go get -u github.com/golang/dep/cmd/dep
```

If you are using Mac OS X then it's even easier

```
brew install dep 
brew upgrade dep
```

I fucking love Homebrew on Mac. It has everything!

### Databse
I am going to use PostgreSQL for this project, so let's create one. The superuser on my computer is `cfeng` so I will
use that to create a database named `goyak_development`

```
$ psql postgres
postgres=# create database goyak_development;
```

Then just exit with `\q`

Actually just in case you don't remember the password to your `ROLE`, do the following

```
postgres=# alter user <your_username> with password <whatever you like>
```

I did mine with
```
postgres=# alter user cfeng with password "cfeng";
```



