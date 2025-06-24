import type { ErrorResponse } from '@/modules/core/types/error';

export const INTERNAL_SERVER_ERROR: ErrorResponse = {
  message: 'An error occurred while processing your request.',
  status: 500,
};

export const NOT_FOUND_ERROR: ErrorResponse = {
  message: 'The requested resource was not found.',
  status: 404,
};

export const UNAUTHORIZED_ERROR: ErrorResponse = {
  message: 'You are not authorized to access this resource.',
  status: 401,
};
