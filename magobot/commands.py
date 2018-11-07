from magobot.magobot import Command
from xml.dom import minidom
from urllib3 import PoolManager
from urllib3.exceptions import InsecureRequestWarning
from urllib3 import disable_warnings
from basc_py4chan import Board, Thread
import random
import rule34
import asyncio


class Start(Command):

    _command = 'start'
    _send_mention = True

    def _on_execution(self, args):
        return 'Hola k ase'


class Roll(Command):

    _command = 'roll'

    def _on_execution(self, args):
        _max = 100

        if len(args) > 0:
            arg = args[0]

            if arg.isnumeric:
                _max = int(arg)

        return random.randint(0, _max)


class Rule34(Command):

    _command = 'rule34'

    api_url = str()
    xml_string = str()
    tags = ''
    limit = 15

    def __init__(self):
        super().__init__()
        self.loop = asyncio.get_event_loop()
        self.loop.set_debug(True)

    def _on_execution(self, args):
        self.tags = args[0]
        posts = self.get_posts()

        if posts is not False:
            index = random.randint(0, len(posts) - 1)
            return posts[index]
        else:
            return 'Prueba con otra cosa...'

    def __parse_xml(self):
        document = minidom.parseString(self.xml_string)
        posts_elements = document.getElementsByTagName('post')

        posts = list()
        for post_element in posts_elements:
            file_url = post_element.attributes['file_url'].value
            posts.append(file_url)

        return posts

    def __random_page_id(self):
        total_posts = self.loop.run_until_complete(rule34.Rule34(self.loop).totalImages(self.tags))
        # calcula el numero de paginas dividiendo el total de post por 40 posts por pagina aproximandamente
        total_pages = int(total_posts / 100)
        page = random.randint(0, total_pages)
        return page

    def get_posts(self):
        try:
            if self.tags == "":
                return False

            disable_warnings(InsecureRequestWarning)  # Desactiva errores sobre peticiones inseguras
            random_pid = self.__random_page_id()
            self.api_url = rule34.Rule34.urlGen(tags=self.tags, limit=self.limit, PID=random_pid)
            response = PoolManager().request(method='GET', url=self.api_url)

            if response.status == 200:
                self.xml_string = response.data
            else:
                return False

            return self.__parse_xml()
        except rule34.Request_Rejected as request_exception:
            print(request_exception.message)
        except rule34.Rule34_Error as rule34_exception:
            print(rule34_exception.message)

        return False


class ChanCommand(Command):

    def _on_execution(self, args):
        raise NotImplementedError()

    def generator_to_array(self, generator) -> list:
        array = list()

        # el bucle no para hasta que el generador no de mas elementos
        for element in generator:
            array.append(element)

        return array

    def get_threads(self, board_name):
        threads = list()

        board = Board(board_name)
        threads_ids = board.get_all_thread_ids()

        for thread_id in threads_ids:
            threads.append(Thread(board, thread_id))

        return threads

    def random_thread(self, thread_array):
        index = random.randint(0, len(thread_array))
        thread = thread_array[index]
        thread.update()
        return thread


class RandomAnimeWallpaper(ChanCommand):

    _command = 'randomw'

    def __init__(self):
        super().__init__()
        self._board_id = 'w'

    def _on_execution(self, args):
        threads = self.get_threads(self._board_id)
        thread = self.random_thread(threads)
        files = self.generator_to_array(thread.files())
        index = random.randint(0, len(files))
        return files[index]


class RandomBThread(ChanCommand):

    _command = 'random'

    def __init__(self):
        super().__init__()
        self._board_id = 'b'

    def _on_execution(self, args):
        thread = self.random_thread(self.get_threads(self._board_id))
        return thread.url


class RandomHentaiThread(ChanCommand):

    _command = 'hentai'

    def __init__(self):
        super().__init__()
        self._board_id = 'h'

    def _on_execution(self, args):
        thread = self.random_thread(self.get_threads(self._board_id))
        files = self.generator_to_array(thread.files())
        index = random.randint(0, len(files))
        return files[index]


class RandomEcchiThread(ChanCommand):

    _command = 'ecchi'

    def __init__(self):
        super().__init__()
        self._board_id = 'e'

    def _on_execution(self, args):
        thread = self.random_thread(self.get_threads(self._board_id))
        files = self.generator_to_array(thread.files())
        index = random.randint(0, len(files))
        return files[index]
