from .helper import get_message_arg, send, random_element
from xml.dom import minidom
from urllib3 import PoolManager
from urllib3.exceptions import InsecureRequestWarning
from urllib3 import disable_warnings
import random
import rule34


class Post(object):
    file_url = str()


class Rule34(object):

    def __init__(self, tags, limit):
        self.api_url = str()
        self.xml_string = str()
        self.tags = tags
        self.limit = limit

    def __parse_xml(self):
        document = minidom.parseString(self.xml_string)
        posts_elements = document.getElementsByTagName('post')

        posts = list()
        for post_element in posts_elements:
            post = Post()
            post.file_url = post_element.attributes['file_url'].value
            posts.append(post)

        return posts

    def __random_page_id(self):
        total_posts = rule34.totalImages(self.tags)
        # calcula el numero de paginas dividiendo el total de post por 40 posts por pagina aproximandamente
        total_pages = int(total_posts / 40)
        page = random.randint(0, total_pages)
        return page

    def get_posts(self):
        try:
            if self.tags == "":
                return False

            disable_warnings(InsecureRequestWarning)  # Desactiva errores sobre peticiones inseguras
            self.api_url = rule34.urlGen(tags=self.tags, limit=self.limit, PID=self.__random_page_id())
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


def send_rule34(bot, update):
    arg = get_message_arg(update.message.text)
    bot.send_sticker(update.message.chat_id, "CAADAwAD0RAAAsKphwW8SGCoG75C8AI")  # sticker de kanna buscando
    # image_url = get_rule34_post(arg)

    posts = Rule34(tags=arg, limit=15).get_posts()

    if posts and len(posts) > 0:
        post = random_element(posts)
        send(bot, update, post.file_url)
        return

    send(bot, update, "Prueba con otra cosa...")
