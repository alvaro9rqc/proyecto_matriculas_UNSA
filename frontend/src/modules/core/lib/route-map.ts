import type { Route } from '@/modules/core/types/route';
import { routesConfig } from '@/modules/core/config/routes';

type RouteMap = Map<string, Route>;

const routeMap: RouteMap = new Map();

function buildFullPath(route: Route, basePath = ''): string {
  const currentPath = route.dynamic ? `:${route.param}` : route.path;
  return [basePath, currentPath].filter(Boolean).join('/');
}

function flattenRoutes(route: Route, basePath = '') {
  const fullPath = buildFullPath(route, basePath);
  routeMap.set(`/${fullPath}`, { ...route, fullPath });

  for (const sub of route.sub || []) {
    flattenRoutes(sub, fullPath);
  }
}

flattenRoutes(routesConfig);

export function getFlattenedRouteMap(): RouteMap {
  return routeMap;
}
