#!/usr/bin/env bash
set -a

# Setup
local_dir="$(dirname "$0")"
PARENT_DIR=`pwd`
MIGRATE=false
if [ "${SKIP_MIGRATION}" == "true" ];then
  MIGRATE=false
fi

args=( "$@" )
if [ $# == 0 ]; then
  args=(  unit cmd testworld integration )
  MIGRATE=true
elif [ $# == 1 ] && [ "${args[0]}" == "unit" ]; then
  MIGRATE=false
else
  MIGRATE=true
fi

GETH_DOCKER_CONTAINER_NAME="geth-node"
CC_DOCKER_CONTAINER_NAME="cc-node"
GETH_DOCKER_CONTAINER_WAS_RUNNING=`docker ps -a --filter "name=${GETH_DOCKER_CONTAINER_NAME}" --filter "status=running" --quiet`
CC_DOCKER_CONTAINER_WAS_RUNNING=`docker ps -a --filter "name=${CC_DOCKER_CONTAINER_NAME}" --filter "status=running" --quiet`

# Code coverage is stored in coverage.txt
echo "" > coverage.txt

################# Run Dependencies #########################
if $MIGRATE; then
  for path in ${local_dir}/test-dependencies/test-*; do
      [ -d "${path}" ] || continue # if not a directory, skip
      source "${path}/env_vars.sh" # Every dependency should have env_vars.sh + run.sh executable files
      echo "Executing [${path}/run.sh]"
      ${path}/run.sh
      if [ $? -ne 0 ]; then
          exit 1
      fi
  done
  ############################################################

  ################# Migrate contracts ########################
  rm -f /tmp/migration.log
  "${PARENT_DIR}"/build/scripts/migrate.sh &> /tmp/migration.log
  if [ $? -ne 0 ]; then
    echo "migrations failed"
    cat /tmp/migration.log
    exit 1
  fi
  rm -f /tmp/migration.log
  ## adding this env here as well since the envs from previous step(child script) is not imported
  export MIGRATION_RAN=true

  ############################################################

  ################# deploy bridge########################
  ## delete any stale bridge containers
  docker rm -f bridge
  path=${local_dir}/test-dependencies/bridge
  source "${path}/env_vars.sh"
  echo "Executing [${path}/run.sh]"
  ${path}/run.sh
  if [ $? -ne 0 ]; then
      exit 1
  fi
  ############################################################
fi

################# Run Tests ################################
if [[ ${status} -eq 0 ]]; then
  statusAux=0
  for path in ${local_dir}/tests/*; do
    [[ -x "${path}" ]] || continue # if not an executable, skip

    for arg in "${args[@]}"; do
        if [[ ${path} == *$arg* ]]; then
            echo "Executing test suite [${path}]"
            ./$path
            statusAux="$(( $statusAux | $? ))"
            continue
        fi
    done
  done
  # Store status of tests
  status=$statusAux
fi
############################################################

if $MIGRATE; then
  ################# CleanUp ##################################
  if [ -n "${GETH_DOCKER_CONTAINER_WAS_RUNNING}" ]; then
      echo "Container ${GETH_DOCKER_CONTAINER_NAME} was already running before the test setup. Not tearing it down as the assumption is that the container was started outside this context."
  else
      echo "Bringing GETH Daemon Down"
      docker rm -f geth-node
  fi

  if [ -n "${CC_DOCKER_CONTAINER_WAS_RUNNING}" ]; then
      echo "Container ${CC_DOCKER_CONTAINER_NAME} was already running before the test setup. Not tearing it down as the assumption is that the container was started outside this context."
  else
      echo "Bringing Centrifuge Chain down"
      docker rm -f cc-node
  fi

  echo "Bringing bridge down..."
  docker rm -f bridge

  ############################################################
fi

################# Propagate test status ####################
echo "The test suite overall is exiting with status [$status]"
exit $status
############################################################
