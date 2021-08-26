# hardware-detection

This sub-project is responsible for creating an iso-image, which advertises information about the host-system.
It should
- boot fast
- be super small
- be extensible (might want more information in the future)
- be rebuildable via docker container

## MML
- small size (<50MB)
- boots fast (<30s)
- extensible
- not rebuildable (isos dont work)
Doesn't work (-> doesn't boot). More info in ./MML/README.md

## Alpine
- medium size (<500MB)
- boots fast
- semi extensible (only found out how to install apks, but even that doesn't work)
- not rebuildable (isos are 4x larger than original, and don't work)
Doesn't work (-y doesn't boot). More info in ./Alpine/README.md

## Ubuntu
- large size (1.3GB)
- boots fast (<40s)
- extensible (full shell before imaged)
- rebuildable (but not with container; requires manual work)
Built from a virtual ubuntu machine.
Works, but is large. More info in ./Ubuntu/README.md

## Debian
- medium size (<1GB)
- boots super fast (<10s)
- extensible (super customizable to be exact)
- rebuildable via docker container
