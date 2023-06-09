.PHONY:	build

build:
	go build -o chainbridge

# chainbridge commands (Build must precede)
SRC_GATEWAY = https://api.baobab.klaytn.net:8651
# DST_GATEWAY = https://rpc-mumbai.maticvigil.com/

SRC_ADDR = 0x2345BF77D1De9eacf66FE81a09a86CfAb212a542
SRC_PK = 56a7a32def4a15c0187e8a076baf4c76db60bbf9ad59b8396d0b2174f9d5533e
# DST_ADDR = 0x48179c14aeDf2bfaCf7bF6328A2eDACd6d3DA426
# DST_PK = 23e44e5019b54dd4e8c02db9cf4a6e4563c6fe32352bcede43cded05fac05404

RESOURCE_ID = 0x0000000000000000000000000000000000000000000000000000000000000001

SRC_BRIDGE = 0x0A03be286E8947acb0d29Ee06BC1003cf8F7f3fC
SRC_TOKEN = 0x0b118820e8dfA0E14E9F2093d72215bc81365DFa
SRC_HANDLER = 0x04D3B69FEcdcCAb9A2F1C930d05134d1b3a27d99
SRC_GENERIC = 0xf25f351c5efb349897dB8A2D8B448684B1d3F9f2


# For BERA Testnet
DST_ADDR = 0xF17aeD3A3e090c096beded279B39b051223A4838
DST_PK = 2b6399ecbebd49c4fdf3ddfab539b306f1f007dfeb1dd1e9fdcd375ccc04f788

DST_GATEWAY = http://121.141.157.230:22001
DST_BRIDGE = 0x02f325776004560fe1e1ddD1965a871e0dD45a87
DST_TOKEN = 0x80cD41f4c8c0d42C6645215BB2f7fDb5c7E3E3be
DST_HANDLER = 0xCE72d7266bA59637d53bA49fE1F7b4f9d1093F07
DST_GENERIC = 0xd9590eea4272387CfeeE9A605B33C8e8ed2be6d4

# DST_BRIDGE = 0x4e17d27605896f145E8fD95BcEee63315148eFed
# DST_TOKEN = 0x63A604721b426794aD9F8E47736383cb8D54A1E0
# DST_HANDLER = 0xe433e545536B675E5222C471C1380e0F5E5E768e
# DST_GENERIC = 0xd46f3d8cc7e88Aa2AD6022a333d2332FacccA02b

# changeable arguments
FEE = 1000000000000000
AMOUNT = 1000000000
APPROVE = 10000000000000000000

# src 체인에 Bridge, ERC20, ERC20 Handler Contract 배포
deploysrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 deploy \
	--bridge --erc20 --erc20-handler --generic-handler \
	--relayers ${SRC_ADDR} \
	--relayer-threshold 1 \
	--domain 0 \
	--fee ${FEE} \
	--erc20-name BERSSWAP \
	--erc20-symbol BSP

# dst 체인에 Bridge, ERC20, ERC20 Handler Contract 배포
deploydst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 deploy \
	--bridge --erc20 --erc20-handler  --generic-handler \
	--relayers ${SRC_ADDR} \
	--relayer-threshold 1 \
	--domain 1 \
	--fee ${FEE} \
	--erc20-name BERASWAP \
	--erc20-symbol BAP

# src 체인에 deposit 하기 전 필요한 과정들을 차례대로 수행하도록 커맨드를 결합시켰다.
readysrc: registersrc burnablesrc mintablesrc approvesrc mintsrc

# dst 체인에 deposit 하기 전 필요한 과정들을 차례대로 수행하도록 커맨드를 결합시켰다.
readydst: registerdst burnabledst mintabledst approvedst mintdst

# src 체인의 Bridge에 Resource ID를 키로 Handler를, Handler에 Token을 등록한다.
registersrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 bridge register-resource \
    --bridge ${SRC_BRIDGE} \
    --handler ${SRC_HANDLER} \
    --resource ${RESOURCE_ID} \
    --target ${SRC_TOKEN}

