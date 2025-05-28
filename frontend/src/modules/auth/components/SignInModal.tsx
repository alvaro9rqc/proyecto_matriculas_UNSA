import GoogleLogo from '@/assets/google.svg?react';
import Logo from '@/assets/logo.svg?react';
import MicrosoftLogo from '@/assets/microsoft.svg?react';
import Button from '@/modules/landing/ui/Button';
import * as Dialog from '@radix-ui/react-dialog';

export default function SignInModal() {
  return (
    <Dialog.Root>
      <Dialog.Trigger asChild>
        <Button>Iniciar sesión</Button>
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="fixed inset-0 bg-black/25 animate-fade-in" />
        <Dialog.Content className="bg-white rounded-lg shadow-lg fixed top-1/2 left-1/2 -translate-1/2 p-4 md:p-8 flex flex-col gap-y-4 items-center w-[90vw] max-w-xl focus:outline-none animate-content-show">
          <Dialog.Title className="inline-flex items-center gap-x-2">
            <Logo className="h-8" />
            <span className="text-2xl font-semibold">Tuitions</span>
          </Dialog.Title>
          <Dialog.Description className="text-center text-balance text-md mb-8">
            Ingresa a tu cuenta y revisa tus matrículas
          </Dialog.Description>
          <Button
            variant="secondary"
            className="inline-flex justify-center items-center gap-x-2 w-full"
          >
            <GoogleLogo className="size-4" />
            Google
          </Button>
          <Button
            variant="secondary"
            className="inline-flex justify-center items-center gap-x-2 w-full"
          >
            <MicrosoftLogo className="size-4" />
            Microsoft
          </Button>
          <Dialog.Close asChild>
            <Button
              variant="muted"
              className="absolute top-2 right-2 md:top-4 md:right-4"
            >
              X
            </Button>
          </Dialog.Close>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}
