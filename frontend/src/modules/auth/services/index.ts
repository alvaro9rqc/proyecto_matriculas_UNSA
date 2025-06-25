import { BACKEND_URL } from 'astro:env/client';
import type { User } from '@/modules/auth/types/user';
import type { ApiResponse } from '@/modules/core/types/api';
import {
  INTERNAL_SERVER_ERROR,
  UNAUTHORIZED_ERROR,
} from '@/modules/core/lib/errors';
import type { AstroCookies } from 'astro';

class AuthService {
  private apiAuthUrl: string;
  private astroCookies: AstroCookies;

  constructor(astroCookies: AstroCookies) {
    this.apiAuthUrl = `${BACKEND_URL}/auth`;
    this.astroCookies = astroCookies;
  }

  async getUser(): ApiResponse<User> {
    const sessionToken = this.astroCookies.get('session_token')?.value;

    if (!sessionToken) {
      return {
        data: undefined,
        error: UNAUTHORIZED_ERROR,
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

export const authService = (astroCookies: AstroCookies) =>
  new AuthService(astroCookies);
