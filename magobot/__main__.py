import magobot.settings
import magobot.magobot
import magobot.commands
# import magobot.ai


bot = magobot.magobot.MagoBot(magobot.settings.TOKEN)

# Add command listeners
print('Initializing commands...')
bot.add_handler(magobot.commands.Start())
bot.add_handler(magobot.commands.Roll())
bot.add_handler(magobot.commands.Rule34())
bot.add_handler(magobot.commands.RandomAnimeWallpaper())
bot.add_handler(magobot.commands.RandomBThread())
bot.add_handler(magobot.commands.RandomEcchiThread())
bot.add_handler(magobot.commands.RandomHentaiThread())

# train ai model
# print('Training with default intents...')
# trainer = magobot.ai.Trainer()
# trainer.prepare_intents("intents.json")
# trainer.create_train_lists()
# trainer.train()
# trainer.save_train()

# add ai message handler with trained model
# ai_responder = magobot.ai.AIResponse(trainer.words, trainer.classes, trainer.intents, trainer.get_model())
# ai_message_handler = magobot.magobot.AIMessageHandler(ai_responder)
# bot.add_handler(ai_message_handler)

# start bot
print('All done! bot running...')
bot.start()
