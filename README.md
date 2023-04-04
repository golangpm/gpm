# `GPM` - golang pocket manager

![logo](docs/imgs/gpm-dirty_2.jpg)

### Пакетный менеджер для создания приложений на языке golang

### Установка

```bash
git clone https://github.com/golangpm/gpm
cd gpm

#Linux
./gpm_linux-setup

# Mac Os
./gpm_mac_setup

# set your configuration
# GitHub username
# Email address
gpm set-config

# Show your configuration
gpm config

# to start application
gpm

# to create a new project
gpm goapp MyApp

```

## Как использовать

[Documentation](./docs/DOCUMENTATION.md)

`gpm` - меню утилиты

`gpm config` - посмотреть конфиг

`gpm set-config` - установить данные пользователя в конфигурацию (github username и email)

`gpm goapp` **MyApp** - создать приложение "MyApp"

[TODO](todo.md)