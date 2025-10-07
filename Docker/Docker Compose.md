# Docker Compose

## Example of docker compose file

```yaml
version: '3.8'
services:
  database:
    image: postgres:latest
    container_name: blog-db
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin123
      POSTGRES_DB: blog_db
    volumes:
      - blog-db-data:/var/lib/postgresql/data
    networks:
      - blog-net
    ports:
      - "5432:5432"

  backend:
    image: python:3.10
    container_name: blog-backend
    working_dir: /app
    volumes:
      - ./backend:/app
    environment:
      DB_HOST: database
      DB_USER: admin
      DB_PASSWORD: admin123
      DB_NAME: blog_db
    networks:
      - blog-net
    ports:
      - "5000:5000"
    command: sh -c "pip install -r requirements.txt && python app.py"

  frontend:
    image: nginx:latest
    container_name: blog-frontend
    volumes:
      - ./frontend:/usr/share/nginx/html:ro
    networks:
      - blog-net
    ports:
      - "8080:80"

volumes:
  blog-db-data:

networks:
  blog-net:
```

## Commands

### up

```shell
docker compose up
```