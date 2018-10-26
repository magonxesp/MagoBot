from telegram.ext import Updater, CommandHandler
from magobot.rule34 import send_rule34
from magobot.dummy import send_welcome, roll
from magobot._4chan import send_4chan_anime_wallpaper, send_4chan_random, send_4chan_ecchi, send_4chan_hentai
from magobot.settings import TOKEN

# crea el updater del bot mediante el token del bot
updater = Updater(TOKEN)

# añade los comandos al bot
updater.dispatcher.add_handler(CommandHandler('start', send_welcome))
updater.dispatcher.add_handler(CommandHandler('rule34', send_rule34))
updater.dispatcher.add_handler(CommandHandler('roll', roll))
updater.dispatcher.add_handler(CommandHandler('random', send_4chan_random))
updater.dispatcher.add_handler(CommandHandler('randomw', send_4chan_anime_wallpaper))
updater.dispatcher.add_handler(CommandHandler('ecchi', send_4chan_ecchi))
updater.dispatcher.add_handler(CommandHandler('hentai', send_4chan_hentai))

updater.start_polling()
updater.idle()
