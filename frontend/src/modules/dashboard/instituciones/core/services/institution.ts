import { authService } from '@/modules/auth/core/services/auth';
import { INTERNAL_SERVER_ERROR } from '@/modules/core/lib/errors';
import type { ApiResponse } from '@/modules/core/types/api';
import type { Institution } from '@/modules/dashboard/instituciones/core/types/institution';
import type { AstroCookies } from 'astro';
import { BACKEND_URL } from 'astro:env/client';

class InstitutionService {
  private apiInstitutionUrl: string;

  constructor() {
    this.apiInstitutionUrl = `${BACKEND_URL}/institutions`;
  }

  async getInstitutions(cookies: AstroCookies): ApiResponse<Institution[]> {
    const { data: sessionToken, error } =
      await authService.validateSessionToken(cookies);

    if (error) {
      return {
        data: undefined,
        error: error,
      };
    }

    try {
      const response = await fetch(
        // TODO: Change real API URL
        this.apiInstitutionUrl,
        {
          method: 'GET',
          headers: {
            Cookie: `session_token=${sessionToken}`,
          },
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        return {
          data: undefined,
          error: {
            status: response.status,
            message: errorData.message || 'Failed to fetch institutions',
          },
        };
      }

      const data = await response.json();
      return { data, error: undefined };
    } catch (error) {
      console.log('[getInstitutions] Error:', error);
      return {
        data: undefined,
        error: INTERNAL_SERVER_ERROR,
      };
    }
  }

  async getInstitutionById(institutionId: string): ApiResponse<Institution> {
    try {
      const response = await fetch(
        `${this.apiInstitutionUrl}/${institutionId}`,
        {
          method: 'GET',
        },
      );

      if (!response.ok) {
        const errorData = await response.json();
        return {
          data: undefined,
          error: {
            status: response.status,
            message: errorData.message || 'Failed to fetch institution',
          },
        };
      }

      const data = await response.json();
      return { data, error: undefined };
    } catch (error) {
      console.log('[getInstitutionById] Error:', error);
      return {
        data: undefined,
        error: INTERNAL_SERVER_ERROR,
      };
    }
  }
}

export const institutionService = new InstitutionService();
