from .helper import send
from basc_py4chan import Board
from basc_py4chan import Thread
from basc_py4chan import Post
import random


def get_4chan_random_thread_from_board(board):
    pass


def get_4chan_random_post_from_thread(thread):
    pass


def send_4chan_anime_wallpaper(bot, update):
    board = Board('w')
    threads_ids = board.get_all_thread_ids()
    index = random.randint(0, len(threads_ids))
    thread = Thread(board, threads_ids[index])
    thread.update()
    file = thread.file_objects()
    send(bot, update, file.file_url)
