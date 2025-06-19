import { SupportedFields } from '@/modules/core/lib/field';

export type SupportFieldType =
  (typeof SupportedFields)[keyof typeof SupportedFields];

type BaseField<T extends string> = {
  name: T;
  label: string;
  description?: string;
  placeholder?: string;
  required?: boolean;
};

export type Field<T extends string> = BaseField<T> &
  (
    | {
        type:
          | typeof SupportedFields.EMAIL
          | typeof SupportedFields.PASSWORD
          | typeof SupportedFields.TEXTAREA
          | typeof SupportedFields.TEXT;
      }
    | {
        type: typeof SupportedFields.SELECT;
        options: string[];
      }
  );
