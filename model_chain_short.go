/*
 * Harmony Connect
 *
 * An easy to use API that helps you access the Factom blockchain.
 *
 * API version: 1.0.19
 * Contact: harmony-support@factom.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package connectclient

// Represents a list of Factom chains that have been or will be saved to the blockchain.
type ChainShort struct {
	// This is the unique identifier created for each chain.
	ChainId string `json:"chain_id"`
	// The unique identifier of the chain's first entry.
	EntryHash string `json:"entry_hash"`
	// The immutability stage that this chain has reached.
	Stage string `json:"stage"`
}