# dst 체인에 Bridge에 Resource ID를 키로 Handler를, Handler에 Token을 등록한다.
registerdst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 bridge register-resource \
    --bridge ${DST_BRIDGE} \
    --handler ${DST_HANDLER} \
    --resource ${RESOURCE_ID} \
    --target ${DST_TOKEN}

# src 체인의 전송자 앞으로 APPROVE만큼 토큰을 민팅한다.
mintsrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 mint \
	--contract ${SRC_TOKEN} \
	--amount ${APPROVE} \
	--decimals 0 \
	--recipient ${SRC_ADDR}

# dst 체인의 전송자 앞으로 APPROVE만큼 토큰을 민팅한다.
mintdst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 erc20 mint \
	--contract ${DST_TOKEN} \
	--amount ${APPROVE} \
	--decimals 0 \
	--recipient ${DST_ADDR}

# src 체인의 Bridge에 등록된 Handler에게 Token을 소각할 권리를 부여한다.
burnablesrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 bridge set-burn \
    --bridge ${SRC_BRIDGE} \
    --handler ${SRC_HANDLER} \
    --token-contract ${SRC_TOKEN}

# dst 체인의 Bridge에 등록된 Handler에게 Token을 소각할 권리를 부여한다.
burnabledst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 bridge set-burn \
    --bridge ${DST_BRIDGE} \
    --handler ${DST_HANDLER} \
    --token-contract ${DST_TOKEN}

# src 체인의 Bridge에 등록된 Handler에게 Token을 민팅할 권리를 부여한다.
mintablesrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 add-minter \
    --minter ${SRC_HANDLER} \
    --contract ${SRC_TOKEN}

# dst 체인의 Bridge에 등록된 Handler에게 Token을 민팅할 권리를 부여한다.
mintabledst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 erc20 add-minter \
    --minter ${DST_HANDLER} \
    --contract ${DST_TOKEN}

# src 체인의 Handler가 민팅할 수 있는 최대 수량을 지정한다.
approvesrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 approve \
    --amount ${APPROVE} \
    --contract ${SRC_TOKEN} \
    --recipient ${SRC_HANDLER} \
	--decimals 0

# dst 체인의 Handler가 민팅할 수 있는 최대 수량을 지정한다.
approvedst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 erc20 approve \
    --amount ${APPROVE} \
    --contract ${DST_TOKEN} \
    --recipient ${DST_HANDLER} \
	--decimals 0

# bridge 컨트랙트에 recipient 앞으로 amount 만큼 토큰을 예치한다. 두 체인에 배포되어있는 bridge 컨트랙트에 의해 토큰이 교환된다.
depositsrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 deposit \
	--amount ${AMOUNT} \
	--domain 1 \
	--bridge ${SRC_BRIDGE} \
	--recipient ${SRC_ADDR} \
	--resource ${RESOURCE_ID} \
	--decimals 0 \
	--value ${FEE}

depositdst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 erc20 deposit \
	--amount ${AMOUNT} \
	--domain 0 \
	--bridge ${DST_BRIDGE} \
	--recipient ${DST_ADDR} \
	--resource ${RESOURCE_ID} \
	--decimals 0 \
	--value ${FEE}

checkbalancesrc:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 balance \
	--address ${SRC_ADDR} \
	--contract ${SRC_TOKEN}

checkbalancedst:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 30000000000 erc20 balance \
	--address ${DST_ADDR} \
	--contract ${DST_TOKEN}

# src 체인의 전송자가 보유한 dst 체인의 토큰 수량을 조회한다.
balancesrc:
	./chainbridge evm-cli --url ${DST_GATEWAY} --private-key ${DST_PK} --gas-price 10000000000 erc20 balance \
	--address ${SRC_ADDR} \
	--contract ${DST_TOKEN}

# dst 체인의 전송자가 보유한 src 체인의 토큰 수량을 조회한다.
balancedst:
	./chainbridge evm-cli --url ${SRC_GATEWAY} --private-key ${SRC_PK} --gas-price 30000000000 erc20 balance \
	--address ${DST_ADDR} \
	--contract ${SRC_TOKEN}
