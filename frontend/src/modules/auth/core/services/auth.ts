import { BACKEND_URL } from 'astro:env/client';
import type { ApiResponse } from '@/modules/core/types/api';
import {
  INTERNAL_SERVER_ERROR,
  UNAUTHORIZED_ERROR,
} from '@/modules/core/lib/errors';
import type { AstroCookies } from 'astro';
import type { User } from '@/modules/auth/core/types/user';

class AuthService {
  private apiAuthUrl: string;

  constructor() {
    this.apiAuthUrl = `${BACKEND_URL}/auth`;
  }

  async validateSessionToken(cookies: AstroCookies): ApiResponse<string> {
    const sessionToken = cookies.get('session_token')?.value;

    if (!sessionToken) {
      return {
        data: undefined,
        error: UNAUTHORIZED_ERROR,
      };
    }
    return { data: sessionToken, error: undefined };
  }

  async getUser(cookies: AstroCookies): ApiResponse<User> {
    const { data: sessionToken, error } =
      await this.validateSessionToken(cookies);

    if (error) {
      return {
        data: undefined,
        error: error,
      };
    }

    try {
      const response = await fetch(`${this.apiAuthUrl}/me`, {
        method: 'GET',
        headers: {
          Cookie: `session_token=${sessionToken}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
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

  async logout(cookies: AstroCookies): ApiResponse<void> {
    const { data: sessionToken, error } =
      await this.validateSessionToken(cookies);

    if (error) {
      return {
        data: undefined,
        error: error,
      };
    }

    try {
      const response = await fetch(`${this.apiAuthUrl}/logout`, {
        method: 'POST',
        headers: {
          Cookie: `session_token=${sessionToken}`,
        },
      });

      if (!response.ok) {
        console.log(response);
        const errorData = await response.json();
        console.error('[authService]', errorData);
        return {
          data: undefined,
          error: {
            status: response.status,
            message: errorData.message || 'Failed to logout',
          },
        };
      }

      cookies.delete('session_token', {
        httpOnly: true,
        sameSite: 'lax',
        path: '/',
      });
      return { data: undefined, error: undefined };
    } catch (error) {
      console.error('Error logging out:', error);
      return {
        data: undefined,
        error: INTERNAL_SERVER_ERROR,
      };
    }
  }
}

export const authService = new AuthService();
