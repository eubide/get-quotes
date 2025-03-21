# Random Line

Una herramienta simple en Go para extraer líneas aleatorias de archivos de texto.

## Estructura del Proyecto

```
get-quote/
├── bin/                    # Directorio de salida de compilación
├── cmd/                    # Comandos ejecutables
│   └── get-quote/        # Comando principal
├── pkg/                    # Código de paquetes
│   ├── config/             # Manejo de configuración
│   └── randomline/         # Funcionalidad de línea aleatoria
├── src/                    # Archivos fuente
│   └── files/              # Archivos de datos
│       ├── quotes.lst      # Citas en inglés
│       └── citas.lst       # Citas en español
├── go.mod                  # Definición del módulo Go
├── go.sum                  # Checksums de dependencias
├── Makefile                # Automatización de construcción
├── .get-quote.yaml    # Archivo de configuración
└── README.md               # Documentación del proyecto
```

## Instalación

```bash
# Clonar el repositorio
git clone https://github.com/eubide/get-quote.git
cd get-quote

# Construir el ejecutable
make build
```

## Uso

```bash
./bin/get-quote nombre_fichero
```

Donde `nombre_fichero` es el nombre de un archivo en el directorio configurado (por defecto `src/files/`).
La extensión `.lst` se añadirá automáticamente si no se proporciona.

## Ejemplos

Extraer una línea aleatoria del archivo quotes.lst:
```bash
./bin/get-quote quotes
```

Extraer una línea aleatoria del archivo citas.lst:
```bash
./bin/get-quote citas
```

## Configuración

La aplicación utiliza un archivo de configuración YAML llamado `.get-quote.yaml`. El archivo de configuración puede colocarse en:
1. El directorio actual
2. `$HOME/.config/.get-quote/`
3. El directorio home del usuario

Si no se encuentra ningún archivo de configuración, se utilizan valores predeterminados.

### Ejemplo de configuración

```yaml
# Configuración de Random Sentence

filesBaseDir: src/files
defaultExtension: .lst
errorMessages:
  fileNotFound: "El archivo %s no existe"
  fileOpenError: "Error al abrir el archivo: %v"
  missingParameter: "Uso: %s <nombre_fichero>\nDebe proporcionar un nombre de fichero %s"
```
