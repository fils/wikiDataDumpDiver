# Wikidata Dump Parser in Go

# About

 Never done anything with NER or NLP or machine learning.  So thought I might 
 play with the WikiData codex and see what hot mess I can make of it.   Plus
 for work I would love to be able to scan and look for wikidata concepts in some 
 of the data files the facilities are exposing. 

 When I started I found https://github.com/JeroenDeDauw/JsonDumpReader which is 
 in PHP, and I don't know PHP.  So my goal will be to at least try and 
 match the functionality of it in Go. 

## WikiData

WikiData is big. You just won't believe how vastly, hugely, mind-bogglingly big it is.
( forgive me Douglas Adams ).

So I pulled down https://dumps.wikimedia.org/wikidatawiki/entities/latest-all.json.bz2
and wow that that is a lot of JSON.   So to deal with it I needed a streaming 
JSON parser.  

That is the wdparse in cmd.  With that I hope to extract elements into a KV store or 
something I can work with to get to the next step.  

Ref:
* https://www.wikidata.org/wiki/Wikidata:Database_download

## Prose

Being a Go fanboy I found https://github.com/jdkato/prose.  So I am going to see
if there is something I can do with it and the Wikidata data. 

More exactly, I am wondering if I can do something like at:
https://medium.com/@errata.ai/prodigy-prose-radically-efficient-machine-teaching-in-go-93389bf2d772

## Spacy 

In poking around I ran across Spacy.  It looks impressive and links to 
it and other references are blow.  
Some of the examples from these resources are coded up in the scripts 
directory but a person should look at the references for them in context. 

* https://www.linkedin.com/pulse/spacyirl-2019-conference-overview-ivan-bilan/

## Random Notes

Look at 
* https://stanbol.apache.org/overview.html
* https://medium.com/@jdkato/prodigy-prose-radically-efficient-machine-teaching-in-go-93389bf2d772
* https://github.com/allenai/scispacy  Looks like it will get a WikiData model soon (ish)
* https://medium.com/@mgalkin/spacy-irl-2019-and-wikidata-based-ner-64a799c17823
* https://www.researchgate.net/publication/322281320_NECKAr_A_Named_Entity_Classifier_for_Wikidata

