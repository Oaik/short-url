<h3 align="center">Shorter url</h3>
<div align="center">
  
  [![GitHub contributors](https://img.shields.io/github/contributors/Oaik/short-url)](https://github.com/Oaik/short-url/contributors)
  [![GitHub issues](https://img.shields.io/github/issues/Oaik/short-url)](https://github.com/Oaik/short-url/issues)
  [![GitHub forks](https://img.shields.io/github/forks/Oaik/short-url)](https://github.com/Oaik/short-url/network)
  [![GitHub stars](https://img.shields.io/github/stars/Oaik/short-url)](https://github.com/Oaik/short-url/stargazers)
  [![GitHub license](https://img.shields.io/github/license/Oaik/short-url)](https://github.com/Oaik/short-url/blob/master/LICENSE)
  <img src="https://img.shields.io/github/languages/count/Oaik/short-url" />
  <img src="https://img.shields.io/github/languages/top/Oaik/short-url" />
  <img src="https://img.shields.io/github/languages/code-size/Oaik/short-url" />
  <img src="https://img.shields.io/github/issues-pr-raw/Oaik/short-url" />
</div>

## Table of Contents
- [Features of the Project](#features)
- [Built with](#built-with)
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



#### Example:

-running http://localhost:7070/short/www.google.com in postman/crul

you will see output of hash string

ex: shorturl:  goossjl548

-running http://localhost:7070/goto/goossjl548 

will redirect to www.google.com

## built with
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