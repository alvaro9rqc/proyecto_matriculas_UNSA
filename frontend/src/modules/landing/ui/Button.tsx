import { cn } from '@/modules/core/lib/utils';
import type { ComponentPropsWithRef } from 'react';
import type { VariantProps } from 'tailwind-variants';
import { buttonVariants, type SupportedTags } from '../lib/buttonUtils';

type ButtonProps<Tag extends SupportedTags> = {
  as?: Tag;
} & VariantProps<typeof buttonVariants> &
  ComponentPropsWithRef<Tag>;
export default function Button<T extends SupportedTags = 'button'>({
  as = 'button' as T,
  variant,
  size,
  children,
  className,
  ...rest
}: ButtonProps<T>) {
  const Tag = as;
  return (
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    <Tag className={cn(buttonVariants({ variant, size }), className)} {...rest}>
      {children}
    </Tag>
  );
}
