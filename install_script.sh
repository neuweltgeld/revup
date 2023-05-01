#!/bin/sh

echo ""
echo "__________.__                 __                                                                 "
echo "\______   \  |   ____   ____ |  | __  _________________    ____  ________________    ____  ____  "
echo " |    |  _/  |  /  _ \_/ ___\|  |/ / /  ___/\____ \__  \ _/ ___\/ __ \_  __ \__  \ _/ ___\/ __ \ "
echo " |    |   \  |_(  <_> )  \___|    <  \___ \ |  |_> > __ \\  \__\  ___/|  | \// __ \\  \__\  ___/ "
echo " |______  /____/\____/ \___  >__|_ \/____  >|   __(____  /\___  >___  >__|  (____  /\___  >___  >"
echo "        \/                 \/     \/     \/ |__|       \/     \/    \/           \/     \/    \/ "
echo "__________             ____ ___                                                                  "
echo "\______   \ _______  _|    |   \______                                                           "
echo " |       _// __ \  \/ /    |   /\____ \                                                          "
echo " |    |   \  ___/\   /|    |  / |  |_> >                                                         "
echo " |____|_  /\___  >\_/ |______/  |   __/                                                          "
echo "        \/     \/               |__|                                                             "
echo " Discord irlandali_turist#7300 // Geralt"
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "
echo "                                                                                                 "

sleep 1

function addToPath {
  source $HOME/.bash_profile
  PATH_EXIST=$(grep ${1} $HOME/.bash_profile)
  if [ -z "$PATH_EXIST" ]; then
    echo "export PATH=$PATH:${1}" >>$HOME/.bash_profile
  fi
}


echo "1. Updating packages..." && sleep 1
sudo apt update

echo "2. Installing dependencies..." && sleep 1
sudo apt install curl tar wget clang pkg-config libssl-dev jq build-essential bsdmainutils git make ncdu gcc git jq chrony liblz4-tool -y && curl https://get.ignite.com/cli! | bash

echo "3. Installing go..." && sleep 1
if ! [ -x "$(command -v go)" ]; then
  source <(curl -s "https://raw.githubusercontent.com/nodejumper-org/cosmos-scripts/master/utils/go_install.sh")
  source .bash_profile
fi

echo "$(go version)"

echo "4. Building binaries..." && sleep 1

cd $HOME && git clone https://github.com/neuweltgeld/revup && cd revup

VALIDATOR_NAME=validator1
CHAIN_ID=revup
KEY_NAME=revup-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

NAMESPACE_ID=$(openssl rand -hex 8)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://rpc-blockspacerace.pops.one/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

ignite chain build
revupd tendermint unsafe-reset-all
revupd init $VALIDATOR_NAME --chain-id $CHAIN_ID

revupd keys add $KEY_NAME --keyring-backend test
revupd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
revupd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
revupd collect-gentxs
revupd start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT
