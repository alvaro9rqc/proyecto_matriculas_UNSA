import type { Route } from '@/modules/core/types/route';
import { routesConfig } from '@/modules/core/config/routes';

export function getRouteFromPath(path: string): Route {
  console.log('Get from path: ', path);
  // Split path and filter out empty segments
  const splittedPath = path.split('/').filter(Boolean);
  // Map segments to Route object
  const result = splittedPath.reduce<Route | undefined>(
    (currentRoute, segment) => {
      if (!currentRoute) return undefined;
      const next = currentRoute.sub?.find(
        (r) => r.path === segment || r.dynamic,
      );
      return next;
    },
    routesConfig,
  );

  return result ?? ({} as Route);
}

export const NOT_FOUND_ROUTE = getRouteFromPath('/404');
