import json
import nltk
from nltk.stem.lancaster import LancasterStemmer
import numpy
import random
import tflearn
import tensorflow
import os
import pickle


class Trainer(object):

    def __init__(self):
        self._stemmer = LancasterStemmer()
        self.words = []
        self.classes = []
        self.documents = []
        self.ignore_words = ['?']
        self._train_x = []
        self._train_y = []
        self.__models_path = 'tf_models'
        self.intents = {}

        if os.path.isdir(self.__models_path) is not True:
            os.mkdir(self.__models_path)

    def prepare_intents(self, intents_path):
        with open(intents_path, encoding="utf-8") as json_data:
            self.intents = json.load(json_data)

        for intent in self.intents['intents']:
            for pattern in intent['patterns']:
                # tokeniza cada palabra de la frase
                w = nltk.word_tokenize(pattern)
                # añade a la lista de palabras
                self.words.extend(w)
                # añade a documents la tupla
                self.documents.append((w, intent['tag']))
                # añade a la lista de classes
                if intent['tag'] not in self.classes:
                    self.classes.append(intent['tag'])

        # ignorar, pasar a minuscula y eliminar duplicados
        self.words = [self._stemmer.stem(w.lower()) for w in self.words if w not in self.ignore_words]
        self.words = sorted(list(self.words))

        # Elimina classes duplicadas
        self.classes = sorted(list(set(self.classes)))

    def create_train_lists(self):
        training = []  # array de datos de entrenamiento
        output_empty = [0] * len(self.classes)  # array vacio para la salida

        # recorrer el set de entrenamiento (palablas de cada frase)
        for document in self.documents:
            # bolsa de palabras
            bag = []
            # lista de palabras tokenizadas
            pattern_words = document[0]
            # procesa (metodo stem) cada palabra
            pattern_words = [self._stemmer.stem(word.lower()) for word in pattern_words]
            # crear bolsa de palabras
            for w in self.words:
                if w in pattern_words:
                    bag.append(1)
                else:
                    bag.append(0)

            # el output es 0 por cada tag y 1 por el tag actual
            output_row = list(output_empty)
            output_row[self.classes.index(document[1])] = 1

            # añadir a la lista de entrenamientos la bolsa y el output
            training.append([bag, output_row])

        # mezclamos los entrenamientos y los ponemos en un numpy.array
        random.shuffle(training)
        training = numpy.array(training)

        # creamos las listas de entrenamientos y tests
        self._train_x = list(training[:, 0])
        self._train_y = list(training[:, 1])

    def __build_neuronal_network(self):
        # reiniciar los datos
        tensorflow.reset_default_graph()
        # Construir la red neuronal
        net = tflearn.input_data([None, len(self._train_x[0])])
        net = tflearn.fully_connected(net, 8)
        net = tflearn.fully_connected(net, 8)
        net = tflearn.fully_connected(net, len(self._train_y[0]), 'softmax')
        net = tflearn.regression(net)

        # definir modelo y preparar el tensorboard
        self._model = tflearn.DNN(net, tensorboard_dir='tflearn_logs')

    def train(self):
        self.__build_neuronal_network()
        # Iniciar entrenamiento
        self._model.fit(self._train_x, self._train_y, 1000, batch_size=8, show_metric=True)
        self._model.save(self.__models_path + '/model.tflearn')

    def save_train(self):
        file = open('training_data', 'wb')

        pickle.dump({
            'words': self.words,
            'classes': self.classes,
            'train_x': self._train_x,
            'train_y': self._train_y
        }, file)

    def restore_train(self):
        if os.path.isfile('training_data') is True:
            file = open('training_data', 'rb')
            data = pickle.load(file)
            self.words = data['words']
            self.classes = data['classes']
            self._train_x = data['train_x']
            self._train_y = data['train_y']
            return True

        return False

    def restore_model(self):
        self.__build_neuronal_network()
        self._model.load(self.__models_path + '/model.tflearn')

    def get_model(self):
        return self._model


class AI(object):

    def __init__(self, model):
        self._model = model
        self._stemmer = LancasterStemmer()


class AIResponse(AI):

    def __init__(self, words, classes, intents, model):
        super().__init__(model)
        self.words = words
        self.classes = classes
        self.intents = intents
        self.__ERROR_THRESHOLD = 0.25  # margen de error

    def clean_up_sentence(self, sentence):
        # tokenize the pattern
        sentence_words = nltk.word_tokenize(sentence)
        # stem each word
        sentence_words = [self._stemmer.stem(word.lower()) for word in sentence_words]
        return sentence_words

    def bow(self, message_string):
        message_words = self.clean_up_sentence(message_string)
        bag = [0] * len(self.words)

        for message_word in message_words:
            for i, word in enumerate(self.words):
                if word == message_word:
                    bag[i] = 1  # palabra encontrada

        return numpy.array(bag)

    def classify(self, message_string):
        # generar probabilidades a partir del modelo
        probabilities = self._model.predict([self.bow(message_string)])
        results = probabilities[0]
        # filtrar predicciones aa traves del margen de error
        results = [[i, r] for i, r in enumerate(results) if r > self.__ERROR_THRESHOLD]
        # ordenar por probabilidad
        results.sort(key=lambda x: x[1], reverse=True)
        return_list = []

        for r in results:
            return_list.append((self.classes[r[0]], r[1]))

        return return_list

    def get_response(self, message_string):
        results = self.classify(message_string)

        if results:
            while results:
                for i in self.intents['intents']:
                    if i['tag'] == results[0][0]:
                        return random.choice(i['responses'])

                results.pop(0)




