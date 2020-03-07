import re
from enum import Enum
import importlib


class TokenType(Enum):
    """
    Token types correspond to module names
    """
    COMMAND = 'commands'
    TEXT = 'text'


class TokenParseError(Exception):
    """
    Token parse exception
    """
    pass


class Token(object):
    """
    Token parse class
    """
    __token_pattern = '\[(.+)(:.*){,2}\]'

    def __init__(self, token_str):
        """
        Token constructor
        :param token_str: token string
        :type token_str: str
        """
        self.__token_str = token_str
        self.token_type = ''  # reference to module name
        self.token_class = ''
        self.token_method = ''

    @staticmethod
    def is_token(token_str):
        """
        :param token_str: token string
        :type token_str: str
        :return: boolean
        """
        regex = re.compile(Token.__token_pattern)

        if regex.match(token_str) is not None:
            return True
        else:
            return False

    def __import_module(self):
        module = importlib.import_module('magobot.' + self.token_type)
        return module

    def get_class_instance(self):
        module = self.__import_module()
        _class = getattr(module, self.token_class)
        return _class()

    def parse(self):
        """
        Parse token
        :return void:
        :raise TokenParseError:
        """
        regex = re.compile(Token.__token_pattern)
        match = regex.match(self.__token_str)

        if match is None:
            raise TokenParseError()

        token_parts = self.__token_str.lstrip('[').rstrip(']').split(':')

        if len(token_parts) < 2:
            raise TokenParseError()

        self.token_type = token_parts[0]

        if self.token_type not in [str(t.name).lower() for t in TokenType]:
            raise TokenParseError()

        self.token_type = getattr(TokenType, self.token_type.upper()).value

        self.token_class = token_parts[1]

        if len(token_parts) > 2:
            try:
                self.token_method = token_parts[2]
            except IndexError:
                pass

        return {
            'type': self.token_type,
            'class': self.token_class,
            'method': self.token_method
        }


class TextTokenPreprocessor(object):

    def __init__(self, text, bot=None, update=None):
        """
        StringTokenPreprocessor constructor
        :param text: string to preprocess tokens if available
        :param bot: Bot object for required tokens like command token
        :param update: Update object for required tokens like command token
        :type text: str
        """
        self.__text = text
        self.__bot = bot
        self.__update = update

    def __get_value(self, token_type, token_class_instance, token_class_method=None):
        """
        Executes a class method and return the value
        :param token_type:
        :param token_class_instance:
        :return: str
        """
        return_value = ''

        if token_type == TokenType.COMMAND.value:
            token_class_instance.prepare_response(self.__bot, self.__update)
            token_class_instance.execute(args=None)

        return return_value

    def parse(self):
        text_words = self.__text.split(' ')

        for i, word in enumerate(text_words):
            if Token.is_token(word):
                token = Token(word)
                token.parse()
                text_words[i] = self.__get_value(token.token_type, token.get_class_instance())

        text = ' '.join(text_words)

        if text != '':
            return text
        else:
            return None



