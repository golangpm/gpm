#!/bin/bash

sudo mkdir -p /opt/gpm-conf

cat <<EOF | sudo tee /opt/gpm-conf/gpm-config.json
{
 "username": "",
 "email": ""
}
EOF

sudo chmod 777 /opt/gpm-conf
sudo chmod 777 /opt/gpm-conf/gpm-config.json


# sudo mkdir -p ~/gpm-conf

# cat <<EOF | sudo tee ~/gpm-conf/gpm-config.json
# {
#  "username": "",
#  "email": ""
# }
# EOF

# sudo chmod 777 ~/gpm-conf
# sudo chmod 777 ~/gpm-conf/gpm-config.json