/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.6.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

// OccultaNovelliaCharacter - Model for an Occulta Novellia character
type OccultaNovelliaCharacter struct {

	// Specification version for parsing the resource
	OccultaNovelliaVersion int32 `json:"occulta_novellia_version"`

	// Character name
	Name string `json:"name"`

	Card OccultaNovelliaCharacterCard `json:"card"`

	Progression OccultaNovelliaCharacterProgression `json:"progression"`

	Stats OccultaNovelliaCharacterStats `json:"stats"`

	// List of attribute strings relating to character abilities, faction, types, etc.
	Attributes []string `json:"attributes"`

	// Description of character
	Description string `json:"description"`
}
