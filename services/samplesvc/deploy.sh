#!/bin/bash

echo "Running '$0' for '${NAME}'"

### The following section will push your container image to ECR. The `$NAME` variable is provided from our
### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
###
# docker tag ${NAME}:${CIRCLE_SHA1} 963826138034.dkr.ecr.ap-northeast-1.amazonaws.com/${NAME}:${CIRCLE_SHA1}
# docker images
# `aws ecr get-login --no-include-email --region ap-northeast-1`
# docker push 963826138034.dkr.ecr.ap-northeast-1.amazonaws.com/${NAME}:${CIRCLE_SHA1}

### If you need to deploy your service to mochi, you need to use the kubectl tool. The setup for kubectl's
### config file is also included in this section.
###
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
mkdir ${HOME}/.kube
cp kubeconf-dev.yaml ${HOME}/.kube/config

if [[ "$CIRCLE_BRANCH" == "production" ]]; then
  echo "Todo (production)"
else
  echo "Config setup for kubectl (branch = ${CIRCLE_BRANCH})"
  kubectl config set clusters.mochi.k8s.local.certificate-authority-data $KUBE_CLUSTER_CERT
  kubectl config set clusters.mochi.k8s.local.server $KUBE_SERVER
  kubectl config set users.mochi.k8s.local.client-certificate-data $KUBE_CLIENT_CERT
  kubectl config set users.mochi.k8s.local.client-key-data $KUBE_CLIENT_KEYDATA
  kubectl version
fi
