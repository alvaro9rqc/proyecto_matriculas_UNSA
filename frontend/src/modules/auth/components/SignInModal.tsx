import GoogleLogo from '@/assets/google.svg?react';
import Logo from '@/assets/logo.svg?react';
import MicrosoftLogo from '@/assets/microsoft.svg?react';
import { Button } from '@/modules/core/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/modules/core/ui/dialog';

export default function SignInModal() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Iniciar sesión</Button>
      </DialogTrigger>
      <DialogContent className="p-10 gap-y-14">
        <DialogHeader className="flex flex-col items-center gap-y-4">
          <DialogTitle className="inline-flex items-center gap-x-2">
            <Logo className="h-8" />
            <span className="text-2xl font-semibold">Tuitions</span>
          </DialogTitle>
          <DialogDescription>
            Ingresa a tu cuenta y revisa tus matrículas
          </DialogDescription>
        </DialogHeader>
        <div className="flex flex-col gap-y-4">
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
        </div>
      </DialogContent>
    </Dialog>
  );
}
