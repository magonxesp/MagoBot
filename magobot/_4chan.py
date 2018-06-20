from .helper import send, random_element
from basc_py4chan import Board
from basc_py4chan import Thread


def generator_to_array(generator) -> list:
    array = list()

    # el bucle no para hasta que el generador no de mas elementos
    for element in generator:
        array.append(element)

    return array


def get_threads(board_name):
    threads = list()

    board = Board(board_name)
    threads_ids = board.get_all_thread_ids()

    for thread_id in threads_ids:
        threads.append(Thread(board, thread_id))

    return threads


def random_thread(thread_array):
    thread = random_element(thread_array)
    thread.update()
    return thread


def send_4chan_anime_wallpaper(bot, update):
    thread = random_thread(get_threads('w'))
    files = generator_to_array(thread.files())
    send(bot, update, random_element(files))


def send_4chan_random(bot, update):
    thread = random_thread(get_threads('b'))
    send(bot, update, thread.url)


def send_4chan_ecchi(bot, update):
    thread = random_thread(get_threads('e'))
    files = generator_to_array(thread.files())
    send(bot, update, random_element(files))
