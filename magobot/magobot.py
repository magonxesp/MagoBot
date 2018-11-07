from telegram.ext import Updater, CommandHandler
from telegram import Chat, ParseMode


class MagoBot(object):

    def __init__(self, token):
        self.__updater = Updater(token)

    def add_command(self, command):
        self.__updater.dispatcher.add_handler(command)

    def start(self):
        self.__updater.start_polling()
        self.__updater.idle()
    """
    def __send_response(self, bot, update, message_text, no_reply_in_group=False, mention_user=False):
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
    """


class Command(CommandHandler):

    _command = ''

    " Only when command is executed on group "
    _send_reply = False
    _send_mention = False

    __bot = None
    __update = None

    def __init__(self):
        super().__init__(self._command, self.__execute, pass_args=True)

    def __execute(self, bot, update, args):
        self.__bot = bot
        self.__update = update

        messages = self._on_execution(args)

        if type(messages) is not 'list':
            messages = [messages]

        for message in messages:
            self.__send(message)

    def __send(self, message_text):
        chat = self.__update.message.chat
        message = self.__update.message

        if chat.type == Chat.GROUP or chat.type == Chat.SUPERGROUP:
            if self._send_reply:
                self.__bot.send_message(chat_id=chat.id, text=message_text, reply_to_message_id=message.message_id)
            elif self._send_mention:
                user = message.from_user
                markdown = "[@{}](tg://user?id={}) " + message_text
                message_text = markdown.format(user.username, user.id)
                self.__bot.send_message(chat_id=chat.id, text=message_text, parse_mode=ParseMode.MARKDOWN)
            else:
                self.__bot.send_message(chat_id=chat.id, text=message_text)
        else:
            self.__bot.send_message(chat.id, message_text)

    def _on_execution(self, args):
        raise NotImplementedError()
