package dexscreener

import (
	"fmt"
)

const ResponseError = `{
    "schemaVersion": "1.0.0",
    "pairs": null
}`

var PairsDataResponse = fmt.Sprintf(`{"schema": "test", "pairs": %v}`, PairsData)

const PairsData = `[
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0xfadc475639131c1eac3655c37eda430851d53716",
            "pairAddress": "0xFadc475639131C1EAC3655c37EDA430851d53716",
            "labels": [
                "v1"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xA1077a294dDE1B09bB078844df40758a5D0f9a27",
                "name": "Wrapped Pulse",
                "symbol": "WPLS"
            },
            "priceNative": "53.8628",
            "priceUsd": "0.002995",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 5
                },
                "h1": {
                    "buys": 10,
                    "sells": 43
                },
                "h6": {
                    "buys": 59,
                    "sells": 128
                },
                "h24": {
                    "buys": 732,
                    "sells": 567
                }
            },
            "volume": {
                "h24": 63015.83,
                "h6": 6796.09,
                "h1": 1022.89,
                "m5": 5.51
            },
            "priceChange": {
                "m5": 0.59,
                "h1": 1.53,
                "h6": 3.8,
                "h24": 3.28
            },
            "liquidity": {
                "usd": 129047.85,
                "base": 21537407,
                "quote": 1160066663
            },
            "fdv": 108700778,
            "marketCap": 108700778,
            "pairCreatedAt": 1683944565000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x6444456960c3f95b5b408f4d9e00220643f06f94",
            "pairAddress": "0x6444456960C3f95b5b408f4d9E00220643f06F94",
            "labels": [
                "v1"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xA1077a294dDE1B09bB078844df40758a5D0f9a27",
                "name": "Wrapped Pulse",
                "symbol": "WPLS"
            },
            "priceNative": "103.3912",
            "priceUsd": "0.005750",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 9
                },
                "h1": {
                    "buys": 7,
                    "sells": 103
                },
                "h6": {
                    "buys": 177,
                    "sells": 186
                },
                "h24": {
                    "buys": 895,
                    "sells": 1114
                }
            },
            "volume": {
                "h24": 111436.48,
                "h6": 15214.5,
                "h1": 4917.63,
                "m5": 25.42
            },
            "priceChange": {
                "m5": 0.59,
                "h1": 2.01,
                "h6": 3.66,
                "h24": 6.38
            },
            "liquidity": {
                "usd": 380951.41,
                "base": 33122122,
                "quote": 3424536109
            },
            "fdv": 163766482,
            "marketCap": 163766482,
            "pairCreatedAt": 1683944275000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x5777d92f208679db4b9778590fa3cab3ac9e2168",
            "pairAddress": "0x5777d92f208679DB4b9778590Fa3CAB3aC9e2168",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "1.0001158",
            "priceUsd": "0.005709",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 3
                },
                "h1": {
                    "buys": 1,
                    "sells": 42
                },
                "h6": {
                    "buys": 21,
                    "sells": 117
                },
                "h24": {
                    "buys": 85,
                    "sells": 526
                }
            },
            "volume": {
                "h24": 129225.84,
                "h6": 18440.66,
                "h1": 5739.65,
                "m5": 1248
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.73,
                "h6": 4,
                "h24": 4.92
            },
            "liquidity": {
                "usd": 621856.12,
                "base": 22888199,
                "quote": 86039221
            },
            "fdv": 162590809,
            "marketCap": 162590809,
            "pairCreatedAt": 1683945295000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x2db5ef4e8a7dbe195defae2d9b79948096a03274",
            "pairAddress": "0x2dB5EF4E8A7Dbe195defAe2d9b79948096a03274",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.9971",
            "priceUsd": "0.005659",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 2
                },
                "h6": {
                    "buys": 11,
                    "sells": 9
                },
                "h24": {
                    "buys": 50,
                    "sells": 33
                }
            },
            "volume": {
                "h24": 30678.79,
                "h6": 4571.71,
                "h1": 3613.81,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.12,
                "h6": 3.06,
                "h24": 2.87
            },
            "liquidity": {
                "usd": 791824.57,
                "base": 69956509,
                "quote": 69755730
            },
            "fdv": 161166368,
            "marketCap": 161166368,
            "pairCreatedAt": 1685651445000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0xc636bfe0bae34824380b4e26bc34e4614e55e483",
            "pairAddress": "0xC636bfE0bAe34824380B4E26Bc34e4614e55e483",
            "labels": [
                "v1"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xCc78A0acDF847A2C1714D2A925bB4477df5d48a6",
                "name": "Atropa",
                "symbol": "ATROPA"
            },
            "priceNative": "0.1376",
            "priceUsd": "0.005701",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 15,
                    "sells": 4
                },
                "h6": {
                    "buys": 30,
                    "sells": 4
                },
                "h24": {
                    "buys": 204,
                    "sells": 27
                }
            },
            "volume": {
                "h24": 32773.89,
                "h6": 1479.6,
                "h1": 799.87,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 2.09,
                "h6": 3.97,
                "h24": 5.33
            },
            "liquidity": {
                "usd": 18035381.75,
                "base": 1581569985,
                "quote": 217712197
            },
            "fdv": 162371762,
            "marketCap": 162371762,
            "pairCreatedAt": 1685325445000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0xdf102c06b6ed69250efef5a5748dd1075e0e2809",
            "pairAddress": "0xDf102c06b6ed69250EfeF5a5748dD1075E0E2809",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xd6c31bA0754C4383A41c0e9DF042C62b5e918f6d",
                "name": "BEAR",
                "symbol": "TEDDY BEAR ㉾"
            },
            "priceNative": "155255.02163",
            "priceUsd": "0.002830",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 0
                },
                "h6": {
                    "buys": 4,
                    "sells": 10
                },
                "h24": {
                    "buys": 147,
                    "sells": 85
                }
            },
            "volume": {
                "h24": 15826.86,
                "h6": 246.95,
                "h1": 23.39,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0.09,
                "h6": -3.11,
                "h24": -0.62
            },
            "liquidity": {
                "usd": 109282.37,
                "base": 19303819,
                "quote": 2997014915426
            },
            "fdv": 102702774,
            "marketCap": 102702774,
            "pairCreatedAt": 1694659095000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x99ac8ca7087fa4a2a1fb6357269965a2014abc35",
            "pairAddress": "0x99ac8cA7087fA4A2A1FB6357269965A2014ABc35",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
                "name": "Wrapped BTC",
                "symbol": "WBTC"
            },
            "priceNative": "0.00002126",
            "priceUsd": "0.005684",
            "txns": {
                "m5": {
                    "buys": 3,
                    "sells": 0
                },
                "h1": {
                    "buys": 10,
                    "sells": 0
                },
                "h6": {
                    "buys": 39,
                    "sells": 22
                },
                "h24": {
                    "buys": 183,
                    "sells": 101
                }
            },
            "volume": {
                "h24": 71839.14,
                "h6": 9230.86,
                "h1": 2534.66,
                "m5": 1244.42
            },
            "priceChange": {
                "m5": 0.33,
                "h1": 1.1,
                "h6": 1.74,
                "h24": 5.12
            },
            "liquidity": {
                "usd": 411411.98,
                "base": 59177860,
                "quote": 280.4711
            },
            "fdv": 161890819,
            "marketCap": 161890819,
            "pairCreatedAt": 1683945135000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x69d91b94f0aaf8e8a2586909fa77a5c2c89818d5",
            "pairAddress": "0x69D91B94f0AaF8e8A2586909fA77A5c2c89818d5",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39",
                "name": "HEX",
                "symbol": "HEX"
            },
            "priceNative": "0.2700",
            "priceUsd": "0.005715",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 1
                },
                "h6": {
                    "buys": 6,
                    "sells": 18
                },
                "h24": {
                    "buys": 47,
                    "sells": 46
                }
            },
            "volume": {
                "h24": 1436.15,
                "h6": 342.08,
                "h1": 67.58,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.84,
                "h6": 3.94,
                "h24": 4.88
            },
            "liquidity": {
                "usd": 432149.3,
                "base": 9583883,
                "quote": 17830685
            },
            "fdv": 162754997,
            "marketCap": 162754997,
            "pairCreatedAt": 1683944295000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x3041cbd36888becc7bbcbc0045e3b1f144466f5f",
            "pairAddress": "0x3041CbD36888bECc7bbCBc0045E3B1f144466f5f",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "priceNative": "1.9147",
            "priceUsd": "0.005654",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 2,
                    "sells": 0
                },
                "h6": {
                    "buys": 8,
                    "sells": 7
                },
                "h24": {
                    "buys": 27,
                    "sells": 80
                }
            },
            "volume": {
                "h24": 6189.44,
                "h6": 642.38,
                "h1": 108.31,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 3.07,
                "h6": 3.37,
                "h24": 5.41
            },
            "liquidity": {
                "usd": 114182.22,
                "base": 10096477,
                "quote": 19331932
            },
            "fdv": 161028290,
            "marketCap": 161028290,
            "pairCreatedAt": 1683955985000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x6c6bc977e13df9b0de53b251522280bb72383700",
            "pairAddress": "0x6c6Bc977E13Df9b0de53b251522280BB72383700",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.9998",
            "priceUsd": "0.005707",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 11
                },
                "h6": {
                    "buys": 14,
                    "sells": 12
                },
                "h24": {
                    "buys": 31,
                    "sells": 75
                }
            },
            "volume": {
                "h24": 6168.48,
                "h6": 1156.37,
                "h1": 219.04,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 3.24,
                "h6": 3.53,
                "h24": 4.92
            },
            "liquidity": {
                "usd": 383455.38,
                "base": 25576664,
                "quote": 41596637
            },
            "fdv": 162548406,
            "marketCap": 162548406,
            "pairCreatedAt": 1683944435000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x3416cf6c708da44db2624d63ea0aaef7113527c6",
            "pairAddress": "0x3416cF6C708Da44DB2624D63ea0AAef7113527C6",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "priceNative": "1.8957",
            "priceUsd": "0.004962",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 0
                }
            },
            "volume": {
                "h24": 1.06,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": 0
            },
            "liquidity": {
                "usd": 168912.44,
                "base": 159918,
                "quote": 64226131
            },
            "fdv": 141318715,
            "marketCap": 141318715,
            "pairCreatedAt": 1683944585000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5",
            "pairAddress": "0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.9990",
            "priceUsd": "0.005607",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 1
                },
                "h6": {
                    "buys": 3,
                    "sells": 1
                },
                "h24": {
                    "buys": 5,
                    "sells": 8
                }
            },
            "volume": {
                "h24": 106.43,
                "h6": 32.54,
                "h1": 3.26,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 3.65,
                "h6": 6.2,
                "h24": 10.25
            },
            "liquidity": {
                "usd": 154800.48,
                "base": 13802788,
                "quote": 13789932
            },
            "fdv": 159690486,
            "marketCap": 159690486,
            "pairCreatedAt": 1683960815000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0xa9607906406cf2cf98bc1a1f0224c85d779042bd",
            "pairAddress": "0xa9607906406Cf2Cf98bC1A1F0224C85d779042BD",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.5206",
            "priceUsd": "0.002963",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 4
                },
                "h6": {
                    "buys": 9,
                    "sells": 30
                },
                "h24": {
                    "buys": 76,
                    "sells": 88
                }
            },
            "volume": {
                "h24": 6891.62,
                "h6": 1304.9,
                "h1": 176.71,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1,
                "h6": 2.92,
                "h24": 2.04
            },
            "liquidity": {
                "usd": 74666.35,
                "base": 12597338,
                "quote": 6558499
            },
            "fdv": 107527965,
            "marketCap": 107527965,
            "pairCreatedAt": 1685661415000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x9db9e0e53058c89e5b94e29621a205198648425b",
            "pairAddress": "0x9Db9e0e53058C89e5B94e29621a205198648425B",
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
                "name": "Wrapped BTC",
                "symbol": "WBTC"
            },
            "priceNative": "0.00001106",
            "priceUsd": "0.002958",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 0
                },
                "h6": {
                    "buys": 1,
                    "sells": 0
                },
                "h24": {
                    "buys": 16,
                    "sells": 16
                }
            },
            "volume": {
                "h24": 1275.52,
                "h6": 35.34,
                "h1": 35.34,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 2.07,
                "h6": 2.07,
                "h24": 3.32
            },
            "liquidity": {
                "usd": 59417.9,
                "base": 18202172,
                "quote": 20.8265
            },
            "fdv": 107339748,
            "marketCap": 107339748,
            "pairCreatedAt": 1683946485000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x1ba4436daa19d77268ae1f13ca8294ec6a4c2486",
            "pairAddress": "0x1BA4436dAA19d77268Ae1f13cA8294eC6a4c2486",
            "labels": [
                "v1"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xCc78A0acDF847A2C1714D2A925bB4477df5d48a6",
                "name": "Atropa",
                "symbol": "ATROPA"
            },
            "priceNative": "0.07213",
            "priceUsd": "0.002987",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 5
                },
                "h6": {
                    "buys": 1,
                    "sells": 6
                },
                "h24": {
                    "buys": 34,
                    "sells": 35
                }
            },
            "volume": {
                "h24": 1539.79,
                "h6": 114.44,
                "h1": 38.43,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.48,
                "h6": 3.76,
                "h24": 3.25
            },
            "liquidity": {
                "usd": 56225.08,
                "base": 9409475,
                "quote": 678715
            },
            "fdv": 108402679,
            "marketCap": 108402679,
            "pairCreatedAt": 1685330735000
        },
        {
            "chainId": "pulsechain",
            "dexId": "0x06393b72873077c523ed10B3Db69087875AF61ef",
            "url": "https://dexscreener.com/pulsechain/0x506a1088f304ea844d3b81a63453dca140e4e7b6",
            "pairAddress": "0x506A1088f304EA844d3B81A63453DCa140E4e7B6",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.9994",
            "priceUsd": "0.005285",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 1
                }
            },
            "volume": {
                "h24": 29.02,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": 4.57
            },
            "liquidity": {
                "usd": 53306.89,
                "base": 5036134,
                "quote": 5046585
            },
            "fdv": 150518329,
            "marketCap": 150518329,
            "pairCreatedAt": 1685163115000
        },
        {
            "chainId": "pulsechain",
            "dexId": "0x06393b72873077c523ed10B3Db69087875AF61ef",
            "url": "https://dexscreener.com/pulsechain/0x52337dcd6aef2d2437005d7aa5de3b802cc63442",
            "pairAddress": "0x52337DcD6aeF2d2437005d7AA5De3B802cC63442",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "priceNative": "1.9133",
            "priceUsd": "0.005631",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 0
                },
                "h6": {
                    "buys": 3,
                    "sells": 3
                },
                "h24": {
                    "buys": 20,
                    "sells": 25
                }
            },
            "volume": {
                "h24": 1648.85,
                "h6": 265.53,
                "h1": 47.93,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 3.79,
                "h6": 3.02,
                "h24": 5.2
            },
            "liquidity": {
                "usd": 38775.4,
                "base": 3229198,
                "quote": 6995312
            },
            "fdv": 160374503,
            "marketCap": 160374503,
            "pairCreatedAt": 1685163355000
        },
        {
            "chainId": "pulsechain",
            "dexId": "9mm",
            "url": "https://dexscreener.com/pulsechain/0xd901f1d4eb400fdd897ecf26be3cca7659041698",
            "pairAddress": "0xD901f1D4eB400fdd897ecf26BE3cca7659041698",
            "labels": [
                "V3"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
                "name": "Dai Stablecoin",
                "symbol": "DAI"
            },
            "priceNative": "0.9984",
            "priceUsd": "0.005326",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 1,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 2
                }
            },
            "volume": {
                "h24": 30.29,
                "h6": 9.12,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 6.19,
                "h24": 6.19
            },
            "liquidity": {
                "usd": 30169.84,
                "base": 1924390,
                "quote": 3733568
            },
            "fdv": 151696007,
            "marketCap": 151696007,
            "pairCreatedAt": 1713616885000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x7858e59e0c01ea06df3af3d20ac7b0003275d4bf",
            "pairAddress": "0x7858E59e0C01EA06Df3aF3D20aC7B0003275D4Bf",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "priceNative": "1.9531",
            "priceUsd": "0.005425",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 0
                }
            },
            "volume": {
                "h24": 0.2,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": 0
            },
            "liquidity": {
                "usd": 29443.68,
                "base": 325254,
                "quote": 9964130
            },
            "fdv": 154507828,
            "marketCap": 154507828,
            "pairCreatedAt": 1683955575000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x0c7d9e061bffb8f6cac598b6322a9968dc463986",
            "pairAddress": "0x0C7D9E061bFfB8f6Cac598B6322a9968dc463986",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xeB2CEed77147893Ba8B250c796c2d4EF02a72B68",
                "name": "Pulse Drip",
                "symbol": "PDRIP"
            },
            "priceNative": "0.006216",
            "priceUsd": "0.005629",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 1,
                    "sells": 1
                },
                "h6": {
                    "buys": 5,
                    "sells": 4
                },
                "h24": {
                    "buys": 40,
                    "sells": 21
                }
            },
            "volume": {
                "h24": 1214.81,
                "h6": 228.05,
                "h1": 64.53,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.3,
                "h6": 2.15,
                "h24": 3.68
            },
            "liquidity": {
                "usd": 23583.31,
                "base": 2094578,
                "quote": 13021
            },
            "fdv": 160317807,
            "marketCap": 160317807,
            "pairCreatedAt": 1713191475000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x3861254332677c8c00cf6c6aad71ce880b283ceb",
            "pairAddress": "0x3861254332677C8c00CF6C6aaD71cE880B283cEb",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "priceNative": "1.9114",
            "priceUsd": "0.005644",
            "txns": {
                "m5": {
                    "buys": 1,
                    "sells": 0
                },
                "h1": {
                    "buys": 2,
                    "sells": 0
                },
                "h6": {
                    "buys": 4,
                    "sells": 7
                },
                "h24": {
                    "buys": 41,
                    "sells": 27
                }
            },
            "volume": {
                "h24": 1589.95,
                "h6": 64.69,
                "h1": 17.95,
                "m5": 3.79
            },
            "priceChange": {
                "m5": 0.08,
                "h1": 3.08,
                "h6": 3.56,
                "h24": 5.4
            },
            "liquidity": {
                "usd": 18414.84,
                "base": 1631122,
                "quote": 3117774
            },
            "fdv": 160751476,
            "marketCap": 160751476,
            "pairCreatedAt": 1685783685000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0xf6dcdce0ac3001b2f67f750bc64ea5beb37b5824",
            "pairAddress": "0xF6DCdce0ac3001B2f67F750bc64ea5beB37B5824",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39",
                "name": "HEX",
                "symbol": "HEX"
            },
            "priceNative": "0.2692",
            "priceUsd": "0.005690",
            "txns": {
                "m5": {
                    "buys": 1,
                    "sells": 0
                },
                "h1": {
                    "buys": 3,
                    "sells": 3
                },
                "h6": {
                    "buys": 5,
                    "sells": 29
                },
                "h24": {
                    "buys": 55,
                    "sells": 63
                }
            },
            "volume": {
                "h24": 1752.48,
                "h6": 370.39,
                "h1": 64.49,
                "m5": 0.08
            },
            "priceChange": {
                "m5": -0.25,
                "h1": 0.77,
                "h6": 2.68,
                "h24": 4.3
            },
            "liquidity": {
                "usd": 18633.62,
                "base": 1637125,
                "quote": 440828
            },
            "fdv": 162064897,
            "marketCap": 162064897,
            "pairCreatedAt": 1683946175000
        },
        {
            "chainId": "ethereumpow",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/ethereumpow/0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
            "pairAddress": "0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
                "name": "Wrapped Ether",
                "symbol": "WETH"
            },
            "priceNative": "0.000002611",
            "priceUsd": "0.00002891",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 4,
                    "sells": 1
                }
            },
            "volume": {
                "h24": 1.96,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": -0.01
            },
            "liquidity": {
                "usd": 14917.57,
                "base": 257938225,
                "quote": 673.5557
            },
            "fdv": 934356,
            "marketCap": 934356
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0xcc229dea3fa09cfddc686ed89e0867f600dfe6f9",
            "pairAddress": "0xcc229DEA3fa09CfdDC686eD89e0867F600DfE6F9",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xeB2CEed77147893Ba8B250c796c2d4EF02a72B68",
                "name": "Pulse Drip",
                "symbol": "PDRIP"
            },
            "priceNative": "0.003234",
            "priceUsd": "0.002928",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 2,
                    "sells": 0
                },
                "h6": {
                    "buys": 4,
                    "sells": 2
                },
                "h24": {
                    "buys": 25,
                    "sells": 24
                }
            },
            "volume": {
                "h24": 647.97,
                "h6": 78.19,
                "h1": 34.79,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.21,
                "h6": 2.27,
                "h24": 0.95
            },
            "liquidity": {
                "usd": 19042.21,
                "base": 3250755,
                "quote": 10513
            },
            "fdv": 106269420,
            "marketCap": 106269420,
            "pairCreatedAt": 1713741355000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0x3150e4162dfdd5c89a8653bdf1cbb9e09c11f42a",
            "pairAddress": "0x3150e4162DfDD5c89A8653bDf1CbB9E09C11F42a",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x5a24D7129B6f3FcAd2220296df28911880AD22B0",
                "name": "PLSZEN",
                "symbol": "PZEN"
            },
            "priceNative": "6.0004391",
            "priceUsd": "0.005367",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 5,
                    "sells": 0
                },
                "h6": {
                    "buys": 6,
                    "sells": 0
                },
                "h24": {
                    "buys": 19,
                    "sells": 7
                }
            },
            "volume": {
                "h24": 506.13,
                "h6": 76.66,
                "h1": 62.35,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 1.43,
                "h6": 1.76,
                "h24": 1.14
            },
            "liquidity": {
                "usd": 17571.69,
                "base": 1636713,
                "quote": 9821001
            },
            "fdv": 152867309,
            "marketCap": 152867309,
            "pairCreatedAt": 1683953815000
        },
        {
            "chainId": "ethereumpow",
            "dexId": "sushiswap",
            "url": "https://dexscreener.com/ethereumpow/0x06da0fd433c1a5d7a4faa01111c044910a184553",
            "pairAddress": "0x06da0fd433C1A5d7a4faa01111c044910A184553",
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
                "name": "Wrapped Ether",
                "symbol": "WETH"
            },
            "priceNative": "0.000002613",
            "priceUsd": "0.00002893",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 0
                }
            },
            "volume": {
                "h24": 2.15,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": 0
            },
            "liquidity": {
                "usd": 10149.55,
                "base": 175363054,
                "quote": 458.2708
            },
            "fdv": 935058,
            "marketCap": 935058
        },
        {
            "chainId": "ethereumpow",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/ethereumpow/0x4e68ccd3e89f51c3074ca5072bbac773960dfa36",
            "pairAddress": "0x4e68Ccd3E89f51C3074ca5072bbAC773960dFa36",
            "labels": [
                "v3"
            ],
            "baseToken": {
                "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "name": "Tether USD",
                "symbol": "USDT"
            },
            "quoteToken": {
                "address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
                "name": "Wrapped Ether",
                "symbol": "WETH"
            },
            "priceNative": "0.000002805",
            "priceUsd": "0.00003107",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 1,
                    "sells": 4
                }
            },
            "volume": {
                "h24": 2.01,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": -4.43
            },
            "liquidity": {
                "usd": 9522.21,
                "base": 71051708,
                "quote": 660.5368
            },
            "fdv": 1003932,
            "marketCap": 1003932
        },
        {
            "chainId": "pulsechain",
            "dexId": "9inch",
            "url": "https://dexscreener.com/pulsechain/0x6518ba656aa9c0d235202772214ff060359d5a8f",
            "pairAddress": "0x6518BA656Aa9C0d235202772214Ff060359D5a8f",
            "labels": [
                "v3"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xA6C4790cc7Aa22CA27327Cb83276F2aBD687B55b",
                "name": "X",
                "symbol": "X"
            },
            "priceNative": "0.03308",
            "priceUsd": "0.005316",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 0
                },
                "h1": {
                    "buys": 0,
                    "sells": 0
                },
                "h6": {
                    "buys": 0,
                    "sells": 0
                },
                "h24": {
                    "buys": 11,
                    "sells": 13
                }
            },
            "volume": {
                "h24": 1302.87,
                "h6": 0,
                "h1": 0,
                "m5": 0
            },
            "priceChange": {
                "m5": 0,
                "h1": 0,
                "h6": 0,
                "h24": -0.87
            },
            "liquidity": {
                "usd": 13804.21,
                "base": 656441,
                "quote": 64188
            },
            "fdv": 151403992,
            "marketCap": 151403992,
            "pairCreatedAt": 1723321155000
        },
        {
            "chainId": "pulsechain",
            "dexId": "pulsex",
            "url": "https://dexscreener.com/pulsechain/0x8f6dfb2fa2f7ccf9d7106e96207d8b947a89998a",
            "pairAddress": "0x8F6DfB2Fa2f7Ccf9d7106E96207d8B947a89998a",
            "labels": [
                "v2"
            ],
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0xA1077a294dDE1B09bB078844df40758a5D0f9a27",
                "name": "Wrapped Pulse",
                "symbol": "WPLS"
            },
            "priceNative": "103.05507",
            "priceUsd": "0.005698",
            "txns": {
                "m5": {
                    "buys": 0,
                    "sells": 2
                },
                "h1": {
                    "buys": 7,
                    "sells": 18
                },
                "h6": {
                    "buys": 28,
                    "sells": 39
                },
                "h24": {
                    "buys": 164,
                    "sells": 201
                }
            },
            "volume": {
                "h24": 1545.09,
                "h6": 251.1,
                "h1": 111.96,
                "m5": 0.5
            },
            "priceChange": {
                "m5": -0.01,
                "h1": 1.24,
                "h6": 3,
                "h24": 4.83
            },
            "liquidity": {
                "usd": 6002.22,
                "base": 526676,
                "quote": 54276646
            },
            "fdv": 162271464,
            "marketCap": 162271464,
            "pairCreatedAt": 1685662125000
        },
        {
            "chainId": "pulsechain",
            "dexId": "uniswap",
            "url": "https://dexscreener.com/pulsechain/0xf60afed42d276507b6bcae93157c66bee2cf332b",
            "pairAddress": "0xf60AfeD42D276507B6bcAe93157C66bee2Cf332b",
            "baseToken": {
                "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
                "name": "USD Coin",
                "symbol": "USDC"
            },
            "quoteToken": {
                "address": "0x444444444444C1a66F394025Ac839A535246FCc8",
                "name": "Genius",
                "symbol": "GENI"
            },
            "priceNative": "2414.3418",
            "priceUsd": "0.005659",
            "txns": {
                "m5": {
                    "buys": 1,
                    "sells": 0
                },
                "h1": {
                    "buys": 5,
                    "sells": 0
                },
                "h6": {
                    "buys": 7,
                    "sells": 0
                },
                "h24": {
                    "buys": 21,
                    "sells": 4
                }
            },
            "volume": {
                "h24": 308.91,
                "h6": 89.41,
                "h1": 83.89,
                "m5": 11.84
            },
            "priceChange": {
                "m5": 0.25,
                "h1": 2.69,
                "h6": 3.7,
                "h24": 5.71
            },
            "liquidity": {
                "usd": 8584.77,
                "base": 1257166,
                "quote": 626995201
            },
            "fdv": 161170621,
            "marketCap": 161170621,
            "pairCreatedAt": 1683951225000
        }
    ]`
