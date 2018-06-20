from .helper import get_message_arg, send
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


def send_rule34(bot, update):
    arg = get_message_arg(update.message.text)
    bot.send_sticker(update.message.chat_id, "CAADAwAD0RAAAsKphwW8SGCoG75C8AI")  # sticker de kanna buscando
    image_url = get_rule34_post(arg)

    if image_url is not None:
        send(bot, update, image_url)
    else:
        send(bot, update, "Prueba con otra cosa...")
