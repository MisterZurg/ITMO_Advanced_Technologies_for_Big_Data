# K3S deploy

1. Install k3s: `curl -sfL https://get.k3s.io | sh -` ([original](https://k3s.io/))
2. [Install HELM](https://www.corpsoft24.ru/knowledge/k3s-ustanovka-stateful-prilozheniya-s-bazoy-dannykh/)
3. [Install Kompose](https://kompose.io/installation/)
4. `cd k3s`
5. `kompose convert --file ../docker-compose.yml --volumes hostPath --out all` (if you want to get separate config files for each service and deployment, use `kompose convert --file ../docker-compose.yml --volumes hostPath`)
6. Now we have a file called `all` in `k3s` directory. The only things left are deployment and forwarding.
7. `sudo kubectl proxy --port=<PORT>` Here instead of `<PORT>` we need port 18080 for the API and ports 9000 and 8123 for the database. Running a new terminal for each of those commands is neccessary.
8. `sudo kubectl apply -f all`

Testing with `curl localhost:8080/ping` works.

TODO:
- test if logs are written to the database
- maybe add replication to the server
