import type { Roles } from '@/modules/auth/core/lib/roles';
import { INTERNAL_SERVER_ERROR } from '@/modules/core/lib/errors';
import type { ApiResponse } from '@/modules/core/types/api';
import { INSTITUTION_MOCK } from '@/modules/dashboard/instituciones/core/lib/mock';
import type { Institution } from '@/modules/dashboard/instituciones/core/types/institution';
import { BACKEND_URL } from 'astro:env/client';

class InstitutionService {
  private apiInstitutionUrl: string;

  constructor() {
    this.apiInstitutionUrl = `${BACKEND_URL}/institutions`;
  }

  async getInstitutionsGroupedByRoleMock(userId: string): ApiResponse<{
    [Roles.ADMIN]: Institution[];
    [Roles.STUDENT]: Institution[];
  }> {
    const l = INSTITUTION_MOCK.length;
    console.log('[getInstitutionsGroupedByRoleMock]: userId --> ', userId);
    return {
      data: {
        admin: INSTITUTION_MOCK.toSpliced(0, l / 2),
        student: INSTITUTION_MOCK.toSpliced(l / 2 + 1),
      },
      error: undefined,
    };
  }

  async getInstitutionsGroupedByRole(userId: string): ApiResponse<{
    [Roles.ADMIN]: Institution[];
    [Roles.STUDENT]: Institution[];
  }> {
    try {
      const response = await fetch(
        // TODO: Change real API URL
        `${this.apiInstitutionUrl}/grouped/${userId}`,
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
            message: errorData.message || 'Failed to fetch institutions',
          },
        };
      }

      const data = await response.json();
      return { data, error: undefined };
    } catch (error) {
      console.log('[getInstitutionsGroupedByRole] Error:', error);
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

export const institutionService = () => new InstitutionService();
