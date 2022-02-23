<h3 align="center">Shorter url</h3>
<div align="center">
  
  [![GitHub contributors](https://img.shields.io/github/contributors/Oaik/Redurl-shorterdit)](https://github.com/Oaik/url-shorter/contributors)
  [![GitHub issues](https://img.shields.io/github/issues/Oaik/url-shorter)](https://github.com/Oaik/url-shorter/issues)
  [![GitHub forks](https://img.shields.io/github/forks/Oaik/url-shorter)](https://github.com/Oaik/url-shorter/network)
  [![GitHub stars](https://img.shields.io/github/stars/Oaik/url-shorter)](https://github.com/Oaik/url-shorter/stargazers)
  [![GitHub license](https://img.shields.io/github/license/Oaik/url-shorter)](https://github.com/Oaik/url-shorter/blob/master/LICENSE)
  <img src="https://img.shields.io/github/languages/count/Oaik/url-shorter" />
  <img src="https://img.shields.io/github/languages/top/Oaik/url-shorter" />
  <img src="https://img.shields.io/github/languages/code-size/Oaik/url-shorter" />
  <img src="https://img.shields.io/github/issues-pr-raw/Oaik/url-shorter" />
</div>

## Table of Contents
- [Features of the Project](#features)
- [Built with](#built)
- [How to Run](#how-to-run)

## Features

Used to short the url links
the website will converted into hashed string that cosists of random digits and letters that will used to map to this website

there is 2 api calls you can make:
http://localhost:7070/short/*website*
http://localhost:7070/goto/*short-url*


1) http://localhost:7070/short/*website*

ex:
http://localhost:7070/short/www.google.com

Will short the website "www.google.com" to shorter link

2) http://localhost:7070/goto/*short-url*

ex:
http://localhost:7070/goto/goossjl548

will redirect to the website that have shorturl that equal to "goossjl548" in the database

Example:
-running http://localhost:7070/short/www.google.com in postman/crul
you will see output of hash string
ex: shorturl:  goossjl548
-running http://localhost:7070/goto/goossjl548 
will redirect to www.google.com

## built
Langague used: Go Langague

## how to run

Make sure that you have a table in your database called "url" that contains 2 fields 
"shortUrl" of type text and primary key 
"website" of type text and not null

next you need to make sure that your const config in postgress/main.go is correct
ex: dbname, password, ...

make sure that you in the parent directory and then run
```console
$ go run main.go
```