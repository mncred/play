[![Visits Badge](https://badges.strrl.dev/visits/mncred/play)](https://badges.strrl.dev)
![GitHub Tag](https://img.shields.io/github/v/tag/mncred/play)
[![codebeat badge](https://codebeat.co/badges/0a36aca9-0bc9-4cac-9a33-c973a9dcf10c)](https://codebeat.co/projects/github-com-mncred-play-main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mncred/play)

<div align="center">
    <table>
        <tr>
            <td align="center" colspan="3">
                <b>Загрузите последнюю версию MNCPlay</b>
            </td>
        </tr>
        <tr>
            <td align="center">Windows</td>
            <td align="center">MacOS</td>
            <td align="center">Linux</td>
        </tr>
        <tr>
            <td>
                <a href="https://github.com/mncred/play/releases/download/latest/mncplay-windows-amd64.exe" target="_blank">
                    <img width="100" src="https://raw.githubusercontent.com/mncred/play/main/.github/assets/logo-windows.png"></img>
                </a>
            </td>
            <td>
                <a href="https://github.com/mncred/play/releases/download/latest/mncplay-darwin-universal.zip" target="_blank">
                    <img width="100" src="https://raw.githubusercontent.com/mncred/play/main/.github/assets/logo-apple.png"></img>
                </a>
            </td>
            <td>
                <a href="https://github.com/mncred/play/releases/download/latest/mncplay-linux-amd64" target="_blank">
                    <img width="100" src="https://raw.githubusercontent.com/mncred/play/main/.github/assets/logo-linux.png"></img>
                </a>
            </td>
        </tr>
        <tr>
            <td align="center" colspan="3">
                <a href="https://github.com/mncred/play/releases" target="_blank">все релизы и версии</a>
            </td>
        </tr>
    </table>
</div>

# ![](https://api.iconify.design/mdi/rocket-launch.svg?height=32&color=royalblue) MNCPlay

> Minecraft лаунчер для серверов проекта mnc.red

## ![](https://api.iconify.design/mdi/code.svg?height=24&color=royalblue) Для разработчиков

> В соответствии с лицензией MIT вы можете использовать и модифицировать этот код под ваши нужды.

### ![](https://api.iconify.design/mdi/crane.svg?height=20&color=royalblue) Стек технологий

Лаунчер работает как нативное приложение, используя [GoLang](https://go.dev) сервер в качестве BackEnd и
[Vue.js](https://vuejs.org) в качестве FrontEnd в связке с [Vuetify](https://vuetifyjs.com) MaterialDesign библиотекой
компонентов. Сборка при помощи [Wails](https://wails.io).

<details>
    <summary>Поддерживаемые архитектуры Wails</summary>

- `windows/amd64` - Windows x64
- `windows/386` - Windows x32
- `windows/arm64` - Windows ARM
- `darwin/amd64` - MacOS Intel
- `darwin/arm64` - MacOS M1
- `darwin/universal` - MacOS Intel и M1
- `linux/amd64` - Linux x64
- `linux/arm` - Данный клиент не удаётся собрать под эту архитектуру
- `linux/arm64` - Данный клиент не удаётся собрать под эту архитектуру

</details>

### ![](https://api.iconify.design/mdi/application-brackets.svg?height=20&color=royalblue) Инструкция по сборке

#### ![](https://api.iconify.design/mdi/world.svg?height=16&color=royalblue) Настройка окружения

В общем случае вам потребуется установить `GoLang`, `Node` и `Wails`.

Опционально возможно использовать `UPX` для сжатия билдов. Однако, `UPX` не поддерживает все доступные `Wails` архитектуры
и ограничивается только `windows/amd64`, `windows/386`, `linux/amd64`.

В зависимости от платформы на которой происходит сборка может потребоваться установить дополнительные компоненты,
такие как `GTK` и `WebKit2 GTK`.
Подробнее смотрите в [workflows/build.yml](.github/workflows/build.yaml)
и в [документации Wails](https://wails.io/docs/gettingstarted/installation#platform-specific-dependencies).

#### ![](https://api.iconify.design/mdi/package-variant-closed.svg?height=16&color=royalblue) Сборка

##### ![](https://api.iconify.design/mdi/gear.svg?height=14&color=royalblue) Используя Make

Требуется установить `git`, т.к. сборка через `Make` поставляет информацию о текущем коммите в приложение.

Для сборки под плаформу `windows/amd64` выполните:  
`make build_windows_amd64`.

Готовый билд ищите в `build/bin`.
Смотрите подробнее возможности сборки под отдельные платформы в [Makefile](Makefile).

##### ![](https://api.iconify.design/mdi/car-manual-transmission.svg?height=14&color=royalblue) Ручная сборка

Вместо использования `make` возможно осуществлять сборку напрямую через Wails, например:

`wails build --webview2 embed --platform windows/amd64`

---

<details>
    <summary>
        Черновик
    </summary>

A fully customizable Minecraft launcher and mod-server.
Powered by [Wails](https://wails.io), and written with Go, TypeScript, Vue, Quasar

TODO:

- [ ] How to install GoLang
- [ ] How to install Node
- [ ] How to install Wails
- [ ] HOWTO: Play on Android (via Pojav Launcher + mcpack from Modrinth)

Получить список релизов по API: `curl -XGET 'https://api.github.com/repos/mncred/play/releases'`

## Dependencies

### User

For more info see [Wails Platform Specific Dependencies](https://wails.io/docs/gettingstarted/installation#platform-specific-dependencies).
Long story short, Windows users need to install [WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/).

### Developer

#### Windows

- WebView2 Runtime
- Go 1.18+
- Node 15+

#### Linux

- gcc
- libgtk3
- libwebkit
- Go 1.18+
- Node 15+

#### MacOS

- Xcode
- Go 1.18+
- Node 15+


## TODO

- [ ] Реализовать поддержку аргументов командной строки для лаунчера. Лаунчер может использоваться
  с их помощью не только для игры, но и как CLI инструмент формирования/валидации данных.
- [ ] Реализовать модуль загрузки JRE/JDK под выбранную платформу.
  В конечном итоге мы получим распакованные файлы архива (сложить в `${launcher}/java/11/...`)
  Есть несколько вендоров поставляющих портативную версию:
  - `https://api.adoptium.net/v3/assets/latest/${major}/hotspot?vendor=eclipse`
    - major - версия JVM (8/11) и мы получаем JSON с информацией где скачать под каждую платформу.
    Отфильтровать список по условию `a.filter(v => v.binary.architecture === 'x32').filter(v => v.binary.os === 'windows').filter(v => v.binary.image_type === 'jdk')`
    И скачать из `.binary.package.link`
  - `https://corretto.aws/downloads/latest/amazon-corretto-${major}-${arch}-${sanitizedOS}-jdk.${ext}`
    Здесь напрямую получаем ссылку на скачивание
    - major - версия JVM (8/11)
    - arch - aarch64 (arm64) или x64 (amd64)
    - sanitizedOS - windows, macos, linux
    - ext - У каждой платформы своё расширение упаковки. windows - zip, macos и linux - tar.gz
    Используя те же переменные получаем ссылку на md5:
    `https://corretto.aws/downloads/latest_checksum/amazon-corretto-${major}-${arch}-${sanitizedOS}-jdk.${ext}`
- [ ] Реализовать модуль загрузки mrpack с Modrinth.
    Он использует формат `.mrpack` для модпаков. Например:
    https://cdn.modrinth.com/data/1KVo5zza/versions/Hrdee8Qh/Fabulously.Optimized-5.8.0-beta.10.mrpack
    Это zip архив и внутри лежит `modrinth.index.json` в котором указаны ссылки откуда качать и куда класть моды/ресурспаки и прочее,
    а так же могут лежать файлы с конфигами, которые специфичны для этого пака.
- [ ] Реализовать модуль запуска Minecraft используя JDK и mrpack

Структура дистрибутива для лаунчера может выглядеть следующим образом.
Загрузчик таким образом может загрузить как клиент, так и сервер.
Добавить поддержку пакетов из MC лаунчера.

Загрузчик складывает разные виды ресурсов в директорию кэша при необходимости и ищет для начала в ней.
При запуске конкретной версии игры клиент лезет в сеть и проверяет целостность файлов.
В настройках мы можем разрешить или запретить (по умолчанию разрешаем) запуск игры оффлайн, если не удалось проверить целостность.

```yaml
---
kind: modpack/mrpack/v1
# Идентификатор объекта для использования в конфигурации
id: fabulously-optimized
spec:
    # Локализованное название модпака
    title: Fabulously Optimized
    # Локализованное описание
    desc: Improve your graphics and performance with this simple modpack. 1.20.4 beta!
    # Ссылки откуда скачивать .mrpack, будет загружен по первой доступной
    download:
        - link: https://cdn.modrinth.com/data/1KVo5zza/versions/Hrdee8Qh/Fabulously.Optimized-5.8.0-beta.10.mrpack
---
kind: modpack/asset/v1
id: single-mod
spec:
    title:
    desc:
    path: mods/my-mod.jar
    download:
        - link: https://...
          md5: ...
          sha1: ...
---
kind: modpack/asset/v1
id: motd
spec:
    title:
    desc:
    path: server.properties
    content: |
        # Конфиг сервера или клиента, или биндинги клавиш
---
kind: modpack/v1
id: optimizations-client
spec:
    title: Оптимизация
    desc: Включает наборы оптимизаций
    mrpacks:
        - fabulously-optimized
    resourcepacks:
        - ...
    shaders:
        - ...
    datapacks:
        - ...
    assets:
        - single-mod
        - motd
---
kind: java/v1
id: java-8
spec:
    title:
    desc:
    platforms:
        'windows/amd64':
            download:
                - https://...
        'macos/amd64':
            download:
                - https://...
---
kind: minecraft/v1
id: minecraft-1.20.1
spec:
    title:
    desc:
    # Здесь разрешена ссылка на forge/paper/spigot и что угодно
    download:
        - https://...
---
kind: client/v1
id: client-stone
spec:
    title:
    desc:
    defaultProps:
        ram:
            max: 4G
    java: java-8
    minecraft: minecraft-1.20.1
    modpacks:
        - optimizations-client
    run: -Xmx{{.props.ram.max}} ... -jar {{.minecraft.jar}}
```

</details>
