# BookStore with Go

![stack_logo](./misc/display.png)


### Prerequisite
- Install [MySQL](https://hub.docker.com/_/mysql)


### Installing Dependencies
```
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
$ go get -u github.com/gorilla/mux
```

### Build Application

```
$ cd cmd/main && go build .
$ ./main
```

If you're using Windows just type `main.exe` and press enter.
