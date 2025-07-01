import type { EnrollmentProcessStates } from '@/modules/dashboard/instituciones/matriculas/core/lib/enrollment';

export type EnrollmentProcessStatus =
  (typeof EnrollmentProcessStates)[keyof typeof EnrollmentProcessStates];

type BaseEnrollmentProcess = {
  id: string;
  institutionId: string;
};

export type EnrollmentProcess = BaseEnrollmentProcess & {
  name: string;
  startAt: string;
  endAt: string;
};
