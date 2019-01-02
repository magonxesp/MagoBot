from telegram.ext import Updater, CommandHandler, RegexHandler
from telegram import Chat, ParseMode
from magobot.token import TextTokenPreprocessor


class MagoBot(object):

    def __init__(self, token):
        self.__updater = Updater(token)

    def add_handler(self, handler):
        self.__updater.dispatcher.add_handler(handler)

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
        super().__init__(self._command, self.__on_call, pass_args=True)

    def __on_call(self, bot, update, args):
        self._response = BotResponse(bot, update)
        self.execute(args)

    def prepare_response(self, bot, update):
        self._response = BotResponse(bot, update)

    def execute(self, args):
        raise NotImplementedError()

    def get_name(self):
        return self._command


class MessageHandler(object):
    
    def __init__(self):
        self.__handlers = []
        self.__callbacks = []
        self.response = ''

    def _trigger(self, bot, update):
        """
        :param bot: bot object
        :type bot: telegram.Bot
        :param update: message update object
        :type update: telegram.Update
        :return: void
        """
        for callback in self.__callbacks:
            callback(update)

        response = BotResponse(bot, update)
        response.send(ResponseType.TEXT, self.response)

    def add_trigger_event_callback(self, callback):
        self.__callbacks.append(callback)

    def add_pattern(self, pattern):
        handler = RegexHandler(pattern, self._trigger)
        self.__handlers.append(handler)

    def get_handlers(self):
        return self.__handlers


class AIMessageHandler(RegexHandler):

    def __init__(self, ai):
        """
        :param ai: AI response picker class
        :type ai: magobot.ai.AIResponse
        """
        super().__init__('(.*)', self.__trigger)
        self.__ai = ai

    def __trigger(self, bot, update):
        """
        :param bot: bot object
        :type bot: telegram.Bot
        :param update: message update object
        :type update: telegram.Update
        :return: void
        """
        response = self.__ai.get_response(update.message.text)
        preprocesor = TextTokenPreprocessor(response, bot=bot, update=update)
        parsed_response = preprocesor.parse()

        if parsed_response is not None:
            BotResponse(bot, update).send(ResponseType.TEXT, response)


