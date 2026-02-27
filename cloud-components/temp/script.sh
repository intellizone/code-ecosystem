#!/bin/sh
su git

sleep 1

mkdir -p /code-ecosystem/cloud-deploy.git
chown git:git /code-ecosystem/cloud-deploy.git

cd /code-ecosystem/cloud-deploy.git
git init --bare ~/cloud-deploy.git

mkdir -p /code-ecosystem/.ssh

cat > /code-ecosystem/.ssh/authorized_keys <<EOF
public key content here
EOF

chown git:git -R /code-ecosystem/.ssh

