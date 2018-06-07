from telegram.ext import Updater, CommandHandler
import rule34
import random
import sys


def get_rule34_post(tag):
    try:
        images_urls = rule34.getImageURLS(tag)

        if images_urls is not None:
            index = random.randint(0, len(images_urls))
            return images_urls[index]
        else:
            return None
    except rule34.Request_Rejected or rule34.Rule34_Error:
        print(sys.exc_info())

    return None


def get_message_arg(message_text):
    message_array = message_text.split(' ')[1:]

    arg = ""  # message argument string
    for message in message_array:
        arg += message + ' '

    arg = arg.rstrip()
    return arg


def send_rule34(bot, update):
    arg = get_message_arg(update.message.text)
    bot.send_sticker(update.message.chat_id, "CAADAwAD0RAAAsKphwW8SGCoG75C8AI")  # sticker de kanna buscando
    image_url = get_rule34_post(arg)

    if image_url is not None:
        bot.send_message(update.message.chat_id, image_url)
    else:
        bot.send_message(update.message.chat_id, "Prueba con otra cosa...")


def send_welcome(bot, update):
    bot.send_message(update.message.chat_id, "Hola k ase")


def roll(bot, update):
    max_random_value = 1000
    arg = get_message_arg(update.message.text)

    if arg.isnumeric():
        max_random_value = int(arg)

    value = random.randint(0, max_random_value)
    bot.send_message(chat_id=update.message.chat_id, text=value, reply_to_message_id=update.message.message_id)


# crea el updater del bot mediante el token del bot
updater = Updater("593801508:AAF0qCsRxbyKG0I-QSCoh4wwW7A-G6HuccU")

# añade los comandos al bot
updater.dispatcher.add_handler(CommandHandler('start', send_welcome))
updater.dispatcher.add_handler(CommandHandler('rule34', send_rule34))
updater.dispatcher.add_handler(CommandHandler('roll', roll))

updater.start_polling()
updater.idle()
