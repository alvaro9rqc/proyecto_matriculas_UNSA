import type { Route } from '@/modules/core/types/route';
import { routesConfig } from '../config/routes';

export function getRouteFromPath(path: string): Route | undefined {
  console.log('Get from path: ', path);
  // Split path and filter out empty segments
  const splittedPath = path.split('/').filter(Boolean);
  // Map segments to Route object
  return splittedPath.reduce<Route | undefined>((currentRoute, segment) => {
    if (!currentRoute) return undefined;
    const next = currentRoute.sub?.find((r) => r.path === segment || r.dynamic);
    return next;
  }, routesConfig);
}
