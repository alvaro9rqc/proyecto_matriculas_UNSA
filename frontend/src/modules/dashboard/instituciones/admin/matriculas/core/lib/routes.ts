import { getRouteFromPath } from '@/modules/core/lib/routes';

// Only for faster matching
export const DASHBOARD_ADMIN_INSTITUTION_ROUTE = getRouteFromPath(
  '/dashboard/admin/instituciones/:institution_id',
);

// Only for faster matching
export const DASHBOARD_ADMIN_INSTITUTION_ENROLLMENT_ROUTE = getRouteFromPath(
  '/dashboard/admin/instituciones/:institution_id/matriculas/:matricula_id',
);
