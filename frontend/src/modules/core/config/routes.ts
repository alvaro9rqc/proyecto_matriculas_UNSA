import { House } from '@lucide/astro';
import type { Route } from '../types/route';

export const routesConfig: Route[] = [
  {
    name: 'landing',
    title: 'Tuitions - Administra tus matrículas',
    path: '',
  },
  {
    name: 'dashboard',
    title: 'Panel de control',
    path: 'dashboard',
    icon: House,
    sub: [
      {
        name: 'dashboard-admin',
        title: 'Panel de administración',
        path: 'admin',
        sub: [
          {
            name: 'dashboard-admin-instituciones',
            title: 'Instituciones - Panel de administración',
            path: 'instituciones',
            sub: [
              {
                name: 'dashboard-admin-instituciones-crear',
                title: 'Crear institución - Panel de administración',
                path: 'crear',
                sub: [
                  {
                    name: 'dashboard-admin-instituciones-crear-bienvenido',
                    title:
                      'Bienvenido - Crear institución - Panel de administración',
                    path: 'bienvenido',
                  },
                  {
                    name: 'dashboard-admin-instituciones-crear-nombre',
                    title:
                      'Nombre - Crear institución - Panel de administración',
                    path: 'nombre',
                  },
                ],
              },
              {
                name: 'dashboard-admin-instituciones-one',
                title: 'Institución - Panel de administración',
                dynamic: true,
                sub: [
                  {
                    name: 'dashboard-admin-instituciones-one-cursos',
                    title: 'Cursos - Institución - Panel de administración',
                    path: 'cursos',
                    sub: [
                      {
                        name: 'dashboard-admin-instituciones-one-cursos-one',
                        title: 'Curso - Institución - Panel de administración',
                        dynamic: true,
                      },
                    ],
                  },
                  {
                    name: 'dashboard-admin-instituciones-one-estudiantes',
                    title:
                      'Estudiantes - Institución - Panel de administración',
                    path: 'estudiantes',
                    sub: [
                      {
                        name: 'dashboard-admin-instituciones-one-estudiantes-one',
                        title:
                          'Estudiante - Institución - Panel de administración',
                        dynamic: true,
                      },
                    ],
                  },
                  {
                    name: 'dashboard-admin-instituciones-one-matriculas',
                    title: 'Matrículas - Institución - Panel de administración',
                    path: 'matriculas',
                    sub: [
                      {
                        name: 'dashboard-admin-instituciones-one-matriculas-crear',
                        title:
                          'Crear matrícula - Institución - Panel de administración',
                        path: 'crear',
                        sub: [
                          {
                            name: 'dashboard-admin-instituciones-one-matriculas-crear-inicio',
                            title:
                              'Inicio - Crear matrícula - Institución - Panel de administración',
                            path: 'inicio',
                          },
                          {
                            name: 'dashboard-admin-instituciones-one-matriculas-crear-subirDatos',
                            title:
                              'Subir Datos - Crear matrícula - Institución - Panel de administración',
                            path: 'subir-datos',
                          },
                          {
                            name: 'dashboard-admin-instituciones-one-matriculas-crear-listos',
                            title:
                              'Todo listo - Crear matrícula - Institución - Panel de administración',
                            path: 'listos',
                          },
                        ],
                      },
                      {
                        name: 'dashboard-admin-instituciones-one-matriculas-one',
                        title:
                          'Matrícula - Institución - Panel de administración',
                        dynamic: true,
                        sub: [
                          {
                            name: 'dashboard-admin-instituciones-one-matriculas-one-horarios',
                            title:
                              'Horarios - Matrícula - Institución - Panel de administración',
                            path: 'horarios',
                          },
                        ],
                      },
                    ],
                  },
                ],
              },
            ],
          },
        ],
      },
      {
        name: 'dashboard-estudiante',
        title: 'Panel de estudiante',
        path: 'estudiante',
        sub: [
          {
            name: 'dashboard-estudiante-instituciones',
            title: 'Instituciones - Panel de estudiante',
            path: 'instituciones',
            sub: [
              {
                name: 'dashboard-estudiante-instituciones-one',
                title: 'Institución - Panel de estudiante',
                dynamic: true,
                sub: [
                  {
                    name: 'dashboard-estudiante-instituciones-one-cursos',
                    title: 'Cursos - Institución - Panel de estudiante',
                    path: 'cursos',
                    sub: [
                      {
                        name: 'dashboard-estudiante-instituciones-one-cursos-one',
                        title: 'Curso - Institución - Panel de estudiante',
                        dynamic: true,
                      },
                    ],
                  },
                  {
                    name: 'dashboard-estudiante-instituciones-one-matriculas',
                    title: 'Matrículas - Institución - Panel de estudiante',
                    path: 'matriculas',
                    sub: [
                      {
                        name: 'dashboard-estudiante-instituciones-one-matriculas-one',
                        title: 'Matrícula - Institución - Panel de estudiante',
                        dynamic: true,
                        sub: [
                          {
                            name: 'dashboard-estudiante-instituciones-one-matriculas-one-constancia',
                            title:
                              'Constancia - Matrícula - Institución - Panel de estudiante',
                            path: 'constancia',
                          },
                          {
                            name: 'dashboard-estudiante-instituciones-one-matriculas-one-horarios',
                            title:
                              'Horarios - Matrícula - Institución - Panel de estudiante',
                            path: 'horarios',
                          },
                        ],
                      },
                    ],
                  },
                ],
              },
            ],
          },
        ],
      },
    ],
  },
] as const;
