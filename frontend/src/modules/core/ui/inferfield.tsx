import type { Field } from '@/modules/core/types/field';
import { Input } from '@/modules/core/ui/input';
import {
  FormControl,
  FormDescription,
  FormItem,
  FormLabel,
} from '@/modules/core/ui/form';
import type {
  ControllerRenderProps,
  FieldPath,
  FieldValues,
} from 'react-hook-form';
import { SupportedFields } from '../lib/field';
import { Select } from './select';

function InferItem<
  FieldName extends string,
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>,
>({
  label,
  description,
  ...props
}: Field<FieldName> & ControllerRenderProps<TFieldValues, TName>) {
  return (
    <FormItem>
      <FormLabel>{label}</FormLabel>
      <FormControl>
        {
          props.type === SupportedFields.SELECT && (
            <Select>
              {
                props.options
              }
            </Select>
          )
        }
        {
          props.type === SupportedFields.TEXTAREA && (

          )
        }
        {
          props.type !== SupportedFields.SELECT && props.type !== SupportedFields.TEXTAREA && (

          )
        }
      </FormControl>
      {description && <FormDescription>{description}</FormDescription>}
    </FormItem>
  );
}

export { InferItem };
