from __future__ import print_function
from collections import defaultdict
from xml.etree import cElementTree as ET
import urllib.request

class Rule34_Error(Exception):
    """Rule34 rejected you"""
    def __init__(self, message, *args):
        self.message = message
        super(Rule34_Error, self).__init__(message, *args)
class Request_Rejected(Exception):
    """The Rule34 API wrapper rejected your request"""
    def __init__(self, message, *args):
        self.message = message
        super(Request_Rejected, self).__init__(message, *args)


def ParseXML(rawXML):
    """Parses entities as well as attributes following this XML-to-JSON "specification"
    Using https://stackoverflow.com/a/10077069"""
    if "Search error: API limited due to abuse" in str(rawXML.items()):
        raise Rule34_Error('Rule34 rejected your request due to "API abuse"')

    d = {rawXML.tag: {} if rawXML.attrib else None}
    children = list(rawXML)
    if children:
        dd = defaultdict(list)
        for dc in map(ParseXML, children):
            for k, v in dc.items():
                dd[k].append(v)
        d = {rawXML.tag: {k:v[0] if len(v) == 1 else v for k, v in dd.items()}}
    if rawXML.attrib:
        d[rawXML.tag].update(('@' + k, v) for k, v in rawXML.attrib.items())
    if rawXML.text:
        text = rawXML.text.strip()
        if children or rawXML.attrib:
            if text:
              d[rawXML.tag]['#text'] = text
        else:
            d[rawXML.tag] = text
    return d

def urlGen(tags=None, limit=None, id=None, PID=None, deleted=None, **kwargs):
    """Generates a URL to access the api using your input:
    Arguments:
        "limit"  ||str ||How many posts you want to retrieve
        "pid"    ||int ||The page number.
        "tags"   ||str ||The tags to search for. Any tag combination that works on the web site will work here. This includes all the meta-tags. See cheatsheet for more information.
        "cid"    ||str ||Change ID of the post. This is in Unix time so there are likely others with the same value if updated at the same time.
        "id"     ||int ||The post id.
        "deleted"||bool||If True, deleted posts will be included in the data
    All arguments that accept strings *can* accept int, but strings are recommended
    If none of these arguments are passed, None will be returned
    """
    #I have no intentions of adding "&last_id=" simply because its response can easily be massive, and all it returns is ``<post deleted="[ID]" md5="[String]"/>`` which has no use as far as im aware
    URL = "https://rule34.xxx/index.php?page=dapi&s=post&q=index"
    
    if PID != None:
        URL += "&pid={}".format(PID)
    if limit != None:
        URL += "&limit={}".format(limit)
    if id != None:
        URL += "&id={}".format(id)
    if tags != None:
        tags = str(tags).replace(" ", "+")
        if str(tags) == "":
            raise Request_Rejected('Submitting this action WILL result in your API access being suspended due to API abuse\nReason: tag="" will request every single image on rule34')
        URL += "&tags={}".format(tags)
    if deleted == True:
        url += "&deleted=show"
    if PID != None or limit != None or id != None or tags != None:
        return URL
    else:
        return None

def totalImages(tags):
    """Get an int of how many images are on rule34.xxx
    Argument: tags (string)"""
    XMLData = urllib.request.urlopen(urlGen(tags)).read()
    XMLData = ET.XML(XMLData)
    XML = ParseXML(XMLData)
    return int(XML['posts']['@count'])

def getImageURLS(tags):
    """Returns a list of all images/webms/gifs it can find
    This function can take a LONG time to finish with huge tags. E.G. in my testing "gay" took 200seconds to finish (740 pages)
    Argument: tags (string)"""
    
    imgList = []
    if totalImages(tags) != 0:
        PID = 0
        t = True
        imgList = []
        while t:
            tempURL = urlGen(tags=tags, PID=PID)
            XML = urllib.request.urlopen(tempURL).read() #<-- This line is what causes this function to take as long as it does

            XML = ParseXML(ET.XML(XML))

            if len(imgList) >= int(XML['posts']['@count']): #"if we're out of images to process"
                t = False #"end the loop"
            else:
                for data in XML['posts']['post']:
                    imgList.append(str(data['@file_url']))
            PID += 1
        return imgList
    else:
        return None

def getPostData(PostID):
    """"Returns a dict with all the information available about the post
    Argument: PostID (string or ID)"""
    url = urlGen(id=str(PostID))
    XML = urllib.request.urlopen(url).read()
    XML = ParseXML(ET.XML(XML))
    data = XML['posts']['post']