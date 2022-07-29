<p align="center">
 <a style="text-decoration:none" href="https://img.shields.io/github/workflow/status/emylincon/dec/Go?style=for-the-badge" target="_blank">
     <img src="https://img.shields.io/github/workflow/status/emylincon/dec/Go?style=for-the-badge" alt="Build Status" />
 </a>
 <a style="text-decoration:none" href="https://img.shields.io/badge/Version-1.4.5-informational?style=for-the-badge" target="_blank">
     <img src="https://img.shields.io/badge/Version-0.0.3-informational?style=for-the-badge" alt="Version: 0.0.3" />
 </a>
 <a style="text-decoration:none" href="https://img.shields.io/github/license/emylincon/dec?style=for-the-badge" target="_blank">
     <img src="https://img.shields.io/github/license/emylincon/dec?color=green&style=for-the-badge" alt="Version: 1.4.5" />
 </a>
 <a style="text-decoration:none" href="https://img.shields.io/github/languages/count/emylincon/dec?style=for-the-badge" target="_blank">
 <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/emylincon/dec?style=for-the-badge">
 </a>
 <a style="text-decoration:none" href="https://godoc.org/github.com/emylincon/dec" target="_blank">
     <img src="https://img.shields.io/badge/Go-Doc-informational?style=for-the-badge" alt="Documentation" />
 </a>
 </p>

# DEC (Development Environment Creator)
Create python or go development environment.
dec create python app or de create go app

## Contributing ![report](https://goreportcard.com/badge/github.com/emylincon/dec)
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

## Autocomplete
You now need to ensure that the `dec` completion script gets sourced in all your shell sessions.

### Bash
To do so in all your `bash` shell sessions.
```bash
echo 'source <(dec completion bash)' >>~/.bashrc
```

### Zsh
To do so in all your `zsh` shell sessions.
```zsh
echo 'source <(dec completion zsh)' >>~/.zshrc
```

### Fish
To do so in your current `fish` shell session.
```bash
dec completion fish | source
```

### Powershell
To do so in all your `powershell` shell sessions, add the following line to your `$PROFILE` file:

```powershell
dec completion powershell | Out-String | Invoke-Expression
```

This command will regenerate the auto-completion script on every `powerShell` start up. You can also add the generated script directly to your `$PROFILE` file.

To add the generated script to your `$PROFILE` file, run the following line in your `powershell` prompt:
```powershell
dec completion powershell >> $PROFILE
```

[Help: problem with completion?](https://github.com/kubernetes-sigs/kind/issues/522)
