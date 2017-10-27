#!/bin/sh

cd /etc/dovecot/auth/
for dominio in *; do
    if ! [ -d /var/mail/vhosts/${dominio} ]; then
      mkdir -p /var/mail/vhosts/${dominio}
    fi
    cp -rf $dominio/passwd  /var/mail/vhosts/$dominio/
    cp -rf $dominio/shadow  /var/mail/vhosts/$dominio/
    cp -rf $dominio/autoreply.sieve  /var/mail/vhosts/$dominio/
done
chown -R 101:104 /var/mail/vhosts/

/usr/sbin/rsyslogd
/usr/sbin/dovecot -F