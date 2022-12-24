# MagoBot
MagoBot, un bot magico para telegram.

Requiere Go >= 1.19

## Instalacion
1. Clonar este repositorio
    ```shell script
    $ git clone https://github.com/MagonxESP/MagoBot.git && cd MagoBot
    ```
2. Crear el fichero ``.env``
    ```shell script
    $ cp .env.template .env
    ```
3. Añadir la variable de entorno ``TOKEN`` en el fichero ``.env``
    ```shell script
    TOKEN=bot_token
    ```
## Arranque del bot   
### Con docker (recomendado)
* Crear los contenedores usando ``docker-compose`` e iniciar el bot
    ```shell script
    $ docker-compose up -d --build
    ```
### Sin docker, compilando y ejecutando
1. Compilar con make
    ```shell script
    $ make build
    ```

2. Iniciar el bot
    ```shell script
    ./build/magobot
    ```
 