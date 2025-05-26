#!/usr/bin/bash

directory="${1%/*}"
file="$1"
base_name=$(basename "$file" )
echo "RUNNING ON:"
echo "$directory"
echo "###########"
if [[ "$directory" == "/home/alvaro9rqc/1_Pacha/1-unsa/5_S/ips/proyecto_matriculas_UNSA/backend/src/tests" ]]
then
  build1="\\i /home/alvaro9rqc/1_Pacha/1-unsa/5_S/ips/proyecto_matriculas_UNSA/backend/src/schema/schema.sql"
  if [[ "$base_name" == "schema.sql"  ]] 
  then
    build1=""
  fi
  echo "\\echo \\set ON_ERROR_STOP on 
  BEGIN;
  $build1
  \\i $file
  ROLLBACK;" | psql ips_1
fi

