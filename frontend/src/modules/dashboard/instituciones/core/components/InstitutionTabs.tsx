import { Roles } from '@/modules/auth/core/lib/roles';
import { Button } from '@/modules/core/ui/button';
import { Card, CardContent } from '@/modules/core/ui/card';
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from '@/modules/core/ui/tabs';
import type { Institution } from '@/modules/dashboard/instituciones/core/types/institution';

interface InstitutionTabsProps {
  studentInstitutions: Institution[];
  adminInstitutions: Institution[];
}

export function InstitutionTabs({
  adminInstitutions,
  studentInstitutions,
}: InstitutionTabsProps) {
  return (
    <Tabs defaultValue={Roles.STUDENT} className="w-64 sm:w-md md:w-xl flex-1">
      <TabsList className="w-full justify-center">
        <TabsTrigger value={Roles.STUDENT}>Estudiante</TabsTrigger>
        <TabsTrigger value={Roles.ADMIN}>Administrador</TabsTrigger>
      </TabsList>
      <TabsContent value={Roles.STUDENT} className="flex flex-col">
        <Card className="grow">
          <CardContent>
            {studentInstitutions.length > 0 ? (
              studentInstitutions.map((institution) => (
                <Button
                  key={institution.id}
                  className="items-center w-full"
                  variant="ghost"
                  size="lg"
                >
                  {institution.logoUrl ? (
                    <img
                      src={institution.logoUrl}
                      alt={institution.name}
                      className="h-full object-contain"
                    />
                  ) : (
                    <span className="h-full bg-gray-200 rounded-full flex items-center justify-center">
                      <span className="text-gray-500">No Logo</span>
                    </span>
                  )}
                  {institution.name}
                </Button>
              ))
            ) : (
              <p>No tienes instituciones asignadas.</p>
            )}
          </CardContent>
        </Card>
      </TabsContent>
      <TabsContent value={Roles.ADMIN} className="flex flex-col">
        <Card className="grow">
          <CardContent>
            {adminInstitutions.length > 0 ? (
              adminInstitutions.map((institution) => (
                <Button
                  key={institution.id}
                  className="items-center w-full"
                  variant="ghost"
                  size="lg"
                >
                  {institution.logoUrl ? (
                    <img
                      src={institution.logoUrl}
                      alt={institution.name}
                      className="h-full object-contain"
                    />
                  ) : (
                    <span className="h-full bg-gray-200 rounded-full flex items-center justify-center">
                      <span className="text-gray-500">No Logo</span>
                    </span>
                  )}
                  {institution.name}
                </Button>
              ))
            ) : (
              <p>No tienes instituciones asignadas.</p>
            )}
          </CardContent>
        </Card>
      </TabsContent>
    </Tabs>
  );
}
