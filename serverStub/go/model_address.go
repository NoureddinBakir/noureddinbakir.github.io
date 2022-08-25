/*
 * Eyada API
 *
 * CRUD operations for Doctors and Patients, and the Eyade-Scheduler
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package eyadaAPI

type Address struct {

	Street string `json:"street,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Zip string `json:"zip,omitempty"`
}
