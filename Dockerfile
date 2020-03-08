FROM python:3.7
# create app directory
RUN mkdir -p /app
# using app directory
WORKDIR /app
# copy requirements txr
COPY requirements.txt requirements.txt
# install app dependencies
RUN pip install -r requirements.txt
# copy all project files
COPY . .
# run app
CMD ["python", "-m", "magobot"]