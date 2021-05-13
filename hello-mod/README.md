# My example software
Brief explanation of what the software is.

## Dependencies
* Java 8 or later
* Memcached
...

The following environment variables must be set when running:
* `JAVA_HOME
...

## How to configure
Edit `conf/dev-config.properties` if you want to change the Memcached port, etc,

## How to run locally
1. Make sure Memcached is running
2. Run `mvn jetty:run`
3. Point your browser at `http://localhost:8080/healthcheck.do`

## How to run tests
1. Make sure Memcached is running
2. Run `mvn test`

## How to build
Run `mvn package` to create WAR file in the `target` folder.

## How to release/deploy
Run `./release.sh` and follow the prompts.