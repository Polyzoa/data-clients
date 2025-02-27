package rugcheck

const tokenResponse = `{
  "tokenProgram": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
  "tokenType": "",
  "risks": [
    {
      "name": "Top 10 holders high ownership",
      "value": "",
      "description": "The top 10 users hold more than 70% token supply",
      "score": 9215,
      "level": "danger"
    },
    {
      "name": "Single holder ownership",
      "value": "80.00%",
      "description": "One user holds a large amount of the token supply",
      "score": 8000,
      "level": "danger"
    },
    {
      "name": "High ownership",
      "value": "",
      "description": "The top users hold more than 80% token supply",
      "score": 1411,
      "level": "danger"
    }
  ],
  "score": 18627
}`

const newTokens = `[
  {
    "mint": "4N5ij2HCyZmfTbWBt4sKPgBxRW7hCYka2sNEt7Nppump",
    "decimals": 6,
    "symbol": "Scientist",
    "creator": "2ginSopnDhN1H6AotvNsRuwhxoLgMgjR8tuGibmQx78q",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:08:15.814896592Z",
    "updatedAt": "2025-02-27T06:08:15.814997153Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "8neoQWrJYw6ecQRhyF4HGhVsuL13Vc456ECiQzXwpump",
    "decimals": 6,
    "symbol": "DOE",
    "creator": "FJoqvcCn3Z2dmLhxBTsCozqkWUwtwArUrPgAravgrEEj",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:08:15.164431824Z",
    "updatedAt": "2025-02-27T06:08:15.164503761Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "3e3M9L8fNR7uysjCt6c4orRDjW5sQuHuwKVESkV9pump",
    "decimals": 6,
    "symbol": "RAMCAT",
    "creator": "47inwdgJWF8P7jmc2aHSxXjsAuQRmxqqwTjMP5ZGUpas",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:08:14.000364512Z",
    "updatedAt": "2025-02-27T06:08:14.000460064Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "DtdfRHdTVR27thCFwQsXh2pcKCSYPMKtftTDZoFJyFkj",
    "decimals": 9,
    "symbol": "MINU",
    "creator": "",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb",
    "createAt": "2025-02-27T06:07:55.689717223Z",
    "updatedAt": "2025-02-27T06:07:55.689847201Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "7T6HttDvtHNCax9WRiWJUtYfzne4Hap5fq4d2MSQpump",
    "decimals": 6,
    "symbol": "Freedom ",
    "creator": "5mjvgKwA9NUk4fzMMWzxcf17zgnp3YRpYyAnWVoWnc6p",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:07:47.966422771Z",
    "updatedAt": "2025-02-27T06:07:47.966513965Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "HjWhWwDsTSg4BBR7WyRZuoP8EUTQX25CmtX8ifJ8pump",
    "decimals": 6,
    "symbol": "USDT",
    "creator": "TSLvdd1pWpHVjahSpsvCXUbgwsL3JAcvokwaKt1eokM",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:07:44.9827965Z",
    "updatedAt": "2025-02-27T06:07:44.982871082Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "9GV2UANcuQ2VTdG15Dj2hGuuCJfGFDNcLvykhKZzpump",
    "decimals": 6,
    "symbol": "GOLD go",
    "creator": "5jjmJg1WVoZULWg8SmAnHn261n7gT7jB2hkyXV4BZe4Z",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:07:42.744876012Z",
    "updatedAt": "2025-02-27T06:07:42.74500588Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "5Qm57LtF5du4XS5JR4F3ZKLT43MvWeo7CFNHdXfBpump",
    "decimals": 6,
    "symbol": "red",
    "creator": "GbPSQdZprFJLoqUbTafBWRbaWSMZhmWToTbvaJLwUMHG",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:07:40.31046596Z",
    "updatedAt": "2025-02-27T06:07:40.310578985Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "4aJgoLrKuevurFYHoELe3h2ZbFtqwinNaNB4wZgAE27W",
    "decimals": 9,
    "symbol": "oGPU",
    "creator": "",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb",
    "createAt": "2025-02-27T06:07:33.767235398Z",
    "updatedAt": "2025-02-27T06:07:33.767302476Z",
    "supply": 0,
    "events": null
  },
  {
    "mint": "8DHPkSPzSQLYTpjDGF9sjjsb3fnLe2DFPGUTqBa5UKqG",
    "decimals": 9,
    "symbol": "NOUSPOLIS!",
    "creator": "Db7W7o2F2QxMRZAVAPtT97YW8bewWbpfw26souoSE2ig",
    "mintAuthority": "",
    "freezeAuthority": "",
    "program": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
    "createAt": "2025-02-27T06:07:22.896841731Z",
    "updatedAt": "2025-02-27T06:07:22.896918868Z",
    "supply": 0,
    "events": null
  }
]`
