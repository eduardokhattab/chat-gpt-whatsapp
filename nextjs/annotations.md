## Run Dev Container

- docker 23.0.4
- docker-compose v2.17.2
- maybe you need to recreate the .docker/ folder to share volume ("permission denied")

## Docker hosts

To connect the nextjs containers with chatservice containers, we need to integrate their hosts/network. To do that, we have two options:

<details>
<summary>1. Create a extra_host</summary>

The IP `172.17.0.1` is the default bridge address from docker. He allows us to connect the containers with our host machine

You can check this address running the command from the container or the host machine:

```sh
ip route

docker network inspect bridge
```

```yml
# /home/gympasser/projects/chat-gpt-whatsapp/nextjs/docker-compose.yaml
services:
  nextjs:
    extra_hosts:
      - "host.docker.internal:172.17.0.1"
```
This will create a new host called `host.docker.internal` in container pointing to a fix IP address `172.17.0.1`. You can check it running a `cat /etc/hosts` inside your container

Now we need to map this new host in our local machine, so add it with

```sh
sudo vim /etc/hosts

# Docker custom host
127.0.0.1       host.docker.internal
```

And then, everytime that the container try to access `host.docker.internal`, it will use our localhost (127.0.0.1) that will be running the docker-compose port-forwarding, in our case the ports (8081, 50052)

```yml
# /home/gympasser/projects/chat-gpt-whatsapp/chatservice/docker-compose.yaml
services:
  chatservice:
    ports:
      - "8081:8080"
      - "50052:50051"
```
```ts
// /home/gympasser/projects/chat-gpt-whatsapp/nextjs/src/grpc/client.ts
export const chatClient = new proto.pb.ChatService(
    "host.docker.internal:50052",
    grpc.credentials.createInsecure()
); 
```
</details>

<details>
<summary> 2. Using a docker network </summary>

- create a common network to integrate the containers of the project

```sh
docker network create chatservice-network
docker network ls
```

- map the custom network to all containers

```yml
# /home/gympasser/projects/chat-gpt-whatsapp/nextjs/docker-compose.yaml
version: '3'

services:
  nextjs:
    networks:
      - chatservice-network

  db:
    networks:
      - chatservice-network

networks:
  chatservice-network:
    external: true

# /home/gympasser/projects/chat-gpt-whatsapp/chatservice/docker-compose.yaml
version: '3'

services:
  chatservice:
    networks:
      - chatservice-network

  mysql:
    networks:
      - chatservice-network

networks:
  chatservice-network:
    external: true
```

</details>

## Keycloak

- Keycloak(localhost:9000) > Admin Console > Users > Add user > Credentials
- Clients > Create > ClientID (nextjs) > Valid URIs (http://localhost:3000/*) > Web origins (http://localhost:3000/*)
- Check Client Authentication (now you can go to Credentials tab) > Client Secret

