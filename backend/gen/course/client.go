// Code generated by goa v3.21.0, DO NOT EDIT.
//
// course client
//
// Command:
// $ goa gen github.com/enrollment/design/api

package course

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "course" service client.
type Client struct {
	UploadAllCoursesEndpoint        goa.Endpoint
	GetAllCoursesEndpoint           goa.Endpoint
	GetUserAvailableCoursesEndpoint goa.Endpoint
}

// NewClient initializes a "course" service client given the endpoints.
func NewClient(uploadAllCourses, getAllCourses, getUserAvailableCourses goa.Endpoint) *Client {
	return &Client{
		UploadAllCoursesEndpoint:        uploadAllCourses,
		GetAllCoursesEndpoint:           getAllCourses,
		GetUserAvailableCoursesEndpoint: getUserAvailableCourses,
	}
}

// UploadAllCourses calls the "upload_all_courses" endpoint of the "course"
// service.
// UploadAllCourses may return the following errors:
//   - "not_found" (type *goa.ServiceError): The resource was not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - "un_authorized" (type *goa.ServiceError): Unauthorized access
//   - error: internal error
func (c *Client) UploadAllCourses(ctx context.Context, p *UploadAllCoursesPayload) (err error) {
	_, err = c.UploadAllCoursesEndpoint(ctx, p)
	return
}

// GetAllCourses calls the "get_all_courses" endpoint of the "course" service.
// GetAllCourses may return the following errors:
//   - "not_found" (type *goa.ServiceError): The resource was not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - "un_authorized" (type *goa.ServiceError): Unauthorized access
//   - error: internal error
func (c *Client) GetAllCourses(ctx context.Context, p *GetAllCoursesPayload) (res []*Course, err error) {
	var ires any
	ires, err = c.GetAllCoursesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Course), nil
}

// GetUserAvailableCourses calls the "get_user_available_courses" endpoint of
// the "course" service.
// GetUserAvailableCourses may return the following errors:
//   - "not_found" (type *goa.ServiceError): The resource was not found
//   - "bad_request" (type *goa.ServiceError): Invalid request
//   - "un_authorized" (type *goa.ServiceError): Unauthorized access
//   - error: internal error
func (c *Client) GetUserAvailableCourses(ctx context.Context) (res []*Course, err error) {
	var ires any
	ires, err = c.GetUserAvailableCoursesEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Course), nil
}
