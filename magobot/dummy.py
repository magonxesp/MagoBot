from .helper import send, get_message_arg
import random


def send_welcome(bot, update):
    # envia el saludo y si el usuario esta en un grupo o supergrupo lo menciona
    send(bot, update, "Hola k ase", mention_user=True)


def roll(bot, update):
    max_random_value = 100
    arg = get_message_arg(update.message.text)

    if arg.isnumeric():
        max_random_value = int(arg)

    value = random.randint(0, max_random_value)
    send(bot, update, value)
