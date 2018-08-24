#!/usr/bin/env bash

# Get latest Anchor and Identity Registry Addresses from contract json
export TEST_TIMEOUT=${TEST_TIMEOUT:-600s}
export TEST_TARGET_ENVIRONMENT=${TEST_TARGET_ENVIRONMENT:-'local'}
export CENT_CENTRIFUGENETWORK=${CENT_CENTRIFUGENETWORK:-'testing'}

## Making Env Var Name dynamic
cent_upper_network=`echo $CENT_CENTRIFUGENETWORK | awk '{print toupper($0)}'`
temp1="CENT_NETWORKS_${cent_upper_network}_CONTRACTADDRESSES_ANCHORREGISTRY"
printf -v $temp1 `cat $CENT_ETHEREUM_CONTRACTS_DIR/deployments/$TEST_TARGET_ENVIRONMENT.json | jq -r '.contracts.AnchorRegistry.address' | tr -d '\n'`
temp2="CENT_NETWORKS_${cent_upper_network}_CONTRACTADDRESSES_IDENTITYFACTORY"
printf -v $temp2 `cat $CENT_ETHEREUM_CONTRACTS_DIR/deployments/$TEST_TARGET_ENVIRONMENT.json | jq -r '.contracts.IdentityFactory.address' | tr -d '\n'`
temp3="CENT_NETWORKS_${cent_upper_network}_CONTRACTADDRESSES_IDENTITYREGISTRY"
printf -v $temp3 `cat $CENT_ETHEREUM_CONTRACTS_DIR/deployments/$TEST_TARGET_ENVIRONMENT.json | jq -r '.contracts.IdentityRegistry.address' | tr -d '\n'`
export $temp1
export $temp2
export $temp3
vtemp1=$(eval "echo \"\$$temp1\"")
vtemp2=$(eval "echo \"\$$temp2\"")
vtemp3=$(eval "echo \"\$$temp3\"")
#

echo "ANCHOR REGISTRY ADDRESS: ${vtemp1}"
echo "IDENTITY REGISTRY ADDRESS: ${vtemp3}"
echo "IDENTITY FACTORY ADDRESS: ${vtemp2}"

if [ -z $vtemp1 ] || [ -z $vtemp2 ] || [ -z $vtemp3 ]; then
    echo "One of the required contract addresses is not set. Aborting."
    exit -1
fi