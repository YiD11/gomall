# Gomall

### Introduction

This is a project from a Bytedance course-based project [biz-demo](https://github.com/cloudwego/biz-demo). I fellow the course on Bilibili during coding ([Link Here](https://www.bilibili.com/video/BV1yr421F7z4?spm_id_from=333.788.videopod.sections&vd_source=0a78546755087ad89e6c4ec0c5de39e7)).

Just aim to learn more by conducting a project.

The main code is in `./app` folder, including frontend, user, product, cart, payment, order, checkout and email services.

The project mainly uses Hertz and Kitex frameworks in [Cloudwego community](https://www.cloudwego.io/). The Hertz hanndles HTTP requests and provides router layer to routing and navigation requests, and use RPC protocol to call backend function. The Kitex receive the RPC calling from frontend, CRUD(Create, Read, Update and Delete) corresponding data that the request need in database.

Service Discovery, Log Collection, Metrics Monitoring and Distributed Tracing are inherited into this project😍.

### Run

#### Docker-compose

To run the project on docker compose, you can run `cluster.docker-compose.yml` and `app.docker-compose.yml` sequentially using Docker in Linux.

``` shell
# run basic infrastructures, including mysql, redis, nats ...
docker compose -f cluster.docker-compose.yml up
# run apps
docker compose -f app.docker-compose.yml up
```

The `docker-compose.yml` is integration of `cluster.docker-compose.yml` and `app.docker-compose.yml`. You can edit the config to build those app images directly before running.


### Future Work

Now the project obviously are too weak to be deployed in the real scenario 😥. There are huge features can be added in the project. 


