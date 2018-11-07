from magobot.magobot import MagoBot
from magobot.settings import TOKEN
import magobot.commands

bot = MagoBot(TOKEN)

# Add command listeners
bot.add_command(magobot.commands.Start())
bot.add_command(magobot.commands.Roll())
bot.add_command(magobot.commands.Rule34())
bot.add_command(magobot.commands.RandomAnimeWallpaper())
bot.add_command(magobot.commands.RandomBThread())
bot.add_command(magobot.commands.RandomEcchiThread())
bot.add_command(magobot.commands.RandomHentaiThread())

# start bot
bot.start()
