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

type EntryListData struct {
	// The SHA256 Hash of this entry.
	EntryHash string `json:"entry_hash"`
	Chain EntryListChain `json:"chain"`
	// The time at which this entry was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: `YYYY-MM-DDThh:mm:ss.ssssssZ`
	CreatedAt string `json:"created_at,omitempty"`
	// An API link to retrieve all information about this entry.
	Href string `json:"href"`
}
