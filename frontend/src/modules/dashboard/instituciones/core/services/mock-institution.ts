import type { Roles } from '@/modules/auth/core/lib/roles';
import { NOT_FOUND_ERROR } from '@/modules/core/lib/errors';
import type { ApiResponse } from '@/modules/core/types/api';
import { INSTITUTION_MOCK } from '@/modules/dashboard/instituciones/core/lib/mock';
import type { Institution } from '@/modules/dashboard/instituciones/core/types/institution';

class InstitutionService {
  async getInstitutionsGroupedByRole(userId: string): ApiResponse<{
    [Roles.ADMIN]: Institution[];
    [Roles.STUDENT]: Institution[];
  }> {
    const l = INSTITUTION_MOCK.length;
    console.log('[getInstitutionsGroupedByRoleMock]: userId --> ', userId);
    return {
      data: {
        admin: INSTITUTION_MOCK.toSpliced(0, l / 2),
        student: INSTITUTION_MOCK.toSpliced(l / 2),
      },
      error: undefined,
    };
  }

  async getInstitutionById(institutionId: string): ApiResponse<Institution> {
    const i = INSTITUTION_MOCK.find((ins) => ins.id);
    console.log('[getInstitutionById] institutionId --> ', institutionId);
    if (!i) {
      return {
        data: undefined,
        error: NOT_FOUND_ERROR,
      };
    }
    return {
      data: i,
      error: undefined,
    };
  }
}

export const institutionService = () => new InstitutionService();
