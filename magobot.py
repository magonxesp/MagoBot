from telegram.ext import Updater, CommandHandler
from telegram.chat import Chat
from telegram.parsemode import ParseMode
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


def send(bot, update, message_text, no_reply_in_group=False, mention_user=False):
    chat = update.message.chat
    message = update.message

    if chat.type == Chat.GROUP or chat.type == Chat.SUPERGROUP:
        if no_reply_in_group:
            bot.send_message(chat_id=chat.id, text=message_text)
        else:
            if mention_user:
                user = message.from_user
                markdown = "[@{}](tg://user?id={}) " + message_text
                message_text = markdown.format(user.username, user.id)
                bot.send_message(chat_id=chat.id, text=message_text, parse_mode=ParseMode.MARKDOWN)
            else:
                bot.send_message(chat_id=chat.id, text=message_text, reply_to_message_id=message.message_id)
    else:
        bot.send_message(chat.id, message_text)


def send_rule34(bot, update):
    arg = get_message_arg(update.message.text)
    bot.send_sticker(update.message.chat_id, "CAADAwAD0RAAAsKphwW8SGCoG75C8AI")  # sticker de kanna buscando
    image_url = get_rule34_post(arg)

    if image_url is not None:
        send(bot, update, image_url)
    else:
        send(bot, update, "Prueba con otra cosa...")


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


# crea el updater del bot mediante el token del bot
updater = Updater("593801508:AAF0qCsRxbyKG0I-QSCoh4wwW7A-G6HuccU")

# añade los comandos al bot
updater.dispatcher.add_handler(CommandHandler('start', send_welcome))
updater.dispatcher.add_handler(CommandHandler('rule34', send_rule34))
updater.dispatcher.add_handler(CommandHandler('roll', roll))

updater.start_polling()
updater.idle()
