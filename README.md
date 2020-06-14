

# Go Serve
Simple go http file server

### Install

Use ``go get`` &`` go install`` to install go-serve
```
go get github.com/ssubedir/go-serve
go install github.com/ssubedir/go-serve
```

### Usage

go-serve usage flags
```
go-serve [-path= directory] [-tls= TLS] [-port= port] [-rt= read timeout] [-wt= write timeout] [-it= idle timeout]
```
#### Example

Serves ```/www/```  directory
```
go-serve
```
Serves ```/public/```  directory
```
go-serve -path=public
```
Serves ```/public/```  directory on port 2000
```
go-serve -path=public -port=2000
```

## Built With

* [GO](https://golang.org/) - Programming language


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/ssubedir/go-serve/blob/master/LICENSE) file for details