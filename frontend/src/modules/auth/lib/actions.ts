import { defineAction } from 'astro:actions';
import { getUser } from '@/modules/auth/lib/authServer';

export const user = {
  getUser: defineAction({ handler: () => getUser() }),
};
