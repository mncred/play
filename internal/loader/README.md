# Loader

This is a files loader for Minecraft servers and clients.
Config example:

```yaml
files:
    forge:41.1:
        place:
            type: url
            
        hash: # Optional section, may include all or any of supported hashing algorithms
            md5: ... # Remote file md5
            sha1: ... # Remote file SHA1
        size: 12345 # Recommended. Remote file size in bytes
        location: mods/forge.jar # where to place this file
        files: # Optional. Recursive section with files required by general file.
            name: ...
        requires: # Optional. Provides dependencies for file. It's like files section, but resolves dependencies collision.
            - minecraft:1.20.1
```

Kubernetes like minecraft server/client configuration

```yaml
api: mncplay/v1
kind: Server
metadata:
    name: mncred-stone-server
    region: [ ru ]
spec:
    title: MNC.RED :: Stone
    motd: http://play.mnc.red
    
---
```