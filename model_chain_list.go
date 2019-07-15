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

// Represents a list of Factom entries that have been or will be saved to the blockchain.
type ChainList struct {
	// An array that contains the chains on this page.
	Data []ChainListData `json:"data"`
	// The index of the first chain returned from the total set (Starting from 0).
	Offset int32 `json:"offset"`
	// The number of chains returned.
	Limit int32 `json:"limit"`
	// The total number of chains seen.
	Count int32 `json:"count"`
}
