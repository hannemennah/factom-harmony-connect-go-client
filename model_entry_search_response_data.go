/*
 * Harmony Connect
 *
 * An easy to use API that helps you access the Factom blockchain.
 *
 * API version: 1.0.17
 * Contact: harmony-support@factom.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package connectclient

type EntrySearchResponseData struct {
	// The SHA256 Hash of this entry.
	EntryHash string `json:"entry_hash"`
	// Tags that can be used to identify this entry.
	ExternalIds []string `json:"external_ids"`
	// The level of immutability that this entry has reached.
	Stage string `json:"stage"`
	// An API link to retrieve all information about this entry.
	Href string `json:"href"`
}
