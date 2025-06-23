import type { ErrorResponse } from '@/modules/core/types/error';

export const INTERNAL_SERVER_ERROR: ErrorResponse = {
  message: 'An error occurred while processing your request.',
  status: 500,
};
