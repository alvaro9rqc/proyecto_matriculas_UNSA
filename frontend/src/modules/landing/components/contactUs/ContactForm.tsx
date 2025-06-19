import type { Field } from '@/modules/core/types/field';
import { Form, FormField } from '@/modules/core/ui/form';
import { useForm, type SubmitHandler } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { InferItem } from '@/modules/core/ui/inferfield';

const contactFormSchema = z.object({
  names: z.string().min(1, 'Name is required'),
  lastName: z.string().min(1, 'Last name is required'),
  email: z.string().email('Invalid email address'),
  institutionName: z.string().min(1, 'Institution name is required'),
  message: z.string().min(1, 'Message is required'),
});

const contactFormFields: Field<keyof z.infer<typeof contactFormSchema>>[] = [
  {
    name: 'names',
    label: 'First Name',
    type: 'text',
    placeholder: 'Enter your first name',
    required: true,
  },
  {
    name: 'lastName',
    label: 'Last Name',
    type: 'text',
    placeholder: 'Enter your last name',
    required: true,
  },
  {
    name: 'email',
    label: 'Email',
    placeholder: 'Enter your email address',
    type: 'email',
    required: true,
  },
  {
    name: 'institutionName',
    label: 'Institution Name',
    type: 'text',
    placeholder: 'Enter your institution name',
    required: true,
  },
  {
    name: 'message',
    label: 'Message',
    placeholder: 'Enter your message',
    type: 'textarea',
    required: true,
  },
];

function ContactForm() {
  const form = useForm<z.infer<typeof contactFormSchema>>({
    resolver: zodResolver(contactFormSchema),
    defaultValues: {
      names: '',
      lastName: '',
      email: '',
      institutionName: '',
      message: '',
    },
  });
  const onSubmit: SubmitHandler<z.infer<typeof contactFormSchema>> = (
    values,
  ) => {
    console.log(values);
  };
  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        {contactFormFields.map((field) => (
          <FormField
            control={form.control}
            name={field.name}
            render={({ field: formField }) => (
              <InferItem {...field} {...formField} />
            )}
          />
        ))}
      </form>
    </Form>
  );
}

export { ContactForm };
