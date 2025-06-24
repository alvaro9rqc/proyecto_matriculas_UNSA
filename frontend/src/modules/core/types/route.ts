import type { AstroComponent } from '@lucide/astro';

type RouteBase = {
  name: string;
  title: string;
  sub?: Route[];
  icon?: AstroComponent;
};
export type Route = RouteBase &
  (
    | {
        path: string;
        dynamic: undefined;
      }
    | {
        path: undefined;
        dynamic: true;
      }
  );
