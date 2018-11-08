from telegram.ext import Updater, CommandHandler, RegexHandler
from telegram import Chat, ParseMode


class MagoBot(object):

    def __init__(self, token):
        self.__updater = Updater(token)

    def add_command(self, command):
        self.__updater.dispatcher.add_handler(command)

    def start(self):
        self.__updater.start_polling()
        self.__updater.idle()


class ResponseType(object):
    STICKER = 'sticker'
    TEXT = 'text'


class BotResponse(object):

    def __init__(self, bot, update):
        self.__bot = bot
        self.__update = update
        self.send_reply = False
        self.send_mention = False
        self.parse_mode = None
        self.__message = None
        self.__response_type = ResponseType.TEXT

    def __prepare_params(self):
        chat_type = self.__update.message.chat.type
        chat_id = self.__update.message.chat.id
        message_id = self.__update.message.message_id

        params = {'chat_id': chat_id}

        if chat_type == Chat.GROUP or chat_type == Chat.SUPERGROUP:
            if self.send_reply:
                params['reply_to_message_id'] = message_id
            elif self.send_mention:
                user = self.__update.message.from_user
                markdown = "[@{}](tg://user?id={}) " + self.__message
                self.__message = markdown.format(user.username, user.id)
                params['parse_mode'] = ParseMode.MARKDOWN

        return params

    def send(self, message_type, message_content):
        self.__message = message_content
        params = self.__prepare_params()

        if message_type == ResponseType.TEXT:
            self.__bot.send_message(text=self.__message, **params)
        elif message_type == ResponseType.STICKER:
            self.__bot.send_sticker(sticker=self.__message, **params)


class Command(CommandHandler):

    _command = ''
    _response = None

    def __init__(self):
        super().__init__(self._command, self.__execute, pass_args=True)

    def __execute(self, bot, update, args):
        self._response = BotResponse(bot, update)
        self._on_execution(args)

    def _on_execution(self, args):
        raise NotImplementedError()


class Answer(RegexHandler):

    pattern = ''

    def __init__(self):
        super().__init__(self.pattern, self.__trigger)

    def __trigger(self):
        pass

    def _on_answer(self):
        raise NotImplementedError()
