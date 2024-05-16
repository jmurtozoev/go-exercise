# Golang Developer Assigment

Develop a service to provide an API for retrieval of Last Traded Price of Bitcoin for the following currency pairs:
Your task is to build an API application using go - that would retrieve the Last

1. BTC/USD
2. BTC/CHF
3. BTC/EUR


The request path is:
/api/v1/ltp

The response shall constitute JSON of the following structure:
```json
{
  "ltp": [
    {
      "pair": "BTC/CHF",
      "amount": "49000.12"
    },
    {
      "pair": "BTC/EUR",
      "amount": "50000.12"
    },
    {
      "pair": "BTC/USD",
      "amount": "52000.12"
    }
  ]
}

```

The public Kraken API might be used to retrieve the above LTP information
[API Documentation](https://docs.kraken.com/rest/#tag/Spot-Market-Data/operation/getTickerInformation) 
(The values of the last traded price is called “last trade closed”)

# Requirements:
1. Code shall be hosted in a remote public repository
2. readme.md includes clear steps to build and run the app
3. Integration tests
4. Dockerized application
