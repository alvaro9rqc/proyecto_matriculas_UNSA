import { tv } from 'tailwind-variants';

export const buttonVariants = tv({
  base: 'rounded-md px-3 py-1 font-medium transition-colors duration-300 cursor-pointer w-max',
  variants: {
    variant: {
      primary: 'bg-blue-600 text-white hover:bg-blue-800',
      secondary: 'bg-gray-200 text-black hover:bg-gray-300',
      muted: 'text-black hover:bg-gray-300',
    },
    size: {
      sm: 'text-sm',
      md: 'text-base',
      lg: 'text-lg px-4 py-2',
    },
  },
  defaultVariants: {
    variant: 'primary',
    size: 'md',
  },
});

export type SupportedTags = 'button' | 'a';
