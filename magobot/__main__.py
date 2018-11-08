from magobot.magobot import MagoBot
from magobot.settings import TOKEN
import magobot.commands

bot = MagoBot(TOKEN)

# Add command listeners
bot.add_handler(magobot.commands.Start())
bot.add_handler(magobot.commands.Roll())
bot.add_handler(magobot.commands.Rule34())
bot.add_handler(magobot.commands.RandomAnimeWallpaper())
bot.add_handler(magobot.commands.RandomBThread())
bot.add_handler(magobot.commands.RandomEcchiThread())
bot.add_handler(magobot.commands.RandomHentaiThread())

# start bot
bot.start()
