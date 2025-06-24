import { getRouteFromPath } from '@/modules/core/lib/routes';
import type { Route } from '@/modules/core/types/route';

export const LOGIN_ROUTE = getRouteFromPath('/auth/login') || ({} as Route);
