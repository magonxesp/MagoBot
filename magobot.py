import telebot  # bot api
import rule34  # the magic is here
import random  # memes

token = "593801508:AAF0qCsRxbyKG0I-QSCoh4wwW7A-G6HuccU"

bot = telebot.TeleBot(token)


def get_message_arg(message_text):
    message_array = message_text.split(' ')[1:]

    arg = ""  # message argument string
    for message in message_array:
        arg += message + ' '

    return arg


def get_image_by_tag(tag):
    images_urls = rule34.getImageURLS(tag)
    #index =
    return images_urls


@bot.message_handler(commands=['rule34'], content_types=['text'])
def send_post(message):
    arg = get_message_arg(message.text)
    image_url = get_image_by_tag(arg)
    bot.send_message(message.chat.id, image_url[1])


@bot.message_handler(['start', 'help'])
def send_welcome(message):
    bot.send_message(message.chat.id, "Hola k ase")


@bot.message_handler(func=lambda m: True)
def echo_all(message):
    bot.reply_to(message, message.text)


bot.polling()
