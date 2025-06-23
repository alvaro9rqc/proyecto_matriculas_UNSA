import { defineAction } from 'astro:actions';
import { authApi } from '@/modules/auth/lib/api';

export const user = {
  getUser: defineAction({ handler: () => authApi.getUser() }),
};
