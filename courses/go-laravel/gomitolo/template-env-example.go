package gomitolo

var envExampleTemplate = `
# Give your application a unique name (no spaces)
APP_NAME=myapp

# false for production, true for development
DEBUG=true

# the port should we listen on
PORT=4000

# the server name, e.g, www.mysite.com
SERVER_NAME=localhost

# should we use https?
SECURE=false

# database config - postgres or mysql
DATABASE_TYPE=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5433
DATABASE_USER=postgres
DATABASE_PASS=postgres
DATABASE_NAME=gomitolo
DATABASE_SSL_MODE=disable

# redis config
REDIS_HOST="localhost:6380"
REDIS_PASSWORD=
REDIS_PREFIX=gomitolo

# cache (currently only redis)
CACHE=redis

# cooking seetings
COOKIE_NAME=gomitolo
COOKIE_LIFETIME=1440
COOKIE_PERSIST=true
COOKIE_SECURE=false
COOKIE_DOMAIN=localhost

# session store: cookie, redis, mysql, or postgres
SESSION_TYPE=redis

# mail settings
SMTP_HOST=localhost
SMTP_USERNAME=535314fc4423b2
SMTP_PASSWORD=ffa2ce4e252d97
SMTP_PORT=1025
SMTP_ENCRYPTION=none
SMTP_FROM=gomitolo@gomitolo.com

# mail settings for api services TODO

# template engine: go or jet
# RENDERER=jet
RENDERER=go

# the encryption key; must be exactly 32 characters long
KEY=rHbaqmfdhmdrDDPIytYhwSRzcvpOesjZ`
