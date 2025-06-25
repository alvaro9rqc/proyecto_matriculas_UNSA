import { defineAction } from 'astro:actions';
import { authService } from '@/modules/auth/services';

export const user = {
  getUser: defineAction({
    handler: (_, context) => authService(context.cookies).getUser(),
  }),
};
