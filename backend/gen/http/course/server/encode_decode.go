// Code generated by goa v3.21.0, DO NOT EDIT.
//
// course HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/enrollment/design/api

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	course "github.com/enrollment/gen/course"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadAllCoursesResponse returns an encoder for responses returned by
// the course upload_all_courses endpoint.
func EncodeUploadAllCoursesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeUploadAllCoursesRequest returns a decoder for requests sent to the
// course upload_all_courses endpoint.
func DecodeUploadAllCoursesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UploadAllCoursesRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUploadAllCoursesRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewUploadAllCoursesPayload(&body)

		return payload, nil
	}
}

// EncodeUploadAllCoursesError returns an encoder for errors returned by the
// upload_all_courses course endpoint.
func EncodeUploadAllCoursesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadAllCoursesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "un_authorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadAllCoursesUnAuthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetAllCoursesResponse returns an encoder for responses returned by the
// course get_all_courses endpoint.
func EncodeGetAllCoursesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*course.Course)
		enc := encoder(ctx, w)
		body := NewGetAllCoursesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAllCoursesRequest returns a decoder for requests sent to the course
// get_all_courses endpoint.
func DecodeGetAllCoursesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body GetAllCoursesRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateGetAllCoursesRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewGetAllCoursesPayload(&body)

		return payload, nil
	}
}

// EncodeGetAllCoursesError returns an encoder for errors returned by the
// get_all_courses course endpoint.
func EncodeGetAllCoursesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetAllCoursesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "un_authorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetAllCoursesUnAuthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetUserAvailableCoursesResponse returns an encoder for responses
// returned by the course get_user_available_courses endpoint.
func EncodeGetUserAvailableCoursesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*course.Course)
		enc := encoder(ctx, w)
		body := NewGetUserAvailableCoursesResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetUserAvailableCoursesError returns an encoder for errors returned by
// the get_user_available_courses course endpoint.
func EncodeGetUserAvailableCoursesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetUserAvailableCoursesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "un_authorized":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetUserAvailableCoursesUnAuthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// unmarshalCoursePayloadRequestBodyToCourseCoursePayload builds a value of
// type *course.CoursePayload from a value of type *CoursePayloadRequestBody.
func unmarshalCoursePayloadRequestBodyToCourseCoursePayload(v *CoursePayloadRequestBody) *course.CoursePayload {
	res := &course.CoursePayload{
		Name:        *v.Name,
		Credits:     *v.Credits,
		CicleNumber: *v.CicleNumber,
	}

	return res
}

// marshalCourseCourseToCourseResponse builds a value of type *CourseResponse
// from a value of type *course.Course.
func marshalCourseCourseToCourseResponse(v *course.Course) *CourseResponse {
	res := &CourseResponse{
		ID:          v.ID,
		Name:        v.Name,
		Credits:     v.Credits,
		CicleNumber: v.CicleNumber,
	}

	return res
}
