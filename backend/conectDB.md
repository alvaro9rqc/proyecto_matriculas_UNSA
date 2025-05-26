Para migrar la base de datos:

```bash
goose up
```

---

### ✅ Paso 2: Ejecutar la migración

Una vez configuradas las variables, ejecuta la migración:

```bash
goose up
```

Esto aplicará todas las migraciones que no han sido ejecutadas aún.

---

### ✅ Otros comandos útiles

* Ver el estado de las migraciones:

```bash
goose status
```

* Migrar a una versión específica:

```bash
goose up-to 20240525123000
```

Hacer rollback:

```bash
goose down
```

Crear una nueva migración (archivo `.sql`):

```bash
goose create nombre_migracion sql
```

Esto generará un archivo como:
`20250525154532_nombre_migracion.sql`

---
