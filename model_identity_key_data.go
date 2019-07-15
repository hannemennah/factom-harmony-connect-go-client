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

type IdentityKeyData struct {
	// The public key string in Base58 idpub format.
	Key string `json:"key"`
	// The height at which this key became valid.
	ActivatedHeight int32 `json:"activated_height,omitempty"`
	// The expiration height of this key. `null` if this key is currently active.
	RetiredHeight int32 `json:"retired_height,omitempty"`
	// The level of this key within the hierarchy. A lower number indicates a key that allows a holder to replace higher numbered keys. The master key is priority 0.
	Priority int32 `json:"priority,omitempty"`
	// The unique identifier of the entry on the Identity chain where this key was activated.
	EntryHash string `json:"entry_hash,omitempty"`
}