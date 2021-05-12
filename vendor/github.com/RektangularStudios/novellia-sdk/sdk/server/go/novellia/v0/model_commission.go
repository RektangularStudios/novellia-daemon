/*
 * novellia
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.8.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia

// Commission - A recommended commission for clients that will abide by it
type Commission struct {

	// Name of entity taking commissions
	Name string `json:"name"`

	// Address to deposit commissions
	Address string `json:"address"`

	// Recommended commission percentage
	Percent float32 `json:"percent,omitempty"`
}