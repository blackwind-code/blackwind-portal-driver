# blackwind-portal-driver
Blackwind service portal, backend driver

# Installation
## Build
```bash
sudo apt install python3-openstackclient
git clone https://github.com/blackwind-code/blackwind-portal-driver.git
cd blackwind-portal-driver/blackwind-portal-driver
go build main.go
```

## Add database user & grant privilege to user on Openstack database (Once)
- Requires fully operational openstack cluster
- Requires fully operational zerotier controller node
```bash
mysql -u root -p

CREATE USER '<OS_DB_USERNAME>'@'localhost' IDENTIFIED BY '<OS_DB_PASSWORD>';
CREATE USER '<OS_DB_USERNAME>'@'%' IDENTIFIED BY '<OS_DB_PASSWORD>';

GRANT SELECT ON keystone.local_user TO '<OS_DB_USERNAME>'@'localhost';
GRANT SELECT ON keystone.local_user TO '<OS_DB_USERNAME>'@'%';
GRANT SELECT ON keystone.password TO '<OS_DB_USERNAME>'@'localhost';
GRANT SELECT ON keystone.password TO '<OS_DB_USERNAME>'@'%';
GRANT UPDATE ON keystone.password TO '<OS_DB_USERNAME>'@'localhost';
GRANT UPDATE ON keystone.password TO '<OS_DB_USERNAME>'@'%';

FLUSH PRIVILEGES;
```
# Run
- Requires following environment variables to be set
```bash
# secret password for backend-driver handshake
export SECRET=<secret-password>

# https://docs.openstack.org/keystone/yoga/install/keystone-install-ubuntu.html
export OS_USERNAME=<openstack-user-username>
export OS_PASSWORD=<openstack-user-password>
export OS_PROJECT_NAME=<openstack-project-name>
export OS_USER_DOMAIN_NAME=<openstack-user-domain-name>
export OS_PROJECT_DOMAIN_NAME=<openstack-project-domain-name>
export OS_AUTH_URL=http://localhost:5000/v3
export OS_IDENTITY_API_VERSION=3

export OS_DB_USERNAME=<OS_DB_USERNAME>
export OS_DB_PASSWORD=<OS_DB_PASSWORD>
export OS_DB_IP=<openstack-database-endpoint-ip>

# Zerotier API variables
export ZEROTIER_API_URL=http://localhost:9993
export ZEROTIER_TOKEN=<zerotier-controller-token>
export ZEROTIER_NODE_ID=<zerotier-controller-node-id>
export ZEROTIER_NETWORK_ID=<zerotier-network-id>

# Run
cd blackwind-portal-driver/cmd/blackwind-portal-driver
./main
```
