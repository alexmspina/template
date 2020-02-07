#!/usr/bin/fish

for port in 80 443
    set node_port (kubectl get service -n ingress-nginx ingress-nginx -o=jsonpath="{.spec.ports[?(@.port == $port)].nodePort}")

    docker run -d --name kind-proxy-$port \
      --publish 192.168.1.197:$port:$port \
      --link kind-control-plane:target \
      alpine/socat -dd \
      tcp-listen:$port,fork,reuseaddr tcp-connect:target:$node_port
end