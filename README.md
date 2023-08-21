# MNCRED :: Play

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
