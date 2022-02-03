# pkg
pkg create-python-app or pkg create-go-app

## Contributing
* application is designed using [cobra](https://www.linode.com/docs/guides/using-cobra/)

### Create new commands
* command
```
cobra add <command_name>
```
* example
```
cobra add create
```

### create subcommands
* subcommand
```
cobra add <subcommand_name> -p 'mainCommand_name'
```
* example
```
cobra add deactivate -p 'pythonCmd'
```

## How to use
* create python app
```
pkg create python --name emeka -e emeka@gmail.com -d /home/emeka/app
```
* create golang app
```
pkg create golang --name emeka -e emeka@gmail.com -d /home/emeka/app
```
