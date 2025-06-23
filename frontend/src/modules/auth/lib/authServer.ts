import type { User } from '@/modules/auth/types/user';
import { timeout } from '@/modules/core/lib/time';

export async function getUser(): Promise<User | undefined> {
  // Simulate a server request to get the user
  await timeout(1000);
  const user: User = {
    id: '1',
    name: 'Yenaro Joel',
    surname: 'Noa Camino',
    email: 'ynoacamino@unsa.edu.pe',
    avatar_url:
      'https://ynoa-uploader.ynoacamino.site/uploads/1750016704_ACg8ocLnHIiNMcd-ltRxMAQZ6Qo1hKAeSyZsktQKBp5kNltpKDzlg4_q=s96-c.webp',
  };
  // If return undefined, the user will be considered as not logged in
  return user;
}
