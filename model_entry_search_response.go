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

type EntrySearchResponse struct {
	// An array of entries that match your search criterion.
	Data []EntrySearchResponseData `json:"data"`
	// The index of the first item returned from the total set (Starting from 0).
	Offset int32 `json:"offset"`
	// The maximum number of entries per page.
	Limit int32 `json:"limit"`
	// The total number of entries seen.
	Count int32 `json:"count"`
}
