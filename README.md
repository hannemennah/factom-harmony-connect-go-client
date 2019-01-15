# Factom Harmony Connect - Go Client Library

This is an automatically generated Go client library for [Factom Harmony Connect](https://www.factom.com/products/harmony-connect/).

Connect is the fastest way to add blockchain capabilities to your app without cryptocurrencies, wallets, or network nodes. [Create an account](https://account.factom.com/) to get your free API key for the sandbox environment.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.17
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.harmony.factom.com](https://docs.harmony.factom.com)

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./connectclient"
```

## Documentation for API Endpoints

All URIs are relative to *https://connect-shared-sandbox-2445582615332.production.gw.apicast.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChainsApi* | [**GetChainByID**](docs/ChainsApi.md#getchainbyid) | **Get** /chains/{chain_id} | Get Chain Info
*ChainsApi* | [**GetChains**](docs/ChainsApi.md#getchains) | **Get** /chains | Get All Chains
*ChainsApi* | [**PostChain**](docs/ChainsApi.md#postchain) | **Post** /chains | Create a Chain
*ChainsApi* | [**PostChainSearch**](docs/ChainsApi.md#postchainsearch) | **Post** /chains/search | Search Chains
*EntriesApi* | [**GetEntriesByChainID**](docs/EntriesApi.md#getentriesbychainid) | **Get** /chains/{chain_id}/entries | Get Chain&#39;s Entries
*EntriesApi* | [**GetEntryByHash**](docs/EntriesApi.md#getentrybyhash) | **Get** /chains/{chain_id}/entries/{entry_hash} | Get Entry Info
*EntriesApi* | [**GetFirstEntry**](docs/EntriesApi.md#getfirstentry) | **Get** /chains/{chain_id}/entries/first | Get Chain&#39;s First Entry
*EntriesApi* | [**GetLastEntry**](docs/EntriesApi.md#getlastentry) | **Get** /chains/{chain_id}/entries/last | Get Chain&#39;s Last Entry
*EntriesApi* | [**PostEntriesSearch**](docs/EntriesApi.md#postentriessearch) | **Post** /chains/{chain_id}/entries/search | Search Chain&#39;s Entries
*EntriesApi* | [**PostEntryToChainID**](docs/EntriesApi.md#postentrytochainid) | **Post** /chains/{chain_id}/entries | Create an Entry
*InfoApi* | [**GetAllInfo**](docs/InfoApi.md#getallinfo) | **Get** / | API Info


## Documentation For Models

 - [AllInfo](docs/AllInfo.md)
 - [AllInfoLinks](docs/AllInfoLinks.md)
 - [Chain](docs/Chain.md)
 - [ChainCreate](docs/ChainCreate.md)
 - [ChainData](docs/ChainData.md)
 - [ChainDataEntries](docs/ChainDataEntries.md)
 - [ChainLink](docs/ChainLink.md)
 - [ChainList](docs/ChainList.md)
 - [ChainListData](docs/ChainListData.md)
 - [ChainShort](docs/ChainShort.md)
 - [Entry](docs/Entry.md)
 - [EntryCreate](docs/EntryCreate.md)
 - [EntryData](docs/EntryData.md)
 - [EntryDataEblock](docs/EntryDataEblock.md)
 - [EntryLink](docs/EntryLink.md)
 - [EntryLinkChain](docs/EntryLinkChain.md)
 - [EntryList](docs/EntryList.md)
 - [EntryListData](docs/EntryListData.md)
 - [EntrySearchResponse](docs/EntrySearchResponse.md)
 - [EntrySearchResponseData](docs/EntrySearchResponseData.md)
 - [EntryShort](docs/EntryShort.md)
 - [SearchBody](docs/SearchBody.md)


## Documentation For Authorization

## AppId
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
})
r, err := client.Service.Operation(auth, args)
```
## AppKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
})
r, err := client.Service.Operation(auth, args)
```

## Support

For more information, you can view the Connect documentation at [https://docs.harmony.factom.com](https://docs.harmony.factom.com)


For additional support, contact us at harmony-support@factom.com

