import type { OAuthProviderName } from '@/modules/auth/types/oauthProvider';
import { BACKEND_URL } from 'astro:env/client';

export function getLoginURL(provider: OAuthProviderName) {
  return `${BACKEND_URL}/auth/${provider}/login`;
}
export async function getLogoutURL() {
  return `${BACKEND_URL}/auth/logout`;
}
