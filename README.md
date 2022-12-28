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
3. Añadir la variable de entorno ``MAGOBOT_TOKEN`` en el fichero ``.env``
    ```shell script
    MAGOBOT_TOKEN=bot_token
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

## Comandos

* **/start** - MagoBot te saluda.
* **/roll** - Genera un numero aleatorio del 0 al 100, se puede cambiar el numero maximo con /roll 50 por ejemplo.
* **/rule34** - Muestra un post aleatorio de rule34 de un tag, por ejemplo /rule34 tag.
* **/4chanrandomwthread** - Muestra un thread aleatorio de 4chan del board W ademas de una imagen aleatoria de algun post del thread aleatorio.
* **/4chanrandombthread** - Muestra un thread aleatorio de 4chan del board B.
* **/4chanrandomhentaithread** - Muestra un thread aleatorio de 4chan del board H ademas de una imagen aleatoria de algun post del thread aleatorio.
* **/4chanrandomecchithread** - Muestra un thread aleatorio de 4chan del board E ademas de una imagen aleatoria de algun post del thread aleatorio.
