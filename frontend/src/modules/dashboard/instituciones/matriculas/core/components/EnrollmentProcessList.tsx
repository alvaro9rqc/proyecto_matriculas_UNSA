import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/modules/core/ui/card';
import { Link } from '@/modules/core/ui/link';
import type { Institution } from '@/modules/dashboard/instituciones/core/types/institution';
import { getEnrollmentProcessPath } from '@/modules/dashboard/instituciones/matriculas/core/lib/routes';
import type { EnrollmentProcess } from '@/modules/dashboard/instituciones/matriculas/core/types/process';
import { LANDING_ROUTE } from '@/modules/landing/core/lib/routes';
import { ArrowLeftIcon } from 'lucide-react';

interface EnrollmentProcessListProps {
  institution: Institution;
  processes: EnrollmentProcess[];
}

export function EnrollmentProcessList({
  institution,
  processes,
}: EnrollmentProcessListProps) {
  return (
    <Card className="w-64 sm:w-md md:w-xl grow">
      <CardHeader className="gap-x-4">
        <CardAction className="col-start-1 row-span-2">
          <Link variant="ghost" href={LANDING_ROUTE.fullPath}>
            <ArrowLeftIcon />
          </Link>
        </CardAction>
        <CardTitle>Procesos de matrícula en {institution.name}</CardTitle>
        <CardDescription>
          Selecciona un proceso y revisa todos sus detalles
        </CardDescription>
      </CardHeader>
      <CardContent>
        {processes.length > 0 ? (
          processes.map((process) => (
            <Link
              key={process.id}
              href={getEnrollmentProcessPath({
                institucion_id: institution.id,
                matricula_id: process.id,
              })}
              className="items-center w-full"
              variant="ghost"
              size="lg"
            >
              {process.startAt && process.endAt ? (
                <span className="text-sm text-gray-500">
                  {new Date(process.startAt).toLocaleDateString()} -{' '}
                  {new Date(process.endAt).toLocaleDateString()}
                </span>
              ) : null}
              {process.name}
            </Link>
          ))
        ) : (
          <p>No tienes instituciones asignadas.</p>
        )}
      </CardContent>
    </Card>
  );
}
