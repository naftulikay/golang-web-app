FROM nginx:alpine

ENV API_PROTO=http API_HOST=api API_PORT=8080
ENV NGINX_ENVSUBST_TEMPLATE_SUFFIX=".template"

COPY --chown=nginx:nginx www/* /var/www/

COPY --chown=root:root etc/nginx/conf.d/default.conf.template /etc/nginx/templates/default.conf.template
COPY --chown=root:root lib/docker/90-detect-resolver.sh /docker-entrypoint.d/

EXPOSE 80