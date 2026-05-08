# Nginx Configuration

The main configuration file is located at `/etc/nginx/nginx.conf`.

The configuration files are based on **Context**.

```shell
events {
    # Handles connection processing (e.g., how many connections per worker)
}

http {
    # Handles all HTTP web traffic
    
    server {
        # Defines a specific website/virtual host
        
        location / {
            # Defines how to handle specific URLs/paths for this website
        }
    }
}

```

Instead of putting everything in `nginx.conf`, you create individual files for each website in `/etc/nginx/sites-available/` and link them to `/etc/nginx/sites-enabled/`.

## Usages

### Reverse Proxy

You build an app using Node.js, Python, or Java running on a local port (like localhost:3000). You don’t want to expose port 3000 directly to the internet. Instead, NGINX acts as a middleman (Reverse Proxy).

User visits port 80
→
NGINX intercepts
→
Forwards to port 3000 secretly.

```shell
server {
    listen 80;
    server_name myapp.com;

    location / {
        proxy_pass http://127.0.0.1:3000; # Forward traffic here
        
        # These headers pass the user's real IP and info to your backend app
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}

```

### Serving Static Files

Like the example in `Get Started.md` and how this website is made.

1. Put your website files in a directory, e.g., `/var/www/mywebsite/`.
2. Create a configuration file for your site: `/etc/nginx/sites-available/mywebsite`
3. Add the following configuration:

```shell
server {
    listen 80;                     # The port to listen on (80 is default for HTTP)
    server_name mywebsite.com;     # Your domain name or IP address

    root /var/www/mywebsite;       # Where your files live
    index index.html;              # The default file to load

    location / {
        try_files $uri $uri/ =404; # Try to find the file, if not found, return 404 error
    }
}
```

### Load Balancing

```shell
# Define the group of servers
upstream my_backend {
    server 127.0.0.1:3001;
    server 127.0.0.1:3002;
    server 127.0.0.1:3003;
}

server {
    listen 80;
    server_name scaleapp.com;

    location / {
        proxy_pass http://my_backend; # Pass traffic to the upstream group
    }
}
```

## Catch All

A Catch all configuration is when we want to send empty response for every request coming to our port 80 and 443 except when the domain is 
what we set in another configuration:

Write a config `/etc/nginx/sites-available/catch-all` and dont forget to link it to enabled

```text
# Catch-all for HTTP (Port 80)
server {
   listen 80 default_server;
   server_name _;
   return 444; 
}

# Catch-all for HTTPS (Port 443)
server {
   listen 443 ssl default_server;
   server_name _;
   
   # This tells Nginx to drop the connection before even 
   # trying to show a certificate (Requires Nginx 1.19.4 or newer)
   ssl_reject_handshake on;
}
```

```shell
   sudo nginx -t
   sudo systemctl restart nginx
```

For this to work, make sure there isnt any other config that uses `default_server`  in their `listen` part. 

```shell
grep -r "default_server" /etc/nginx/sites-enabled/*
```