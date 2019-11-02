import re 
import string 
import nltk 
import spacy 
import pandas as pd 
import numpy as np 
import math 
from tqdm import tqdm 

from spacy.matcher import Matcher 
from spacy.tokens import Span 
from spacy import displacy 

pd.set_option('display.max_colwidth', 200)

# load spaCy model
nlp = spacy.load("en_core_web_sm")

# sample text 
text = "GDP in developing countries such as Vietnam will continue growing at a high rate." 

# create a spaCy object 
doc = nlp(text)


# print token, dependency, POS tag 
for tok in doc: 
  print(tok.text, "-->",tok.dep_,"-->", tok.pos_)

#define the pattern 
pattern = [{'POS':'NOUN'}, 
           {'LOWER': 'such'}, 
           {'LOWER': 'as'}, 
           {'POS': 'PROPN'}] #proper noun

# Matcher class object 
matcher = Matcher(nlp.vocab) 
matcher.add("matching_1", None, pattern) 

matches = matcher(doc) 
span = doc[matches[0][1]:matches[0][2]] 

print(span.text)

# Matcher class object
matcher = Matcher(nlp.vocab)

#define the pattern
pattern = [{'DEP':'amod', 'OP':"?"}, # adjectival modifier
           {'POS':'NOUN'},
           {'LOWER': 'such'},
           {'LOWER': 'as'},
           {'POS': 'PROPN'}]

matcher.add("matching_1", None, pattern)
matches = matcher(doc)

span = doc[matches[0][1]:matches[0][2]]
print(span.text)

doc = nlp("Here is how you can keep your car and the other vehicles clean.") 

# print dependency tags and POS tags
for tok in doc: 
  print(tok.text, "-->",tok.dep_, "-->",tok.pos_)


# Matcher class object 
matcher = Matcher(nlp.vocab) 

#define the pattern 
pattern = [{'DEP':'amod', 'OP':"?"}, 
           {'POS':'NOUN'}, 
           {'LOWER': 'and', 'OP':"?"}, 
           {'LOWER': 'or', 'OP':"?"}, 
           {'LOWER': 'other'}, 
           {'POS': 'NOUN'}] 
           
matcher.add("matching_1", None, pattern) 

matches = matcher(doc) 
span = doc[matches[0][1]:matches[0][2]] 
print(span.text)
