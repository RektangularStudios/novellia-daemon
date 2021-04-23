/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

// ProductArtist - Artist attribution
type ProductArtist struct {

	// Name of artist for token asset
	Name string `json:"name,omitempty"`

	// URLs to learn more about artist (e.g. portfolio). These can point to different places.
	Urls []string `json:"urls,omitempty"`
}
