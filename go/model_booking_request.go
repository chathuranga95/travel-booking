/*
 * Travel Management service
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BookingRequest struct {
	EmpId int32 `json:"empId"`

	From string `json:"from,omitempty"`

	Destination string `json:"destination,omitempty"`
}
