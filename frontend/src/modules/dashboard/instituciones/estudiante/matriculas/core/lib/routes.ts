import { getRouteFromPath } from '@/modules/core/lib/routes';

// Only for faster matching
export const DASHBOARD_STUDENT_INSTITUTION_ROUTE = getRouteFromPath(
  '/dashboard/estudiante/instituciones/:institution_id',
);
// Only for faster matching
export const DASHBOARD_STUDENT_INSTITUTION_ENROLLMENT_ROUTE = getRouteFromPath(
  '/dashboard/estudiante/instituciones/:institution_id/matriculas/:matricula_id',
);
