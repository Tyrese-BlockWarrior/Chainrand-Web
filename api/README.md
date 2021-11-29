# Chainrand-api

This api syncs the data between the Chainrand smart contract and a MySQL database for fast approximate name search.  

We use https://www.alchemy.com/ as our provider to connect to the Mumbai testnet via websockets.

# Requirements

- go1.16 (this is the version we used, but you can try other versions)

# Installation

This is if you want to set it up on your own server.

- Create an `.env` file and fill in the configuration. For example:

  ```
  DB_USER="chainrand"
  DB_PASS="dwkirwq2vqmsfft6o98cnq4pgctrzhlq"
  DB_HOST="tcp(127.0.0.1:3306)"
  DB_NAME="chainrand"

  WSS_ENDPOINT_4="wss://68e69fddfa54bc77b079a8a8d9b256c6@rinkeby.infura.io/ws/v3/7092fea6dae8b4e1467682bedf69c665"
  CONTRACT_ADDRESS_4=0x4ddfa9040279f94fcd76ef1d7b06deaefe9e25c0

  WSS_ENDPOINT_80001="wss://polygon-mumbai.g.alchemy.com/v2/KjTc1Q6rCUfUA6FKZDaJMbiyHOYt2n2T"
  CONTRACT_ADDRESS_80001=0x13fd5a9c917c15d37e4E00EFF2d3fae87485822F

  CHAIN_IDS=4,80001
  ```

- Run the `schema.sql` on the MySQL database to create the tables.

- Install the go dependencies.

  ```
  go get -d ./...
  ```

- Build the executable.
  ```
  go build
  ```

- Run the executable. If you want to run it as a background service on Ubuntu, look at `build.sh`.

# Note

- The code is configured to connect to port `8081`. See `server/main.go` if you want to change the port.  

  If you use nginx, add the following to your nginx config file (may be in `/etc/nginx/nginx.conf`) under your server config block to do a reverse proxy to your go app.

  ```
  location ~ ^/api(/|$) {
    proxy_pass http://localhost:8081;
  }
  ```

- The adapter class for the Chainrand smart contract is in `cc` directory. It is generated with `abigen`.

  Take a look at `abigen.sh` to see the exact command.

# License

MIT
