# User Manual

For those going to deploy this project themselves

## Build from source and run

First start by cloning this project to your `GOPATH`:

```shell
$ mkdir -p $GOPATH/src/github.com/qzyse2017/bleBird
$ cd $GOPATH/src/github.com/qzyse2017/bleBird
$ git clone https://github.com/qzyse2017/bleBird.git .
```

run 

```shell
$ ./build/build-all.sh
```


## Configure the database
set your configuration for database in a file and pass the filename as a parameter while starting the program, e.g. you may use `server_conf.yaml` and run

```shell
$ dist/server server_conf.yaml
```

replace the server_conf.yaml with with file you have written configuration in, and you need to provide the necessary infomation as `server_conf.yaml` has written.
//TODO -- add some more details

you could also replace a new database yourself, you just need to add Go's driver of this database to `datastore`(import it). It should work fine. All the database operations in this project is based on `database/sql` package.

