import telebot  # bot api
import rule34  # the magic is here
import random  # memes
import sys

token = "593801508:AAF0qCsRxbyKG0I-QSCoh4wwW7A-G6HuccU"

bot = telebot.TeleBot(token)


def get_message_arg(message_text):
    message_array = message_text.split(' ')[1:]

    arg = ""  # message argument string
    for message in message_array:
        arg += message + ' '

    return arg


def get_image_by_tag(tag):
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


@bot.message_handler(commands=['rule34'])
def send_post(message):
    arg = get_message_arg(message.text)
    bot.send_sticker(message.chat.id, "CAADAwAD0RAAAsKphwW8SGCoG75C8AI") # sticker de kanna buscando
    image_url = get_image_by_tag(arg)

    if image_url is not None:
        bot.send_message(message.chat.id, image_url)
    else:
        bot.send_message(message.chat.id, "Prueba con otra cosa...")


@bot.message_handler(commands=['start'])
def send_welcome(message):
    bot.send_message(message.chat.id, "Hola k ase")


bot.polling()
