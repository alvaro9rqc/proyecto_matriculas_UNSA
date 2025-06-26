import { defineAction } from 'astro:actions';
import { authService } from '@/modules/auth/services';

export const user = {
  getUser: defineAction({
    handler: async (_, context) => {
      const { data, error } = await authService(context.cookies).getUser();
      if (error) {
        throw error;
      }
      return data;
    },
  }),
  logout: defineAction({
    handler: async (_, context) => {
      const { data, error } = await authService(context.cookies).logout();
      if (error) {
        throw error;
      }
      return data;
    },
  }),
};
