[![Visits Badge](https://badges.strrl.dev/visits/mncred/play)](https://badges.strrl.dev)![GitHub Tag](https://img.shields.io/github/v/tag/mncred/play)[![codebeat badge](https://codebeat.co/badges/0a36aca9-0bc9-4cac-9a33-c973a9dcf10c)](https://codebeat.co/projects/github-com-mncred-play-main)![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mncred/play)

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

# ![](https://api.iconify.design/mdi/rocket-launch.svg?height=32&color=royalblue)MNC Play

> Minecraft launcher for mnc.red project servers

## ![](https://api.iconify.design/mdi/code.svg?height=24&color=royalblue)For developers

> Under the MIT license, you can use and modify this code to suit your needs.

### ![](https://api.iconify.design/mdi/crane.svg?height=20&color=royalblue)Technology stack

The launcher works like a native application using[GoLang](https://go.dev)server as BackEnd and[Yeah. js](https://vuejs.org)as a FrontEnd in conjunction with[Vuetify](https://vuetifyjs.com)MaterialDesign library
components. Assembly using[Wails](https://wails.io).

<details>
    <summary>Поддерживаемые архитектуры Wails</summary>

-   `windows/amd64`- Windows x64
-   `windows/386`- Windows x32
-   `windows/arm64`- Windows ARM
-   `darwin/amd64`- MacOS Intel
-   `darwin/arm64`- MacOS M1
-   `darwin/universal`- MacOS Intel и M1
-   `linux/amd64`- Linux x64
-   `linux/arm`- This client cannot be assembled for this architecture
-   `linux/arm64`- This client cannot be assembled for this architecture

</details>

### ![](https://api.iconify.design/mdi/application-brackets.svg?height=20&color=royalblue)Assembly instructions

#### ![](https://api.iconify.design/mdi/world.svg?height=16&color=royalblue)Setting up the environment

In general, you will need to install`GoLang`,`Node`and`Wails`.

Optionally it is possible to use`UPX`for compressing builds. However,`UPX`does not support all available`Wails`architecture
and is limited only`windows/amd64`,`windows/386`,`linux/amd64`.

Depending on the platform on which the assembly takes place, you may need to install additional components,
such as`GTK`and`WebKit2 GTK`.
See more details in[workflows/build.yml](.github/workflows/build.yaml)and in[Wails documentation](https://wails.io/docs/gettingstarted/installation#platform-specific-dependencies).

#### ![](https://api.iconify.design/mdi/package-variant-closed.svg?height=16&color=royalblue)Assembly

##### ![](https://api.iconify.design/mdi/gear.svg?height=14&color=royalblue)Using Make

Required to install`git`, because assembly via`Make`supplies information about the current commit to the application.

For assembly under the platform`windows/amd64`do:  
`make build_windows_amd64`.

Find the finished build in`build/bin`.
See more details about the build options for individual platforms in[Makefile](Makefile).

##### ![](https://api.iconify.design/mdi/car-manual-transmission.svg?height=14&color=royalblue)Manual assembly

Instead of using`make`It is possible to build directly through Wails, for example:

`wails build --webview2 embed --platform windows/amd64`

* * *

<details>
    <summary>
        Черновик
    </summary>

A fully customizable Minecraft launcher and mod-server.
Powered by[Wails](https://wails.io), and written with Go, TypeScript, Vue, Quasar

ALL:

-   [ ] How to install GoLang
-   [ ] How to install Node
-   [ ] How to install Wails
-   [ ] HOW TO: Play on Android (via Pojav Launcher + modpack from Modrinth)

Get a list of releases by API:`curl -XGET 'https://api.github.com/repos/mncred/play/releases'`

## Dependencies

### User

For more info see[Wails Platform Specific Dependencies](https://wails.io/docs/gettingstarted/installation#platform-specific-dependencies).
Long story short, Windows users need to install[WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/).

### Developer

#### Windows

-   WebView2 Runtime
-   Go 1.18+
-   Node 15+

#### Linux

-   gcc
-   Let him cry
-   libwebkit
-   Go 1.18+
-   Node 15+

#### MacOS

-   Xcode
-   Go 1.18+
-   Node 15+

## ALL

-   [ ] Implement support for command line arguments for the launcher. Launcher can be used
    with their help not only for the game, but also as a CLI tool for generating/validating data.
-   [ ] Implement a JRE/JDK loading module for the selected platform.
    Ultimately, we will receive unpacked archive files (folded into`${launcher}/java/11/...`)
    There are several vendors providing a portable version:
    -   `https://api.adoptium.net/v3/assets/latest/${major}/hotspot?vendor=eclipse`
        -   major - JVM version (8/11) and we receive JSON with information where to download for each platform.
            Filter the list by condition`a.filter(v => v.binary.architecture === 'x32').filter(v => v.binary.os === 'windows').filter(v => v.binary.image_type === 'jdk')`And download from`.binary.package.link`
    -   `https://corretto.aws/downloads/latest/amazon-corretto-${major}-${arch}-${sanitizedOS}-jdk.${ext}`Here we directly get the download link
        -   major - JVM version (8/11)
        -   arch - aarch64 (arm64) or x64 (amd64)
        -   sanitizedOS - windows, macos, linux
        -   ext - Each platform has its own packaging extension. windows - zip, macos and linux - tar.gz
            Using the same variables we get a link to md5:`https://corretto.aws/downloads/latest_checksum/amazon-corretto-${major}-${arch}-${sanitizedOS}-jdk.${ext}`
-   [ ] Implement mrpack loading module with Modrinth.
      It uses the format`.mrpack`for modpacks. For example:<https://cdn.modrinth.com/data/1KVo5zza/versions/Hrdee8Qh/Fabulously.Optimized-5.8.0-beta.10.mrpack>This is a zip archive and inside is`modrinth.index.json`which contains links from where to download and where to put mods/resource packs, etc.
      and there may also be files with configs that are specific to this pack.
-   [ ] Implement a Minecraft launcher using JDK and mrpack

The distribution structure for the launcher may look like this.
The bootloader can thus boot both the client and the server.
Add support for packages from MC launcher.

The loader puts different types of resources into the cache directory when necessary and searches there first.
When you launch a specific version of the game, the client goes online and checks the integrity of the files.
In the settings, we can allow or prohibit (by default we allow) the launch of the game offline if the integrity could not be verified.

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
