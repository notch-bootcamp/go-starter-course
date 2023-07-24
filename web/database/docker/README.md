This README file defines process of running a PostgreSQL using docker.

* STEP 1: install docker following steps defined here: https://docs.docker.com/engine/install/
* STEP2: inside this directory build a docker image from a Dockerfile using following line:
    ```
  docker build -t psql-devcademy .```
* STEP3: run a created docker image using following line: 
    ```
  docker run --name psql-devcademy-container -p 5432:5432 -d psql-devcademy```

You can use `docker logs CONTAINER_ID` for checking the status of database.
Use `psql` or some other tools (e.g. IDE) to connect to the database. Following is connection URL:
```jdbc:postgresql://docker@mypassword:localhost:5432/docker```

You can use `docker container stop CONTAINER_ID` to stop container, and a `docker container rm CONTAINER_ID` to remove the container from docker instance.

When this docker container started it automatically contains table called `positions` which we use in examples.
