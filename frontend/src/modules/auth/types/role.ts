import type { Roles } from '@/modules/auth/lib/roles';

export type Role = (typeof Roles)[keyof typeof Roles];
