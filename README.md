# DEC (Development Environment Creator)
Create python or go development environment.
dec create python app or de create go app

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
dec python create --name emeka -e emeka@gmail.com -d /home/emeka/app
```
* create golang app
```
dec golang create --name emeka -e emeka@gmail.com -d /home/emeka/app
```
