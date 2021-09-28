#!/bin/sh

NGINX_DNS_RESOLVER="$(grep '^nameserver' /etc/resolv.conf | head -1 | awk '{print $2;}')"

export NGINX_DNS_RESOLVER

# replace all references to the DNS resolver in the template file
sed -i -e "s/\$NGINX_DNS_RESOLVER/${NGINX_DNS_RESOLVER}/g" \
  -e "s/\${NGINX_DNS_RESOLVER}/${NGINX_DNS_RESOLVER}/g" \
  /etc/nginx/conf.d/default.conf