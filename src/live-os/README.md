# custom debian live

## How to generate the custom iso

Run `./generate.sh`, wait, enjoy the custom iso.

## How to customize the live os within the iso

During `./generate.sh`, the script `./container/chroot.sh` is executed in the livesystem.
This means, by editing the contents of that script, one can fully customize the live os.

## Notes

lighttpd config location: /etc/lighttpd/lighttpd.conf
lighttpd web-root: /var/www/html/

The iso could be further minified if other run-options are removed. Might not be feasable though, since it sounds like a lot work and the image is less than 1GB already.
