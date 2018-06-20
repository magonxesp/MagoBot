from telegram import Chat, ParseMode


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
