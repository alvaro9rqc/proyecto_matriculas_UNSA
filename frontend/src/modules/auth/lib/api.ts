import { BACKEND_URL } from 'astro:env/client';
import type { User } from '@/modules/auth/types/user';
import type { ApiResponse } from '@/modules/core/types/api';
import { INTERNAL_SERVER_ERROR } from '@/modules/core/lib/errors';

class AuthApi {
  private apiAuthUrl: string;

  constructor() {
    this.apiAuthUrl = `${BACKEND_URL}/auth`;
  }

  async getUser(sessionToken?: string): ApiResponse<User> {
    try {
      console.log('sessionToken on aoi request: ', sessionToken);
      const response = await fetch(`${this.apiAuthUrl}/me`, {
        method: 'GET',
        headers: {
          Cookie: `session_token=${sessionToken}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        console.log(errorData);
        return {
          data: undefined,
          error: {
            status: response.status,
            message: errorData.message || 'Failed to fetch user',
          },
        };
      }

      const user: User = await response.json();
      return { data: user, error: undefined };
    } catch (error) {
      console.error('Error fetching user:', error);
      return {
        data: undefined,
        error: INTERNAL_SERVER_ERROR,
      };
    }
  }
}

export const authApi = new AuthApi();
