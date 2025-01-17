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

type EntryShort struct {
	// The SHA256 Hash of the entry you just created. You can use this hash to referece this entry in the future.
	EntryHash string `json:"entry_hash,omitempty"`
	// The current immutability stage of the new entry.
	Stage string `json:"stage,omitempty"`
}
