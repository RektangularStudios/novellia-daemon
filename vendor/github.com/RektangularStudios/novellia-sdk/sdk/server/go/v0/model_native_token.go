/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.4.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

// NativeToken - Definitions required to reference a native token
type NativeToken struct {

	// On-chain policy-id
	PolicyId string `json:"policy_id"`

	// On-chain asset-id
	AssetId string `json:"asset_id"`
}