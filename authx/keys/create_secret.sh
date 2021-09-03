k create secret generic niche-srv-keys --from-file=ssh-privatekey=<path>/keys/id_rsa --from-file=ssh-publickey=<path>/keys/id_rsa.pub -n niche-dev
