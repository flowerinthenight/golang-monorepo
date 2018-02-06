#!/bin/bash

docker tag ${NAME}:${CIRCLE_SHA1} 963826138034.dkr.ecr.ap-northeast-1.amazonaws.com/${NAME}:${CIRCLE_SHA1}
docker images
`aws ecr get-login --no-include-email --region ap-northeast-1`
docker push 963826138034.dkr.ecr.ap-northeast-1.amazonaws.com/${NAME}:${CIRCLE_SHA1}

# Install kubectl.
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
kubectl -h

# Setup kubectl config.
mkdir ${HOME}/.kube
cp kubeconf-dev.yaml ${HOME}/.kube/config
ls -laF
