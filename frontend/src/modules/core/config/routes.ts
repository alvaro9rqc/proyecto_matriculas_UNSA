import { House } from '@lucide/astro';
import type { Route } from '../types/route';

export const routesConfig: Route = {
  name: 'landing',
  title: 'Tuitions - Administra tus matrículas',
  path: '',
  sub: [
    {
      name: 'not-found',
      title: 'Página no encontrada',
      path: 'not-found',
      fullPath: '/not-found',
    },
    {
      name: 'unauthorized',
      title: 'No autorizado',
      path: 'unauthorized',
      fullPath: '/unauthorized',
    },
    {
      name: 'dashboard',
      title: 'Panel de control',
      path: 'dashboard',
      icon: House,
      fullPath: '/dashboard',
      sub: [
        {
          name: 'dashboard-admin',
          title: 'Panel de administración',
          path: 'admin',
          fullPath: '/dashboard/admin',
          sub: [
            {
              name: 'dashboard-admin-instituciones',
              title: 'Instituciones - Panel de administración',
              path: 'instituciones',
              fullPath: '/dashboard/admin/instituciones',
              sub: [
                {
                  name: 'dashboard-admin-instituciones-crear',
                  title: 'Crear institución - Panel de administración',
                  path: 'crear',
                  fullPath: '/dashboard/admin/instituciones/crear',
                  sub: [
                    {
                      name: 'dashboard-admin-instituciones-crear-bienvenido',
                      title:
                        'Bienvenido - Crear institución - Panel de administración',
                      path: 'bienvenido',
                      fullPath:
                        '/dashboard/admin/instituciones/crear/bienvenido',
                    },
                    {
                      name: 'dashboard-admin-instituciones-crear-nombre',
                      title:
                        'Nombre - Crear institución - Panel de administración',
                      path: 'nombre',
                      fullPath: '/dashboard/admin/instituciones/crear/nombre',
                    },
                  ],
                },
                {
                  name: 'dashboard-admin-instituciones-one',
                  title: 'Institución - Panel de administración',
                  dynamic: true,
                  param: 'institucion_id',
                  fullPath: '/dashboard/admin/instituciones/:institucion_id',
                  sub: [
                    {
                      name: 'dashboard-admin-instituciones-one-cursos',
                      title: 'Cursos - Institución - Panel de administración',
                      path: 'cursos',
                      fullPath:
                        '/dashboard/admin/instituciones/:institucion_id/cursos',
                      sub: [
                        {
                          name: 'dashboard-admin-instituciones-one-cursos-one',
                          title:
                            'Curso - Institución - Panel de administración',
                          dynamic: true,
                          param: 'curso_id',
                          fullPath:
                            '/dashboard/admin/instituciones/:institucion_id/cursos/:curso_id',
                        },
                      ],
                    },
                    {
                      name: 'dashboard-admin-instituciones-one-estudiantes',
                      title:
                        'Estudiantes - Institución - Panel de administración',
                      path: 'estudiantes',
                      fullPath:
                        '/dashboard/admin/instituciones/:institucion_id/estudiantes',
                      sub: [
                        {
                          name: 'dashboard-admin-instituciones-one-estudiantes-one',
                          title:
                            'Estudiante - Institución - Panel de administración',
                          dynamic: true,
                          param: 'estudiante_id',
                          fullPath:
                            '/dashboard/admin/instituciones/:institucion_id/estudiantes/:estudiante_id',
                        },
                      ],
                    },
                    {
                      name: 'dashboard-admin-instituciones-one-matriculas',
                      title:
                        'Matrículas - Institución - Panel de administración',
                      path: 'matriculas',
                      fullPath:
                        '/dashboard/admin/instituciones/:institucion_id/matriculas',
                      sub: [
                        {
                          name: 'dashboard-admin-instituciones-one-matriculas-crear',
                          title:
                            'Crear matrícula - Institución - Panel de administración',
                          path: 'crear',
                          fullPath:
                            '/dashboard/admin/instituciones/:institucion_id/matriculas/crear',
                          sub: [
                            {
                              name: 'dashboard-admin-instituciones-one-matriculas-crear-inicio',
                              title:
                                'Inicio - Crear matrícula - Institución - Panel de administración',
                              path: 'inicio',
                              fullPath:
                                '/dashboard/admin/instituciones/:institucion_id/matriculas/crear/inicio',
                            },
                            {
                              name: 'dashboard-admin-instituciones-one-matriculas-crear-subirDatos',
                              title:
                                'Subir Datos - Crear matrícula - Institución - Panel de administración',
                              path: 'subir-datos',
                              fullPath:
                                '/dashboard/admin/instituciones/:institucion_id/matriculas/crear/subir-datos',
                            },
                            {
                              name: 'dashboard-admin-instituciones-one-matriculas-crear-listos',
                              title:
                                'Todo listo - Crear matrícula - Institución - Panel de administración',
                              path: 'listos',
                              fullPath:
                                '/dashboard/admin/instituciones/:institucion_id/matriculas/crear/listos',
                            },
                          ],
                        },
                        {
                          name: 'dashboard-admin-instituciones-one-matriculas-one',
                          title:
                            'Matrícula - Institución - Panel de administración',
                          dynamic: true,
                          param: 'matricula_id',
                          fullPath:
                            '/dashboard/admin/instituciones/:institucion_id/matriculas/:matricula_id',
                          sub: [
                            {
                              name: 'dashboard-admin-instituciones-one-matriculas-one-horarios',
                              title:
                                'Horarios - Matrícula - Institución - Panel de administración',
                              path: 'horarios',
                              fullPath:
                                '/dashboard/admin/instituciones/:institucion_id/matriculas/:matricula_id/horarios',
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
          fullPath: '/dashboard/estudiante',
          sub: [
            {
              name: 'dashboard-estudiante-instituciones',
              title: 'Instituciones - Panel de estudiante',
              path: 'instituciones',
              fullPath: '/dashboard/estudiante/instituciones',
              sub: [
                {
                  name: 'dashboard-estudiante-instituciones-one',
                  title: 'Institución - Panel de estudiante',
                  dynamic: true,
                  param: 'institucion_id',
                  fullPath:
                    '/dashboard/estudiante/instituciones/:institucion_id',
                  sub: [
                    {
                      name: 'dashboard-estudiante-instituciones-one-cursos',
                      title: 'Cursos - Institución - Panel de estudiante',
                      path: 'cursos',
                      fullPath:
                        '/dashboard/estudiante/instituciones/:institucion_id/cursos',
                      sub: [
                        {
                          name: 'dashboard-estudiante-instituciones-one-cursos-one',
                          title: 'Curso - Institución - Panel de estudiante',
                          dynamic: true,
                          param: 'curso_id',
                          fullPath:
                            '/dashboard/estudiante/instituciones/:institucion_id/cursos/:curso_id',
                        },
                      ],
                    },
                    {
                      name: 'dashboard-estudiante-instituciones-one-matriculas',
                      title: 'Matrículas - Institución - Panel de estudiante',
                      path: 'matriculas',
                      fullPath:
                        '/dashboard/estudiante/instituciones/:institucion_id/matriculas',
                      sub: [
                        {
                          name: 'dashboard-estudiante-instituciones-one-matriculas-one',
                          title:
                            'Matrícula - Institución - Panel de estudiante',
                          dynamic: true,
                          param: 'matricula_id',
                          fullPath:
                            '/dashboard/estudiante/instituciones/:institucion_id/matriculas/:matricula_id',
                          sub: [
                            {
                              name: 'dashboard-estudiante-instituciones-one-matriculas-one-constancia',
                              title:
                                'Constancia - Matrícula - Institución - Panel de estudiante',
                              path: 'constancia',
                              fullPath:
                                '/dashboard/estudiante/instituciones/:institucion_id/matriculas/:matricula_id/constancia',
                            },
                            {
                              name: 'dashboard-estudiante-instituciones-one-matriculas-one-horarios',
                              title:
                                'Horarios - Matrícula - Institución - Panel de estudiante',
                              path: 'horarios',
                              fullPath:
                                '/dashboard/estudiante/instituciones/:institucion_id/matriculas/:matricula_id/horarios',
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
  ],
};
