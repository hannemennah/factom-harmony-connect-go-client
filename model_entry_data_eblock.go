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

// Represents the Entry Block that contains the entry. This will be null if the entry is not at least at the `factom` immutability stage.
type EntryDataEblock struct {
	// The Key Merkle Root for this entry block.
	Keymr string `json:"keymr,omitempty"`
	// An API link to retrieve all information about this entry block.
	Href string `json:"href,omitempty"`
}
