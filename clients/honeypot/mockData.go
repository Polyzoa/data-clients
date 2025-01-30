package honeypot

const ResponseData = `{
    "token": {
        "name": "Tether USD",
        "symbol": "USDT",
        "decimals": 6,
        "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
        "totalHolders": 6354736
    },
    "withToken": {
        "name": "Wrapped Ether",
        "symbol": "WETH",
        "decimals": 18,
        "address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
        "totalHolders": 948647
    },
    "summary": {
        "risk": "low",
        "riskLevel": 1,
        "flags": []
    },
    "simulationSuccess": true,
    "honeypotResult": {
        "isHoneypot": false
    },
    "simulationResult": {
        "buyTax": 0,
        "sellTax": 0,
        "transferTax": 0,
        "buyGas": "176752",
        "sellGas": "137663"
    },
    "flags": [],
    "contractCode": {
        "openSource": true,
        "rootOpenSource": true,
        "isProxy": false,
        "hasProxyCalls": false
    },
    "chain": {
        "id": "1",
        "name": "Ethereum",
        "shortName": "eth",
        "currency": "ETH"
    },
    "router": "0xE592427A0AEce92De3Edee1F18E0157C05861564",
    "pair": {
        "pair": {
            "name": "Uniswap V3: WETH-USDT",
            "address": "0x4e68Ccd3E89f51C3074ca5072bbAC773960dFa36",
            "token0": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
            "token1": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
            "type": "UniswapV3"
        },
        "chainId": "1",
        "reserves0": "22674475147151207423007",
        "reserves1": "47864183967113",
        "liquidity": 119074507.63794056,
        "router": "0xE592427A0AEce92De3Edee1F18E0157C05861564",
        "createdAtTimestamp": "1620232628",
        "creationTxHash": "0x2e07c690f149223e4f290986277304ea6a05c6ee47ba303732166bc1b15cbafb"
    },
    "pairAddress": "0x4e68Ccd3E89f51C3074ca5072bbAC773960dFa36"
}`

const ResponseError = `{
    "code": 404,
    "error": "No pairs found"
}`

const WeirdoResponseData = `{
    "token": {
        "name": "CandleAI",
        "symbol": "CNDL",
        "decimals": 18,
        "address": "0x6EFb32bc7893b793603E39643D86594CE3638157",
        "totalHolders": 1212,
        "airdropSummary": {
            "totalTxs": 1,
            "totalAmountWei": "33900000000000000000000000",
            "totalTransfers": 34
        }
    },
    "withToken": {
        "name": "Wrapped Ether",
        "symbol": "WETH",
        "decimals": 18,
        "address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
        "totalHolders": 951092
    },
    "summary": {
        "risk": "low",
        "riskLevel": 1,
        "flags": []
    },
    "simulationSuccess": true,
    "honeypotResult": {
        "isHoneypot": false
    },
    "simulationResult": {
        "maxBuy": {
            "token": 2000000,
            "tokenWei": "1999999999999999999994650",
            "withToken": 7.327062258687804,
            "withTokenWei": "7327062258687804391"
        },
        "buyTax": 4.999999999999993,
        "sellTax": 4.979118170850721,
        "transferTax": 0,
        "buyGas": "160900",
        "sellGas": "177373"
    },
    "holderAnalysis": {
        "holders": "676",
        "successful": "676",
        "failed": "0",
        "siphoned": "0",
        "averageTax": 5.142810650887551,
        "averageGas": 164508.4571005917,
        "highestTax": 7.19,
        "highTaxWallets": "0",
        "taxDistribution": [
            {
                "tax": 6,
                "count": 36
            },
            {
                "tax": 5,
                "count": 170
            },
            {
                "tax": 4,
                "count": 467
            },
            {
                "tax": 7,
                "count": 3
            }
        ],
        "snipersFailed": 0,
        "snipersSuccess": 0
    },
    "flags": [],
    "contractCode": {
        "openSource": true,
        "rootOpenSource": true,
        "isProxy": false,
        "hasProxyCalls": true
    },
    "chain": {
        "id": "1",
        "name": "Ethereum",
        "shortName": "eth",
        "currency": "ETH"
    },
    "router": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
    "pair": {
        "pair": {
            "name": "Uniswap V2: CNDL-WETH",
            "address": "0xbeda23eeE41163F686ca875F0974E789362FDed5",
            "token0": "0x6EFb32bc7893b793603E39643D86594CE3638157",
            "token1": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
            "type": "UniswapV2"
        },
        "chainId": "1",
        "reserves0": "7833287471064078480920804",
        "reserves1": "21306318945945053581",
        "liquidity": 141615.43175887625,
        "router": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
        "createdAtTimestamp": "1715100863",
        "creationTxHash": "0x85634377c07ca0f9583ea2c36a454f389fe12130edfce1da22058a8dfe87e330"
    },
    "pairAddress": "0xbeda23eeE41163F686ca875F0974E789362FDed5"
}`
