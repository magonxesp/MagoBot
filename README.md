# MagoBot
MagoBot, un bot magico para telegram.

Requiere python >= 3.7

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
### Sin docker, usando virtualenv
1. Crear el virtualenv y activarlo
    ```shell script
    $ pip install virtualenv
    $ virtualenv venv
    $ source venv/bin/activate
    ```
2. Instalar las dependencias
    ```shell script
    (venv) $ pip install -r requirements.txt
    ```
3. Iniciar el bot
    ```shell script
    (venv) $ python -m magobot
    ```
 