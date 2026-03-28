# Nginx

This is the current nginx configuration for this website


```shell
server {
    listen 80;
    server_name docs.kshyst.ir doc.kshyst.ir;

        location / {
        root /var/www/docs;
        index index.html;
        try_files $uri $uri/ =404;
        }
        #return 301 https://$host$request_uri;
}

server {
        listen 443 ssl;
        server_name docs.kshyst.ir doc.kshyst.ir;

        ssl_certificate /etc/ssl/certs/local_nginx.crt;
        ssl_certificate_key /etc/ssl/private/local_nginx.key;

        location / {

        root /var/www/docs;
        index index.html;

        try_files $uri $uri/ =404;
    }
}

```

Each server listens on the specified port and only accepts requests
from valid `server_name`s.

A root directory for all folders and html files is defined in `root` variable and the `index` as the index page too.  

This configuration file is under:

```shell
sudo nano /etc/nginx/sites-available/docs
```

And there should be a symlink (soft link) of this in `sites-enabled`:

```shell
sudo ln -s /etc/nginx/sites-available/docs /etc/nginx/sites-enabled/
```

Checking if the configurations are valid and no syntax error occurs:

```shell
sudo nginx -t
```

And finally restart the nginx service:

```shell
sudo systemctl reload nginx
```

Also the directory you define for web files should be accessible by the `www-data` group:

```shell
sudo chown -R www-data:www-data /var/www/my-docs
sudo chmod -R 755 /var/www/my-docs
```