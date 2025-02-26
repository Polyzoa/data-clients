package solsniffer

const tokenResponse = `{
    "tokenData": {
        "indicatorData": {
            "high": {
                "count": 4,
                "details": "{\"Mintable risks found\":true,\"Freeze risks found\":true,\"A private wallet owns a significant share of the supply\":false,\"Tokens auto-freeze risks found\":true,\"Significant ownership by top 10 wallets\":false,\"Significant ownership by top 20 wallets\":false,\"Permanent control risks found\":true,\"Presence of token metadata\":true,\"High locked supply risks found\":false,\"Sufficient liquidity detected\":true,\"Very low liquidity\":true}"
            },
            "moderate": {
                "count": 0,
                "details": "{\"Token metadata are immutable\":true,\"Token operates without custom fees\":true,\"Token has recent user activity\":true,\"Unknown liquidity pools\":true,\"Low count of LP providers\":true}"
            },
            "low": {
                "count": 0,
                "details": "{\"Contract was not recently deployed\":true}"
            },
            "specific": {
                "count": 0,
                "details": "{\"Recent interaction within the last 30 days\":true}"
            }
        },
        "tokenOverview": {
            "deployer": "9k8jWWqfmTTXrc8gmZqYVxSFJhygURvyVRztivJXFTYP",
            "mint": "5e2qRc1DNEXmyxP8qwPwJhRWjef7usLyi7v5xjqLr5G7",
            "address": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
            "type": "spl-token"
        },
        "address": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
        "tokenName": "OFFICIAL TRUMP",
        "deployTime": "2025-01-17T14:01:48.000Z",
        "externals": "{\"website\":\"https://gettrumpmemes.com\",\"coingecko_coin_id\":\"official-trump\",\"twitter_handle\":\"realDonaldTrump\"}",
        "liquidityList": [
            {
                "fluxbeam": {
                    "address": "7jPqJRFiVJ9Gbru3DU4TMow4r14nD5zox9VEUFXJ5w8n",
                    "amount": 0.03,
                    "lpPair": "SOL"
                }
            },
            {
                "raydium": {
                    "address": "C6aQ3EM15Z1UnFU7RPx97U8FBP18gddTtwbgofKxkWkd",
                    "amount": 1.74,
                    "lpPair": "SOL"
                }
            },
            {
                "raydium": {
                    "address": "uCk125EJf7iCjz43aSdzuyMAT5eii6c5zKVxEnGnosa",
                    "amount": 7858.8,
                    "lpPair": "SOL"
                }
            },
            {
                "raydium": {
                    "address": "HPUFRpydYrA5FCBm7W1oL6JdbRg26qt6H5wo4yAud4vq",
                    "amount": 1778.56,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "6nD6d8gG17wakW6Wu5URktBZQp3uxp5orgPa576QXigJ",
                    "amount": 931844.38,
                    "lpPair": "USDC"
                }
            },
            {
                "orca": {
                    "address": "6KX9iiLFBcwfjq3uMqeeMukaMZt5rQYTsbZZTnxbzsz6",
                    "amount": 229684.11,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "EdRQgfs2oRyyqsGDNH5XPJKMUDUBdpPcy597ZHdYY7uk",
                    "amount": 80368.93,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "BYUb9Z3CJY4T8QjUQKS1npAsqX8qyoSR22B1GUtFQ45t",
                    "amount": 3424.95,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "AbTXfZfd2YnR8w1uv81NzHT4ggmq6jombkAFPcpprnfr",
                    "amount": 41018.03,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "6PeMbWGxqf4EBHBjCND8KC3Mh9t2KXwNRVkQrxFis5j8",
                    "amount": 1067.78,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "CSXm32rwS78fy5fvHqCADduYkdpinrPZ42b8LH6bQQjo",
                    "amount": 30.4,
                    "lpPair": "USDC"
                }
            },
            {
                "orca": {
                    "address": "Ecjn9eVBDC5h1eKDyLABGrauuGQTvGmJYUWjKMCmQmBG",
                    "amount": 16.98,
                    "lpPair": "SOL"
                }
            },
            {
                "orca": {
                    "address": "23B8CFoTCcHFvtujatuJWgzgRW9MNjHiVJN2KRLziHzt",
                    "amount": 57.48,
                    "lpPair": "USDC"
                }
            },
            {
                "orca": {
                    "address": "3F5onEzLdRdbT3F7gdSh8AnTBmtyxiRHvjxYcbtBRDTe",
                    "amount": 5.48,
                    "lpPair": "USDC"
                }
            },
            {
                "raydiumCpmm": {
                    "address": "HKuJrP5tYQLbEUdjKwjgnHs2957QKjR2iWhJKTtMa1xs",
                    "amount": 5460205.22,
                    "lpPair": "SOL"
                }
            },
            {
                "raydiumCpmm": {
                    "address": "Grub1v4mcDtkBmTmsUPJ1dNbELt3ayB4PkFwFBxmqu74",
                    "amount": 22655.53,
                    "lpPair": "USDC"
                }
            },
            {
                "raydiumCpmm": {
                    "address": "622n7BiBqYUtn1gfid3ucGXwRCqLvyL5B4rWwBNHTTKZ",
                    "amount": 0.02,
                    "lpPair": "USDC"
                }
            },
            {
                "raydiumCpmm": {
                    "address": "ATuMwubNp1bTFSxpC8fuF6tHUksufeSQUJDVX3MiByfj",
                    "amount": 1.07,
                    "lpPair": "USDC"
                }
            },
            {
                "meteora": {
                    "address": "4ZGaDNvLF31pfmyCLbMD4MqiR24VRda3UGmov41y4AvV",
                    "amount": 32780.4,
                    "lpPair": "SOL"
                }
            },
            {
                "meteora": {
                    "address": "4eNPDEYyckG5jgz5EJVuMzk4VGrcJEwFoHNFjcv6fbC5",
                    "amount": 323.42,
                    "lpPair": "USDC"
                }
            },
            {
                "meteora": {
                    "address": "88PJWPNq776eV9VoJThMQdDG6htkRKF5kLkDahAPDHaC",
                    "amount": 48.17,
                    "lpPair": "SOL"
                }
            },
            {
                "meteora": {
                    "address": "F7qRGffz4oRBJgPs9Bj1PRjUsasqSkw5KgeSyGYjsMRL",
                    "amount": 26.2,
                    "lpPair": "SOL"
                }
            },
            {
                "meteora": {
                    "address": "8hRU9Hh5aCoYzTJZwDAu4PtBCgaexm6E9ezHELvb2PTC",
                    "amount": 586.01,
                    "lpPair": "SOL"
                }
            }
        ],
        "marketCap": 13048054732.76,
        "ownersList": [
            {
                "address": "2RH6rUTPBJ9rUDPpuV9b8z1YL56k1tYU6Uk5ZoaEFFSK",
                "amount": "800000021.56",
                "percentage": "80.00"
            },
            {
                "address": "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM",
                "amount": "30151426.03",
                "percentage": "3.02"
            },
            {
                "address": "3gd3dqgtJ4jWfBfLYTX67DALFetjc5iS72sCgRhCkW2u",
                "amount": "20000000.00",
                "percentage": "2.00"
            },
            {
                "address": "5e2qRc1DNEXmyxP8qwPwJhRWjef7usLyi7v5xjqLr5G7",
                "amount": "15019349.62",
                "percentage": "1.50"
            },
            {
                "address": "9un5wqE3q4oCjyrDkwsdD48KteCJitQX5978Vh7KKxHo",
                "amount": "14991849.90",
                "percentage": "1.50"
            },
            {
                "address": "GRvmQxtFRqtQZYaeF8bUqH8FrDHoK5Eabs5aYyU34bPg",
                "amount": "12540010.00",
                "percentage": "1.25"
            },
            {
                "address": "9d9mb8kooFfaD3SctgZtkxQypkshx6ezhbKio89ixyy2",
                "amount": "9724757.28",
                "percentage": "0.97"
            },
            {
                "address": "8Tp9fFkZ2KcRBLYDTUNXo98Ez6ojGb6MZEPXfGDdeBzG",
                "amount": "7700000.00",
                "percentage": "0.77"
            },
            {
                "address": "8N2ssXZGJbvVLszanERuCJpLZd2nuADgpZwLkDDxwNnS",
                "amount": "5846368.20",
                "percentage": "0.58"
            },
            {
                "address": "42brAgAVNzMBP7aaktPvAmBSPEkehnFQejiZc53EpJFd",
                "amount": "5150000.00",
                "percentage": "0.52"
            },
            {
                "address": "6brjeZNfSpqjWoo16z1YbywKguAruXZhNz9bJMVZE8pD",
                "amount": "4851815.08",
                "percentage": "0.49"
            },
            {
                "address": "5tzFkiKscXHK5ZXCGbXZxdw7gTjjD1mBwuoFbhUvuAi9",
                "amount": "4417282.87",
                "percentage": "0.44"
            },
            {
                "address": "u6PJ8DtQuPFnfmwHbGFULQ4u4EgjDiyYKjVEsynXq2w",
                "amount": "2759322.78",
                "percentage": "0.28"
            },
            {
                "address": "8Mm46CsqxiyAputDUp2cXHg41HE3BfynTeMBDwzrMZQH",
                "amount": "2349649.61",
                "percentage": "0.23"
            },
            {
                "address": "22Wnk8PwyWZV7BfkZGJEKT9jGGdtvu7xY6EXeRh7zkBa",
                "amount": "2264839.73",
                "percentage": "0.23"
            },
            {
                "address": "A77HErqtfN1hLLpvZ9pCtu66FEtM8BveoaKbbMoZ4RiR",
                "amount": "1970534.88",
                "percentage": "0.20"
            },
            {
                "address": "AC5RDfQFmDS1deWZos921JfqscXdByf8BKHs5ACWjtW2",
                "amount": "1965229.25",
                "percentage": "0.20"
            },
            {
                "address": "AP84nNhMKzqTcj1dWs7deDk2BYawcosh6iuf5c6Hgj8x",
                "amount": "1926711.19",
                "percentage": "0.19"
            },
            {
                "address": "FWznbcNXWQuHTawe9RxvQ2LdCENssh12dsznf4RiouN5",
                "amount": "1736357.00",
                "percentage": "0.17"
            },
            {
                "address": "3B7XAQrLoEMDEGvX8569F9GRb9PcXXocJmj5wvhhei9z",
                "amount": "1681862.67",
                "percentage": "0.17"
            },
            {
                "address": "6FEVkH17P9y8Q9aCkDdPcMDjvj7SVxrTETaYEm8f51Jy",
                "amount": "1652213.59",
                "percentage": "0.17"
            },
            {
                "address": "7qtDv72fGzuuGmyg1FkTxBybYH2sXkUd6u8WXwnkRnpE",
                "amount": "1571346.14",
                "percentage": "0.16"
            },
            {
                "address": "9cNE6KBg2Xmf34FPMMvzDF8yUHMrgLRzBV3vD7b1JnUS",
                "amount": "1247277.21",
                "percentage": "0.12"
            },
            {
                "address": "8NBEbxLknGv5aRYefFrW2qFXoDZyi9fSHJNiJRvEcMBE",
                "amount": "1226002.40",
                "percentage": "0.12"
            },
            {
                "address": "EucF6LTAvLgksSwWNPjuyBG9S1iVkDABdmnGs8ficiTy",
                "amount": "1195106.34",
                "percentage": "0.12"
            },
            {
                "address": "DBmae92YTQKLsNzXcPscxiwPqMcz9stQr2prB5ZCAHPd",
                "amount": "1173504.02",
                "percentage": "0.12"
            },
            {
                "address": "F7RkX6Y1qTfBqoX5oHoZEgrG1Dpy55UZ3GfWwPbM58nQ",
                "amount": "1133004.80",
                "percentage": "0.11"
            },
            {
                "address": "ASTyfSima4LLAdDgoFGkgqoKowG1LZFDr9fAQrg7iaJZ",
                "amount": "938040.45",
                "percentage": "0.09"
            },
            {
                "address": "hTvwKr1RvQdPS5xiWfXM2UZYuF55Ei8zzsuB7e58feu",
                "amount": "875249.53",
                "percentage": "0.09"
            },
            {
                "address": "FixDET3jRBMjT9M69Wim8iRpMQiJHkSSfD7wivStuuph",
                "amount": "836590.46",
                "percentage": "0.08"
            },
            {
                "address": "2Ejnns2Fd5gsZdFJbnkZmQEwbrgoiMc2ikKcS2z2Ps3e",
                "amount": "811471.80",
                "percentage": "0.08"
            },
            {
                "address": "A8nPhpCJqtqHdqUk35Uj9Hy2YsGXFkCZGuNwvkD3k7VC",
                "amount": "628661.91",
                "percentage": "0.06"
            },
            {
                "address": "3AioVXJzq2P8FX3hDEa3k8nBE1WrfRDqbZTFQLwGGK4E",
                "amount": "558805.09",
                "percentage": "0.06"
            },
            {
                "address": "Cj22juEv1LjsHFmyu7ALaThwNeXMWPDJde8QayNNtBz3",
                "amount": "558805.09",
                "percentage": "0.06"
            },
            {
                "address": "4qchfp339j853WVeh1jbmBerEHKnzn7W6dNaH8BWN6rV",
                "amount": "555861.85",
                "percentage": "0.06"
            },
            {
                "address": "JqDc3BE4pVpUtcBfUMhy8RT97yuDRcjJ3Yt7q8Pwjsv",
                "amount": "555630.45",
                "percentage": "0.06"
            },
            {
                "address": "2yXxaSQzMFnqewM2jsjECvSWs6TXA9r5doJnPyTbs3xo",
                "amount": "553795.77",
                "percentage": "0.06"
            },
            {
                "address": "7KxBq2D862Ab7E1daJGQasNnqqgwXbqrnMNmwYmg5rWY",
                "amount": "539162.40",
                "percentage": "0.05"
            },
            {
                "address": "22KV28a9dz7U8SBRv8tbwW7burFAzHQFWTj9Qed3QcUP",
                "amount": "537801.54",
                "percentage": "0.05"
            },
            {
                "address": "6qgBGeZgPyxdobeHhcNtAqVe927zodpiuoufhwGN8BhP",
                "amount": "533649.79",
                "percentage": "0.05"
            },
            {
                "address": "6brjeZNfSpqjWoo16z1YbywKguAruXZhNz9bJMVZE8pD",
                "amount": "4851815.08",
                "percentage": "0.49"
            },
            {
                "address": "5tzFkiKscXHK5ZXCGbXZxdw7gTjjD1mBwuoFbhUvuAi9",
                "amount": "4417282.87",
                "percentage": "0.44"
            },
            {
                "address": "u6PJ8DtQuPFnfmwHbGFULQ4u4EgjDiyYKjVEsynXq2w",
                "amount": "2759322.78",
                "percentage": "0.28"
            },
            {
                "address": "8Mm46CsqxiyAputDUp2cXHg41HE3BfynTeMBDwzrMZQH",
                "amount": "2349649.61",
                "percentage": "0.23"
            },
            {
                "address": "22Wnk8PwyWZV7BfkZGJEKT9jGGdtvu7xY6EXeRh7zkBa",
                "amount": "2264839.73",
                "percentage": "0.23"
            },
            {
                "address": "A77HErqtfN1hLLpvZ9pCtu66FEtM8BveoaKbbMoZ4RiR",
                "amount": "1970534.88",
                "percentage": "0.20"
            },
            {
                "address": "AC5RDfQFmDS1deWZos921JfqscXdByf8BKHs5ACWjtW2",
                "amount": "1965229.25",
                "percentage": "0.20"
            },
            {
                "address": "AP84nNhMKzqTcj1dWs7deDk2BYawcosh6iuf5c6Hgj8x",
                "amount": "1926711.19",
                "percentage": "0.19"
            },
            {
                "address": "FWznbcNXWQuHTawe9RxvQ2LdCENssh12dsznf4RiouN5",
                "amount": "1736357.00",
                "percentage": "0.17"
            },
            {
                "address": "3B7XAQrLoEMDEGvX8569F9GRb9PcXXocJmj5wvhhei9z",
                "amount": "1681862.67",
                "percentage": "0.17"
            }
        ],
        "score": 55,
        "tokenImg": "https://arweave.net/VQrPjACwnQRmxdKBTqNwPiyo65x7LAT773t8Kd7YBzw",
        "tokenSymbol": "TRUMP",
        "auditRisk": {
            "mintDisabled": true,
            "freezeDisabled": true,
            "lpBurned": false,
            "top10Holders": false
        }
    },
    "tokenInfo": {
        "price": "12.906179",
        "supplyAmount": 999999596.266929,
        "mktCap": 12906173789.348719
    }
}`
